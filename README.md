## Описание
___

Приложение, которое при запросе на него POST /status опросит все переданные сервисы, получит их статусы и если
статусы этих сервисов равны 200 (статус ОК), то выдаст сообщение "Всё работает".

Если любой сервис не отдаст статус за 3 секунды или статус не будет равен 200, то выдаст сообщение "Сервис {url} не
работает"

Пример запроса:

```
POST http://localhost:8080/status

{
    "adr": [
        "http://192.168.88.254:8088/status",
        "http://192.168.88.254:8089/status"
    ]
}
```