# PeopleHub API

A production-ready RESTful API built in Go for managing people records using PostgreSQL and Clean Architecture principles.

## 🔧 Features

- ✅ PostgreSQL Integration via sqlx
- ✅ Environment configuration with `.env`
- ✅ RESTful endpoints for people management
- ✅ Clean layered architecture (`models`, `repository`, `handlers`, `utils`)
- ✅ Full test coverage using `httptest` and Go testing
- ✅ Ready for Docker & Auth integration (WIP)

## 📁 Project Structure

```
peoplehub-api/
├── go.mod
├── main.go
├── .env
├── schema.sql
├── README.md
├── models/
│   └── person.go
├── repository/
│   └── person_repository.go
├── utils/
│   └── db.go
├── handlers/
│   └── person_handlers.go
├── tests/
│   └── main_test.go
```

## 📦 Tech Stack

- Go 1.20+
- PostgreSQL
- `sqlx`, `pq`, `godotenv`
- Clean Architecture
- GitHub + Version control
- Postman + JSON API
- Full HTTP integration tests

## 🚀 Endpoints

| Method | Route              | Description             |
|--------|--------------------|-------------------------|
| GET    | /people            | List all people         |
| POST   | /people            | Add new person          |
| DELETE | /people/{id}       | Delete person by ID     |
| PATCH  | /people/{id}       | Update person fields    |
| GET    | /people/search     | Search by name          |
| GET    | /people/stats      | Show total count        |

## 📚 Setup & Run

1. Clone the repo  
2. Configure `.env` file:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=person_api
```

3. Run SQL schema:
```bash
psql -U postgres -d person_api -f schema.sql
```

4. Install & run:
```bash
go mod tidy
go run main.go
```

---

## 🧪 Run Tests

```bash
go test ./tests -v
```

---

## 👨‍💻 Author

Built by [@Diony-source](https://github.com/Diony-source) — passionate backend developer in training.  
This project is part of a 30+ day backend mastery roadmap in Go.
