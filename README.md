# Микросервис библиотеки песен

## О проекте

Тестовое задание на тему "Микросервис библиотеки песен".

Дополнительные инструменты:

- Документация Swagger доступна по маршруту `http://localhost:8081/swagger/index.html`

## Локальный запуск

В проекте присутствуют taskfile скрипты: в корневой директории и в директории микросервиса. Для локального запуска необходимо выполнить следующие шаги:

1. Запуск всего окружения микросервиса - PostgreSQL.

```shell
task docker:compose-local
```

2. Запуск микросервиса.

```shell
сd service
task run:local
```
