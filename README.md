# pz4-todo (Go + chi)

CRUD «Список задач» на Go с chi. Выполнено всё по заданию:
- chi-маршрутизация, REST CRUD (GET/POST/PUT/DELETE)
- in-memory хранилище + **сохранение на диск (JSON)**: файл `./data/tasks.json` (или через `DATA_FILE`)
- middleware: **логирование** и **CORS**
- тестирование через curl/Postman
- доп. задания: **валидация title (3..100)**, **пагинация/фильтр done**, **версионирование `/api/v1`**

## Запуск
```bash
go mod tidy
go run .
# переменная для пути файла (опционально)
# DATA_FILE=/path/to/tasks.json go run .
```
Проверка: http://localhost:8080/health → `OK`.

## Примеры запросов
Создать:
```bash
curl -X POST http://localhost:8080/api/v1/tasks   -H "Content-Type: application/json"   -d '{"title":"Выучить chi"}'
```
Список:
```bash
curl "http://localhost:8080/api/v1/tasks"
```
Фильтр по `done`:
```bash
curl "http://localhost:8080/api/v1/tasks?done=true"
```
Пагинация:
```bash
curl "http://localhost:8080/api/v1/tasks?page=1&limit=10"
```
Получить по id:
```bash
curl http://localhost:8080/api/v1/tasks/1
```
Обновить:
```bash
curl -X PUT http://localhost:8080/api/v1/tasks/1   -H "Content-Type: application/json"   -d '{"title":"Выучить chi глубже","done":true}'
```
Удалить:
```bash
curl -X DELETE http://localhost:8080/api/v1/tasks/1
```

## Поведение и коды
- POST → 201 + JSON созданной задачи (валидация: title 3..100)
- GET список → 200, поддерживает `done`, `page`, `limit`
- GET по id → 200 или 404 `{"error":"task not found"}`
- PUT → 200, валидация title 3..100
- DELETE → 204 без тела
- Неверный id/JSON → 400 `{"error":"..."}`

## Сохранение на диск
Файл `./data/tasks.json` создаётся автоматически. Можно указать путь через `DATA_FILE`.
Данные читаются при старте и сохраняются при каждом Create/Update/Delete.
