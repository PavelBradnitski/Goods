# Auth & Book Service
Этот проект состоит из двух микросервисов: 
 - Auth Service (сервис авторизации)  
 - Book Service (сервис хранения книг). 
## Технологии
Проект разработан с использованием следующих технологий:

- **Go** (Golang) - для разработки серверной логики.
- **MongoDb** - для хранения данных.
- **Docker** - для контейнеризации проекта и упрощения 
# Запуск проекта 
1. Убедитесь, что установлен Docker и Docker Compose
2. Клонируйте репозиторий

* git clone https://github.com/PavelBradnitski/Goods.git
* cd Goods
3. Настройте .env файлы

4. Запустите сервисы в Docker

* docker-compose up --build -d`

После успешного запуска сервисы будут доступны по:

- Auth Service → http://localhost:8080

- Book Service → http://localhost:8081

- MongoDB → mongodb://localhost:27017

# Работа с MongoDB

## Подключение к базе через CLI

* docker exec -it mongodb mongosh

## Просмотр существующих баз

* show dbs

## Просмотр коллекций

* show collections
## Вывод всех пользователей

* db.users.find().pretty()

## Вывод всех книг

* db.books.find().pretty()

# API Документация
Swagger-документация доступна по адресу:
- Auth Service → http://localhost:8080/swagger/index.html
- Book Service → http://localhost:8081/swagger/index.html