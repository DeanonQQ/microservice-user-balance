# microservice-user-balance
## Microservice for working with user balance
Микросервис для работы с балансом пользователей.
Реализовано небольшое приложение, которое позволяет изменять баланс пользователей в базе данных через API.

## Перечень запросов

> GET /user/balance?Id=< int >

Пример возвращаемого результата:

    {
        "id": 3,
        "name": "Ivan",
        "age": 24,
        "email": "simpleEmai1@qwe.qw",
        "balance": 10500.00
    }

> POST /user/balance/

### Action "Add" - увеличить баланс пользователя

Пример body:

    {
        "Id": 3,
        "Sum": 200,
        "Action": "Add"
    }

Пример возвращаемого результата:

    {
        "id": 3,
        "name": "Ivan",
        "age": 24,
        "email": "simpleEmai1@qwe.qw",
        "balance": 15900.00
    }

### Action "Substract" - уменьшить баланс пользователя

Пример body:

    {
        "Id": 3,
        "Sum": 200,
        "Action": "Substract"
    }

Пример возвращаемого результата:

    {
        "id": 3,
        "name": "Ivan",
        "age": 24,
        "email": "simpleEmai1@qwe.qw",
        "balance": 15700.00
    }

### Action "Send" - перевести сумму пользователю

Пример body:

    {
        "id": 4,
        "sum": 2000,
        "action": "Send",
        "destination": 3
    }

Пример возвращаемого результата:

    "Success"