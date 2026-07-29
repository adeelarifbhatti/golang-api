# Golang REST API

A simple RESTful CRUD API built with Go and MySQL, containerized with Docker. It manages a `languages` resource and exposes standard HTTP endpoints for create, read, update, and delete operations.

## Tech Stack

- **Language:** Go 1.21
- **Router:** [gorilla/mux](https://github.com/gorilla/mux)
- **Database:** MySQL
- **Containerization:** Docker & Docker Compose
- **CI:** GitHub Actions

## Project Structure

```
.
├── main.go                        # Entry point
├── app.go                         # App struct, router setup, HTTP handlers
├── model.go                       # Database queries and language model
├── constants.go                   # DB credentials and name
├── app_test.go                    # Integration tests
├── Dockerfile                     # Production image
├── DockerfileTest                 # Test image
├── docker-compose.yml             # Local development stack
├── docker-compose-ci.yml          # CI stack
├── init-script.sql                # DB initialization (table + seed data)
└── .github/workflows/ci.yaml      # GitHub Actions CI pipeline
```

## API Endpoints

| Method | Path             | Description           |
|--------|------------------|-----------------------|
| GET    | `/languages`     | List all languages    |
| GET    | `/language/{id}` | Get a language by ID  |
| POST   | `/language`      | Create a new language |
| PUT    | `/language/{id}` | Update a language     |
| DELETE | `/language/{id}` | Delete a language     |

### Example Request Body (POST / PUT)

```json
{
  "name": "Rust"
}
```

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/)

### Run with Docker Compose

```bash
docker-compose up --build
```

This starts two services:
- **mysql** — MySQL database on port `3306`, seeded with `init-script.sql`
- **golang-api** — Go API server on port `8080`

The API will be available at `http://localhost:8080`.

### Run Locally (without Docker)

Ensure a MySQL instance is running and accessible at `mysql:3306` (or update the connection string in `app.go`), then:

```bash
go build -o golang-api
./golang-api
```

## Running Tests

Tests require a running MySQL instance. Using Docker Compose:

```bash
docker-compose -f docker-compose-ci.yml up --build
docker exec golang-api_golang-api_1 go test
```

Or, if MySQL is available locally:

```bash
go test
```

## CI Pipeline

GitHub Actions runs on every push:

1. Starts MySQL and seeds the database
2. Builds the Go binary (`go build`)
3. Runs the test suite (`go test`)

See [.github/workflows/ci.yaml](.github/workflows/ci.yaml) for details.

## Database

The `languages` table is created and seeded automatically via `init-script.sql`:

```sql
CREATE TABLE `languages` (
  `id`   INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255)
);
-- Seeded with: Java, python, golang
```

## Configuration

Database credentials are defined in `constants.go`:

| Constant     | Default    |
|--------------|------------|
| `DBName`     | `language` |
| `DbUser`     | `dbuser`   |
| `DbPassword` | `changeme` |

> **Note:** For production use, replace hardcoded credentials with environment variables or a secrets manager.
