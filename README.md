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
