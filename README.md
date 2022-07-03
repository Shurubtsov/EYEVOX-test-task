# EYEVOX-test-task

## Описание
Тестовое задание от компании EYEVOX. Простой API сервис с парой эндпоинтов для реализации простой back-end системы чата.
Используемые технологии: **Go, CleanEnv, Httprouter, Pgx, Testify, PostgreSQL, Docker.**

## Запуск приложения
*Для поднятия приложения в контейнере нужно запустить команду*:
>     docker-compose up

Контейнеры будут запущены на следующих портах:
1. PostgreSQL -> 5432
2. Adminer -> 8080
3. Приложение -> 4000

Для изменения портов отредактируйте `docker-compose.yml` файл

Для локального запуска требуется отредактировать `config.yml` файл:
```yml
storage:
  host: localhost #заменив эту строку
  port: 5432
  database: go_test_db
  username: postgres
  password: admin
```
И иметь установленный PostgresSQL.

## Тестирование
В приложении реализованые следующие endpoints:
```
[POST] http://localhost:порт/chats/create
[POST] http://localhost:порт/message/create/:название чата
[GET] http://localhost:порт/messages/:название чата/:страница 
[GET] http://localhost:порт/message/:id сообщения
```
Примеры тела POST запросов в `json` формате:
```json
// Тело запроса для создания чата
{
  "name": "чат",
  "founder_nickname": "кто создал чат"
}

// Тело запроса для создания сообщения в чате
{
  "creator_nickname": "создатель",
  "text_message": "текст сообщения"
}
```
