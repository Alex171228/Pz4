package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

var ErrNotFound = errors.New("task not found")

type Repo struct {
	mu       sync.RWMutex
	seq      int64
	items    map[int64]*Task
	filePath string
}

// NewRepo creates in-memory repo without persistence.
func NewRepo() *Repo {
	return &Repo{items: make(map[int64]*Task)}
}

// NewRepoWithFile creates repo and loads tasks from JSON file if it exists.
// All mutating operations persist changes to the same file.
func NewRepoWithFile(path string) (*Repo, error) {
	r := &Repo{items: make(map[int64]*Task), filePath: path}
	if err := r.load(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Repo) List() []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		out = append(out, t)
	}
	// stable order by ID
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

func (r *Repo) Get(id int64) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (r *Repo) Create(title string) *Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	now := time.Now()
	t := &Task{ID: r.seq, Title: title, CreatedAt: now, UpdatedAt: now, Done: false}
	r.items[t.ID] = t
	r.saveLocked()
	return t
}

func (r *Repo) Update(id int64, title string, done bool) (*Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	t.Title = title
	t.Done = done
	t.UpdatedAt = time.Now()
	r.saveLocked()
	return t, nil
}

func (r *Repo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	r.saveLocked()
	return nil
}

// load reads existing tasks from the JSON file, if present.
func (r *Repo) load() error {
	if r.filePath == "" {
		return nil
	}
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			// ensure directory exists for later saves
			_ = os.MkdirAll(filepath.Dir(r.filePath), 0o755)
			return nil
		}
		return fmt.Errorf("read file: %w", err)
	}
	var arr []*Task
	if len(data) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, &arr); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	r.items = make(map[int64]*Task, len(arr))
	var maxID int64
	for _, t := range arr {
		if t.ID > maxID { maxID = t.ID }
		r.items[t.ID] = t
	}
	r.seq = maxID
	return nil
}

// saveLocked writes tasks to the JSON file. Caller must hold r.mu (write).
func (r *Repo) saveLocked() {
	if r.filePath == "" {
		return
	}
	arr := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		arr = append(arr, t)
	}
	// stable order by ID
	sort.Slice(arr, func(i, j int) bool { return arr[i].ID < arr[j].ID })
	b, err := json.MarshalIndent(arr, "", "  ")
	if err != nil {
		return // best-effort; could log if logger injected
	}
	_ = os.MkdirAll(filepath.Dir(r.filePath), 0o755)
	_ = os.WriteFile(r.filePath, b, 0o644)
}
