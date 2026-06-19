# URL Shortener

Сервис сокращения URL на Go.

## Запуск

```bash
go run cmd/shortener/main.go
```

Сервер запустится на `http://localhost:8080`.

## API

### POST / — создать короткую ссылку

Принимает оригинальный URL в теле запроса как `text/plain`, возвращает сокращённый URL.

```bash
curl -X POST http://localhost:8080/ \
  -H "Content-Type: text/plain" \
  -d "https://practicum.yandex.ru/"
```

Ответ `201 Created`:
```
http://localhost:8080/EwHXdJfB
```

### GET /{id} — перейти по короткой ссылке

Перенаправляет на оригинальный URL.

```bash
curl -L http://localhost:8080/EwHXdJfB
```

Ответ `307 Temporary Redirect` с заголовком `Location: https://practicum.yandex.ru/`.

На некорректный запрос сервер возвращает `400 Bad Request`.

## Тесты

```bash
go test ./...
```

## Структура проекта

```
cmd/shortener/   — точка входа
handlers/        — HTTP-обработчики и роутер
store/           — хранилище ссылок (sync.Map)
utils/           — генерация короткого кода
```
