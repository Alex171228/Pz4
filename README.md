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

## Запросы и их результаты
