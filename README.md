### О сервисе

Сервис, который получает по апи ФИО, из открытых апи обогащает
ответ наиболее вероятными возрастом, полом и национальностью и сохраняет данные в
БД. По запросу выдает инфу о найденных людях.

### Деплой

Сначала создать в корне файл .env (пример .env_example) и настроить config.yaml 

Docker-compose команды:
```
docker-compose build
docker-compose up
```

### Ручки

Добавить человека
```
POST /api/person 
body: json
{
    "name":"Valeriy",
    "surname":"Korolev",
    "patronymic":"Pavlovich" // не обязательный
}
```
Изменить человека
```
PUT /api/person/:id
id - int
body: json
{
    "name": "Valeriy",
    "surname": "Korolev",
    "patronymic": {
        "String": "Pavlovich",
        "Valid": true
    },// не обязательный параметр
    "age": 23,
    "gender": "male",
    "nationality": "RU"
}
```
Получить человека
```
GET /api/person/:id
id - int
```
Удалить человека
```
DELETE /api/person/:id
id - int
```
Список людей с пагинацией
```
GET /api/person?page=0&limit=20
page, limit - int
```

### Postman collection

Примеры запросов находятся в постман коллекции в файле effective_mobile_task.postman_collection.json

### _Сделано:_
1. Развернута бд и настроен деплой в Docker-compose
2. Добавил переменные окружения и конфиг
3. Добавил миграции (goose)
4. Миграции накатываются при деплое с помощью Docker
5. REST CRUD с обогащением по внешним API:
   * Подключение к базе ранных
   * crud в repository для бд:
      * добавление
      * изменение
      * удаление
      * чтение
   * получение данных по внешним api
   * http ручки:
      * добавление
      * изменение
      * удаление
      * чтение
6. Пагинация (limit и page по умолчанию 20 и 0 соответственно, максимум лимит: 100, минимум: 1)
     

### _Осталось сделать_:
1. Debug и info логи

### _Было бы круто сделать_:
1. Написать тесты
2. Swagger
