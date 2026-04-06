# Tool-and-materials-accounting-system
A student's coursework on the topic "Tool and materials accounting system"

**Процесс развертывания **

 Настройка базы данных

Установите **PostgreSQL**
Создайте базу данных (например: `inventory`)
Выполните SQL-скрипт для создания таблиц:

```
psql -U postgres -d inventory -f migration.sql
```

Настройка backend

Отредактируйте подключение к БД в файле:

```
internal/database/db.go
```

Укажите свои данные:

``` (golang)
connStr := "user=postgres password=YOUR_PASSWORD dbname=inventory sslmode=disable"
```

4. Запуск backend

```
go mod tidy
go run cmd/main.go
```

Сервер запустится на:

```
http://localhost:8080
```

---

5. Запуск frontend

Откройте файл:

```
frontend/index.html
```

ИЛИ запустите локальный сервер:

```
python -m http.server
```


!!! Требования !!!

* Go 1.20+ (Фреймворк Gin актуальной версии )
* PostgreSQL (PgAdmin4)
