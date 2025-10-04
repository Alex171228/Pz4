# Практическое занятие №4  
## Тема
Маршрутизация с chi (альтернатива — gorilla/mux). Создание небольшого CRUD-сервиса «Список задач».
## Цели
1.	Освоить базовую маршрутизацию HTTP-запросов в Go на примере роутера chi.
2.	Научиться строить REST-маршруты и обрабатывать методы GET/POST/PUT/DELETE.
3.	Реализовать небольшой CRUD-сервис «ToDo» (без БД, хранение в памяти).
4.	Добавить простое middleware (логирование, CORS).
5.	Научиться тестировать API запросами через curl/Postman/HTTPie.
## Структура проекта
<img width="326" height="546" alt="изображение" src="https://github.com/user-attachments/assets/bdea8e6c-ccbc-4f58-925a-d96e8ef0ebad" />

## Запуск проекта
```
cd pz4-todo
go mod tidy
go run .
```
## Проверка работы сервера
<img width="1697" height="591" alt="изображение" src="https://github.com/user-attachments/assets/8ab50271-9da4-4dbd-a052-5afbb7e14700" /> 

## Запросы и результат их выполения
### 1. Создание задачи

<img width="1704" height="681" alt="изображение" src="https://github.com/user-attachments/assets/3a7e0d72-3643-4fa5-9968-4b2443349168" /> 

### 2. Получение списка

<img width="1700" height="644" alt="изображение" src="https://github.com/user-attachments/assets/a3756911-9788-4f6b-8de6-88e103542ea6" /> 

### 3. Получение по ID

<img width="1694" height="631" alt="изображение" src="https://github.com/user-attachments/assets/5a775b5b-f98e-483f-98c8-5dd661d3f855" /> 

### 4. Обновление задачи

<img width="1696" height="685" alt="изображение" src="https://github.com/user-attachments/assets/214995fb-866a-41dd-bb76-2da8cbc087e4" />

### 5. Удаление задачи

<img width="1698" height="414" alt="изображение" src="https://github.com/user-attachments/assets/2015c9be-df10-4f3c-b665-3f5e7d3c7127" />

### 6. Проверка валидации

<img width="1031" height="119" alt="изображение" src="https://github.com/user-attachments/assets/b9ebfc2e-497b-4d5b-8b2c-494258bfa703" />

### 7. Проверка невалидного ID

<img width="895" height="58" alt="изображение" src="https://github.com/user-attachments/assets/a26d6c84-7ebf-43c1-b7cb-e3c39b7647b5" />

### 8. Фильтрация по статусу

<img width="1683" height="601" alt="изображение" src="https://github.com/user-attachments/assets/af9d990f-f2ff-4a64-866c-58642d3afb82" /> 

<img width="1699" height="625" alt="изображение" src="https://github.com/user-attachments/assets/2311ddb5-adc8-425c-8fa8-fd418c97c1a3" /> 

### 9. Пагинация

<img width="1700" height="628" alt="изображение" src="https://github.com/user-attachments/assets/20677b96-ed1a-40de-a1fe-2919a6beac94" /> 

<img width="1698" height="625" alt="изображение" src="https://github.com/user-attachments/assets/e166c34f-63c9-41e3-8e22-d19f48039998" />

### 10. Проверка CORS (OPTIONS)

<img width="1692" height="433" alt="изображение" src="https://github.com/user-attachments/assets/dca3910b-11b2-4c6d-9fa1-7805862f531d" />

### 11. Проверка сохранения на диск
Останавливаем сервер  
В data/tasks.json сохранены задачи

<img width="658" height="639" alt="изображение" src="https://github.com/user-attachments/assets/d49ad8bf-44fb-4c98-a8c3-06351143225b" />

Запускаем сервер снова и выводим задачи


<img width="1690" height="633" alt="изображение" src="https://github.com/user-attachments/assets/db1ae169-cf2d-44ea-9b04-4473374ddac1" />

## Фрагменты кода роутера, middleware, обработчиков

### 1. Основной роутер (`main.go`)
Реализует запуск сервера, подключение middleware и маршрутов с версионированием `/api/v1`.

```
go
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"example.com/pz4-todo/internal/task"
	myMW "example.com/pz4-todo/pkg/middleware"
)

func main() {
	repo := task.NewRepo()
	h := task.NewHandler(repo)

	r := chi.NewRouter()
	r.Use(chimw.RequestID)
	r.Use(chimw.Recoverer)
	r.Use(myMW.Logger)
	r.Use(myMW.SimpleCORS)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Версионирование API
	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/tasks", h.Routes())
	})

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
2. Middleware: логирование (pkg/middleware/logger.go)
go
Копировать код
package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger — middleware, выводящее информацию о каждом запросе в консоль.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

```
### 2. Middleware: логирование

```
package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger — middleware, выводящее информацию о каждом запросе в консоль.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}


```
### 3. Middleware: CORS

```
package middleware

import "net/http"

// SimpleCORS — middleware для разрешения кросс-доменных запросов.
func SimpleCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}


```
### 4. Обработчик списка задач

```
// list — получение списка задач с поддержкой фильтра и пагинации.
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	items := h.repo.List()
	q := r.URL.Query()

	// Фильтр по статусу done
	doneStr := q.Get("done")
	if doneStr == "true" || doneStr == "false" {
		want := doneStr == "true"
		filtered := make([]*Task, 0, len(items))
		for _, t := range items {
			if t.Done == want {
				filtered = append(filtered, t)
			}
		}
		items = filtered
	}

	// Пагинация
	page, limit := 1, 10
	if v, err := strconv.Atoi(q.Get("page")); err == nil && v > 0 {
		page = v
	}
	if v, err := strconv.Atoi(q.Get("limit")); err == nil && v > 0 && v <= 100 {
		limit = v
	}
	start := (page - 1) * limit
	if start > len(items) {
		start = len(items)
	}
	end := start + limit
	if end > len(items) {
		end = len(items)
	}

	writeJSON(w, http.StatusOK, items[start:end])
}

```
### 5. Обработчик создания задачи

```
type createReq struct {
	Title string `json:"title"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		httpError(w, http.StatusBadRequest, "invalid json: require non-empty title")
		return
	}
	if n := len(req.Title); n < 3 || n > 100 {
		httpError(w, http.StatusBadRequest, "title length must be 3..100")
		return
	}
	t := h.repo.Create(req.Title)
	writeJSON(w, http.StatusCreated, t)
}

```
### 6. Обработчик обновления задачи

```
type updateReq struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, bad := parseID(w, r)
	if bad {
		return
	}
	var req updateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		httpError(w, http.StatusBadRequest, "invalid json: require non-empty title")
		return
	}
	if n := len(req.Title); n < 3 || n > 100 {
		httpError(w, http.StatusBadRequest, "title length must be 3..100")
		return
	}
	t, err := h.repo.Update(id, req.Title, req.Done)
	if err != nil {
		httpError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, t)
}


```
## Обработка ошибок и коды ответа

Система обработки ошибок реализована централизованно через вспомогательные функции `httpError` и `writeJSON`, которые отвечают за единообразный формат ответов и корректные HTTP-коды.

### Основные принципы
1. Все ответы сервера возвращаются в формате **JSON**.  
2. Ошибки всегда имеют структуру:
   ```json
   {"error": "описание ошибки"}

## Таблица результатов тестирования


| №  | Метод  | Маршрут                      | Код  | Результат |
|:--:|:-------|:-----------------------------|:----:|:-----------|
| 1  | GET    | /health                      | 200  | ОК |
| 2  | POST   | /api/v1/tasks                | 201  | ОК |
| 3  | GET    | /api/v1/tasks                | 200  | ОК |
| 4  | GET    | /api/v1/tasks/1              | 200  | ОК |
| 5  | PUT    | /api/v1/tasks/1              | 200  | ОК |
| 6  | DELETE | /api/v1/tasks/1              | 204  | ОК |
| 7  | POST   | /api/v1/tasks (короткий title) | 400  | ОК |
| 8  | GET    | /api/v1/tasks/abc            | 400  | ОК |
| 9  | GET    | /api/v1/tasks?done=true      | 200  | ОК |
| 10 | GET    | /api/v1/tasks?page=2&limit=2 | 200  | ОК |
| 11 | OPTIONS| /api/v1/tasks                | 204  | ОК |
| 12 | Проверка файла JSON | data/tasks.json | — | ОК |

## Выводы

В ходе выполнения практического задания был разработан полнофункциональный REST API-сервис на языке Go с использованием фреймворка `chi`.  
Реализованы все обязательные и дополнительные пункты задания:

- настроен HTTP-сервер и маршрутизация с версионированием `/api/v1`;  
- созданы CRUD-обработчики для сущности `Task`;  
- добавлены middleware для логирования и CORS;  
- реализовано хранение данных в файле `data/tasks.json`;  
- обеспечена валидация входных данных, обработка ошибок и корректные коды ответов;  
- добавлены дополнительные возможности: пагинация, фильтрация и проверка сохранения данных на диск.


