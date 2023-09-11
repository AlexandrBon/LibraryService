# LibraryService

Сервис, который храник наборы книг и их авторов и умеет отвечать
на запросы вида:
* `GetAuthorsByBookTitle`: получение авторов книги по ее названию 
* `GetBooksByAuthorName`: получение списка книг автора по его имени\

### Запуск docker-контейнеров
1. Внесение данных в `mysql.env` файл. Значение переменной 
`DB_NETWORK=tcp(db:3306)` для запуска в *docker*-контейнерах.
2. Заходим в корневую директорию и вводим: 
    ```
    make start-all
    ```
   Сервис запущен.

### Тестирование
### Postman
1. Скачать приложение можно тут: https://www.postman.com/downloads/postman-agent/
2. Запускаем контейнеры
3. Открываем приложение, выбираем `New->gRPC`. 
Выбираем вкладку `Service definition`, в ней `import .proto file` и импортируем
файл нашего сервиса: `LibraryService/internal/ports/grpc/service.proto`.
4. Далее вводим `URL` сервиса: `localhost:5700` и выбираем метод, который хотим 
проверить. 
5. Во вкладке `Message` нажимаем `Use Example Message` и корректируем под данные.

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
Тесты для сервиса находятся в `LibraryService/internal/ports/grpc/grpc_test.go`
1. Изменяем переменную окружения: `DB_NETWORK=tcp(localhost:3306)`
2. В `docker-compose.yml` в сервисе `db` комментируем `expose` и
раскомментируем `ports`
3. Запускаем `MySQL`-контейнер:
   ```
   make start-db
   ```
4. Далее, после того как `MySQL` будет готов принимать соединения, можем 
запускать/изменять/добавлять тесты и проверять их с помощью команды:
   ```
   make test
   ```
   
### Команды *make*-файла
   `build_and_run_service_locally`: запуск сервиса на своем хосте\
   `make lint`: запуск линтера

