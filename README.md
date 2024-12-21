# CalcServer

### Сервер подсчета арифметических выражений на Golang.
### Данный сервер решает арифметические действия с однозначными положительными числами, также учитывая скобки и другие арифметические знаки.
### Сервер проверяет выражение на такие ошибки: некорректное выражение, пустое выражение.
#
Для запуска сервера выполнить команду в папке проекта:

```
go run ./cmd/main.go
```

при старте сервера в терминале выdодится строка "Server started"
#
Для проверки работы сервера выполнить запрос:

```
curl --location --request GET ' http://localhost:8080 ' --header 'Content-Type: application/json' --data '\{"expression": "3\+(8\*3)"\}'
```

При правильном запросе выдается код ответа 200 OK — успешный запрос
#
Пример некорректного запроса (отсуствует body)

```
curl --location ' http://localhost:8080 '
```

или неправильно выражение

```
curl --location --request GET ' http://localhost:8080 ' --header 'Content-Type: application/json' --data '\{"expression": "33"\}'
```

выводит код ошибки 400 Bad Request
#
Пустое выражение выводит код ошибки 422 Unprocessable Entity:

```
curl --location --request GET ' http://localhost:8080 ' --header 'Content-Type: application/json' --data '\{"expression": ""\}'
```
#
Для запуска тестов http сервера выполните команду в папке проекта:

```
go test ./internal/application
```
#
Для запуска тестов подсчета арифметических выражений выполните команду:

```
go test ./pkg/calculation
```
