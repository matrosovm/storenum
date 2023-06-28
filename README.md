# storenum

Сервис по взаимодействию с целыми числами.

Запуск сервиса осуществляется через ```make run```, 

линтер - ```make lint/format```, 

зависимости - ```make bin-deps```.

Сервис предоставляет следующее API:
* [получение числа](#получение-числа)
* [добавление числа](#добавление-числа)
* [установка значения](#установка-значения)

Сваггер доступен по http://localhost:8080/swagger/

## Получение числа
Пример запроса:
```
curl -X 'GET' \
  'http://localhost:8080/current_number' \
  -H 'accept: application/json'
```

## Добавление числа
На вход подается натуральное число, которое прибавляется к текущему. 

Пример запроса:
```
curl -X 'POST' \
  'http://localhost:8080/add_number' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "number": "12"
}'
```

## Установка значения
На вход подается натуральное число, которое заменяет текущее.

Пример запроса:
```
curl -X 'POST' \
  'http://localhost:8080/set_number' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "number": "1"
}'
```