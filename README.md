# Book Management Project

A simple **Book Management** application built with **Golang**, using **GORM** and **PostgreSQL**.

---

## Project Structure

```
book-management/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── config.go               # Load configuration from file or environment
├── internal/
│   ├── handlers/               # HTTP controllers
│   │   ├── book_handler.go
│   │   └── author_handler.go
│   ├── middlewares/            # Middleware (auth, logging, recover)
│   │   └── auth.go
│   ├── models/                 # GORM models
│   │   ├── book.go
│   │   └── author.go
│   ├── router/                 # HTTP route definitions
│   │   └── router.go
│   ├── services/               # Business logic
│   │   ├── book_service.go
│   │   └── author_service.go
│   ├── storages/               # Repository / database layer
│   │   ├── book_storage.go
│   │   └── author_storage.go
│   └── wire/                   # Dependency injection (wire or manual)
│       └── wire.go
├── pkg/
│   └── databases/
│       └── postgres.go         # Database connection
├── .air.toml                   # Air configuration (hot reload)
├── .env.local                  # Environment configuration for development (DB, port)
├── go.mod
├── go.sum
├── Dockerfile.local            # Dockerfile for development
└── docker-compose-local.yaml   # Docker Compose for development
```

---

## Running the Project with Docker

1. Build and start containers:

```bash
docker compose -f docker-compose-local.yaml up -d --build
```

* Builds the `book_app:dev` image and starts the app + PostgreSQL containers.

2. Subsequent runs:

```bash
docker compose -f docker-compose-local.yaml up -d
```

* Starts containers without rebuilding.

3. Stop containers:

```bash
docker compose -f docker-compose-local.yaml down
```

* Stops and removes containers.

---

## Notes

* Database connection and migrations are handled automatically via GORM.
* App runs on port `8080` inside Docker, accessible at `http://localhost:8080/`.
