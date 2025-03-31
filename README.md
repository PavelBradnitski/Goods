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
```shell
git clone https://github.com/PavelBradnitski/Goods.git
cd Goods
```
3. Настройте .env файлы

4. Запустите сервисы в Docker
```shell
docker-compose up --build -d
```

После успешного запуска сервисы будут доступны по:

- Auth Service → http://localhost:8080

- Book Service → http://localhost:8081

- MongoDB → mongodb://localhost:27017

# Работа с MongoDB

## Подключение к базе через CLI
```shell
docker exec -it mongodb mongosh
```
## Просмотр существующих баз
```shell
show dbs
```
## Просмотр коллекций
```shell
show collections
```
## Вывод всех пользователей
```shell
db.users.find().pretty()
```
## Вывод всех книг
```shell
db.books.find().pretty()
```
# API Документация
Swagger-документация доступна по адресу:
- Auth Service → http://localhost:8080/swagger/index.html
- Book Service → http://localhost:8081/swagger/index.html