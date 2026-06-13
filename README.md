# User Age API

A RESTful API built with **GoFiber**, **PostgreSQL**, and **SQLC** to manage users and dynamically calculate their age from their date of birth.

---

## Features

* Create a user with `name` and `dob`
* Retrieve a user by ID with dynamically calculated age
* List all users with age
* Update user details
* Delete users
* PostgreSQL database integration
* SQLC for type-safe database queries
* Input validation using `go-playground/validator`
* Logging using Uber Zap
* Request ID middleware
* Request duration logging middleware
* Docker support
* Unit tests for age calculation

---

## Tech Stack

| Technology              | Purpose                    |
| ----------------------- | -------------------------- |
| Go                      | Programming Language       |
| GoFiber                 | Web Framework              |
| PostgreSQL              | Database                   |
| SQLC                    | Type-safe query generation |
| pgx/v5                  | PostgreSQL Driver          |
| Uber Zap                | Logging                    |
| go-playground/validator | Input Validation           |
| Docker                  | Containerization           |

---

## Project Structure

```text
.
├── cmd
│   └── server
│       └── main.go

├── config
│   ├── config.go
│   └── database.go

├── db
│   ├── migrations
│   │   ├── 000001_create_users_table.up.sql
│   │   └── 000001_create_users_table.down.sql
│   │
│   ├── query
│   │   └── users.sql
│   │
│   ├── sqlc
│   │   ├── db.go
│   │   ├── models.go
│   │   └── users.sql.go
│   │
│   └── sqlc.yaml

├── internal
│   ├── handler
│   │   └── user_handler.go
│   │
│   ├── logger
│   │   └── zap.go
│   │
│   ├── middleware
│   │   ├── logger.go
│   │   └── request_id.go
│   │
│   ├── models
│   │   ├── request.go
│   │   └── response.go
│   │
│   ├── repository
│   │   └── user_repository.go
│   │
│   ├── routes
│   │   └── routes.go
│   │
│   └── service
│       ├── age_service.go
│       └── user_service.go

├── tests
│   └── age_test.go

├── Dockerfile
├── docker-compose.yml
├── .env
├── go.mod
├── go.sum
└── README.md
```

---

## Database Schema

### users

| Column | Type   | Constraints |
| ------ | ------ | ----------- |
| id     | SERIAL | PRIMARY KEY |
| name   | TEXT   | NOT NULL    |
| dob    | DATE   | NOT NULL    |

SQL:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

---

## Dynamic Age Calculation

The `age` field is **not stored** in the database.

It is calculated dynamically using Go's `time` package whenever a user is fetched.

Example:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 36
}
```

---

# API Endpoints

## Create User

### Request

```http
POST /users
```

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### Response

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

Status:

```text
201 Created
```

---

## Get User By ID

### Request

```http
GET /users/1
```

### Response

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 36
}
```

Status:

```text
200 OK
```

---

## List Users

### Request

```http
GET /users
```

### Response

```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 36
  }
]
```

Status:

```text
200 OK
```

---

## Update User

### Request

```http
PUT /users/1
```

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

### Response

```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

Status:

```text
200 OK
```

---

## Delete User

### Request

```http
DELETE /users/1
```

### Response

```text
204 No Content
```

---

# Local Setup

## Prerequisites

* Go 1.26+
* PostgreSQL
* SQLC
* Docker (optional)

---

## Clone Repository

```bash
git clone <your-repository-url>

cd user-age-api
```

---

## Configure Environment Variables

Create `.env`

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=user_age_db
DB_SSLMODE=disable
```

---

## Run Database Migration

```bash
migrate -path db/migrations \
-database "postgres://postgres:1234@localhost:5432/user_age_db?sslmode=disable" up
```

---

## Generate SQLC Code

```bash
sqlc generate -f db/sqlc.yaml
```

---

## Run the Application

```bash
go run cmd/server/main.go
```

Server runs on:

```text
http://localhost:8080
```

---

# Docker Setup

## Build Image

```bash
docker build -t user-age-api .
```

---

## Run Container

```bash
docker run -p 8080:8080 --env-file .env user-age-api
```

---

# Running Tests

Run all tests:

```bash
go test ./...
```

Example output:

```text
ok github.com/ShasiChowdam/user-age-api/tests
```

---

# Middleware

### Request ID Middleware

Adds:

```text
X-Request-ID
```

to every response.

---

### Request Duration Logging

Logs request method, path, and duration.

Example:

```text
GET /users - 23ms

POST /users - 8ms

DELETE /users/2 - 3ms
```

---

# Logging

The application uses **Uber Zap** for structured logging.

Example:

```text
Database connected successfully

Server started
```

---

# Future Improvements

* Pagination for `/users`
* Integration tests
* Swagger/OpenAPI documentation
* Authentication and Authorization
* CI/CD pipeline with GitHub Actions

---

# Author

**Sasi Chowdam**

Built as part of a backend development assignment using Go, PostgreSQL, SQLC, and Fiber.
