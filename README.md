# JWT auth server
## Тестовое задание от [Мельник Руслана Сергеевича](https://surgut.hh.ru/resume/f983c45cff0bee08e50039ed1f723673737344) 

При работе необходимы данные конфигурации `.env`. Его cтруктура описана в файле [.env.example](./.env.example).
Также необходима структура таблиц БД. Она описана в файле [setup.sql](./setup.sql).

Приложение реализуют следующую функциональность:
- `HTTP [GET] /{guid}` - создание и сохраннеия в cookies пары токенов 
- `HTTP [GET] /refresh` - обработка пары токенов и последующее их пересоздание с сохранением в cookies

Более подробное описание работы в [файле architecture.md](./architecture.md) 
