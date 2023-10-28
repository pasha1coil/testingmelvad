# Testing Melvad
<details>
  <summary>Содержание</summary>
  <ol>
    <li><a href="#установка-и-запуск">Установка и запуск</a></li>
    <li><a href="#реализовано">Реализовано</a></li>
    <li><a href="#примеры-запросов">Примеры запросов</a></li>
  </ol>
</details>

## Установка и запуск

Клонировать проект.

Возможно вам нужно будет выставить свои конфиги, для этого есть файл .env.

Далее через `терминал`:
- `go test -v ./... ` - запустить тесты
- `docker-compose up` - загрузить докер образы
- `go run cmd/main.go` - запустить проект

## Реализовано

1. Метод добавления данных в  redis db с инкрементацией по ключу.
2. Метод преобразования данных по `text` и `key` в `hex` строку HMAC-SHA512.
3. Метод добавления пользователя по полям `name` и `age`, также проверяет на существование таблицу `users`, если ее нет, то создает.
4. Тесты для `service` и `repository`.

## Примеры запросов

Для взаимодействия с сервером есть 7 способов:
1. (POST) /redis/incr
2. (POST) /sign/hmacsha512
3. (POST) /postgres/users

- input (POST) /redis/incr
```
{
     "key": "age", 
     "value": 19
}
```
- Output 

```
{
    "value": 20
}
```

- input (POST) /sign/hmacsha512
```
{
     "text": "test",
     "key": "test123"
}
```

- Output
```
{
    "HMAC": "55b5bb82607f1c64d187d1089b6cc27e1667cb24de14bb26f8a0ebc3b3dc7595096290b35d7df477a57e69059f2a00946f1737d3eaef3e6c73fa29ac400b8bdb"
}
```

- input (POST) /postgres/users
```
{
     "name": "Alex",
     "age": 21
}
```

- Output
```
{
    "id": 1
}
```