# Тестовое задание "Интернет-провайдер"

## Примеры запросов

## GET, localhost:8000/user_tariffs , показывает тариф у выбранного пользователя
### Пример запроса:
{
    "user_id": 1
}
### Пример ответа:
{
    "userId":1,
    "tariffs":"low"
}

## PUT, localhost:8000/user_tariffs , обновляет цену выбранного тарифа
### Пример запроса:

{
    "name": "low",
    "price": 10
}

### Пример ответа: 
"Successfully!"

## POST, localhost:8000/user_tariffs , добавляет новый тариф с указанными именем и ценой
### Пример запроса:
{
    "name": "2024",
    "price": 50 
}

### Пример ответа:
"Successfully!"


