# LibraryService

Сервис, который хранит наборы книг и их авторов и умеет отвечать
на запросы вида:
* `GetAuthorsByBookTitle`: получение авторов книги по ее названию 
* `GetBooksByAuthorName`: получение списка книг автора по его имени

### БД
Скрипты для создания таблиц и тестовые данные расположены в `mysql/scripts`. 
Директория `mysql/storage` смонтированна в `/var/lib/mysql`.

### Запуск docker-контейнеров
1. Необходимо внести данные в `mysql.env` файл. Значение переменной 
`DB_NETWORK=tcp(db:3306)`
2. Заходим в корневую директорию и вводим: 
    ```
    make start-all
    ```
   Сервис и БД запущены

## Тестирование
### С помощью Postman
Скачать приложение можно тут: https://www.postman.com/downloads/postman-agent/
1. Запускаем *docker*-контейнеры
2. Открываем `Postman`, выбираем `New -> gRPC`.
Выбираем вкладку `Service definition`, в ней `import .proto file` и импортируем
файл нашего сервиса: `internal/ports/grpc/service.proto`
3. Далее вводим `URL` сервиса: `localhost:5700` и выбираем метод, который хотим 
проверить
4. Во вкладке `Message` нажимаем `Use Example Message` и корректируем для тестов

   Пример запроса:
   ```
   {
       "author_name": "a_1"
   }
   ```
   
   Пример ответа:
   ```
   {
       "books": [
           {
               "id": 1,
               "title": "t_1",
               "page_count": 40,
               "publishing_year": 1988
           }
       ]
   }
   ```

### С помощью кода
Тесты для сервиса и БД находятся в `internal/ports/grpc/grpc_test.go`
1. Изменяем переменную окружения `DB_NETWORK=tcp(localhost:3306)` в файле `mysql.env`
2. В `docker-compose.yml` в сервисе `db` комментируем `expose` часть и
убираем комментарии с `ports` части
3. Запускаем `MySQL`-контейнер:
   ```
   make start-db
   ```
4. Далее, после того как `MySQL` сервер будет готов принимать соединения, можем 
запускать тесты:
   ```
   make test
   ```
### Заметки к командам *make*-файла
   `make build`: сборка сервиса на хосте.\
   `make run`: запуск сервиса на хосте. Перед выполнением команды необходимо запустить `MySQL` контейнер \
   `make lint`: запуск линтера\
   `make test`: если не запускать БД, то все тесты кроме `internal/ports/grpc/grpc_test.go` должны выполниться успешно
