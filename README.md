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
Останавливаем сервер \n  
В data/tasks.json сохранены задачи

<img width="658" height="639" alt="изображение" src="https://github.com/user-attachments/assets/d49ad8bf-44fb-4c98-a8c3-06351143225b" />

Запускаем сервер снова и выводим задачи


<img width="1690" height="633" alt="изображение" src="https://github.com/user-attachments/assets/db1ae169-cf2d-44ea-9b04-4473374ddac1" />
