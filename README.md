# Book Management Project

A simple **Book Management** application built with **Golang**, using **GORM** and **PostgreSQL**.

---

## Project Structure
```
├── 📁 .git/ 🚫 (auto-hidden)          # Git metadata, do not touch
├── 📁 book-management-backend/        # Main project source
│   ├── 📁 cmd/                        # Application entry point
│   │   └── 🐹 main.go                 # Main file, starts the server
│   ├── 📁 configs/                    # Configuration management, loads from .env
│   │   └── 🐹 config_env.go
│   ├── 📁 docs/                       # Auto-generated Swagger docs (do not edit manually)
│   │   ├── 🐹 docs.go
│   │   ├── 📄 swagger.json
│   │   └── ⚙️ swagger.yaml
│   ├── 📁 internal/                   # Business logic (Clean Architecture)
│   │   ├── 📁 handlers/               # Controllers: handle requests → call services
│   │   │   └── 🐹 author_handler.go
│   │   ├── 📁 middlewares/            # Middleware (auth, logging, CORS, etc.)
│   │   ├── 📁 models/                 # Entities / structs mapping to DB
│   │   │   ├── 🐹 author.go
│   │   │   └── 🐹 book.go
│   │   ├── 📁 repositories/           # Repository layer: DB queries
│   │   │   └── 🐹 author_repository.go
│   │   ├── 📁 routers/                # HTTP route definitions
│   │   │   ├── 🐹 author_routes.go
│   │   │   └── 🐹 router.go
│   │   ├── 📁 services/               # Service layer: business logic
│   │   │   └── 🐹 author_service.go
│   │   └── 📁 wire/                   # Dependency injection (Google Wire / manual DI)
│   ├── 📁 notes/                      # Development notes (internal docs)
│   │   ├── 📝 RUN.md                  # How to run the project
│   │   ├── 📄 init.txt
│   │   ├── 📄 lib.txt
│   │   ├── 📄 run.txt
│   │   └── 📄 structure.txt
│   ├── 📁 pkg/                        # Reusable packages (utils, db, etc.)
│   │   └── 📁 databases/
│   │       └── 🐹 postgresql.go       # PostgreSQL connection & migrations
│   ├── 📁 tmp/ 🚫 (auto-hidden)       # Temporary files (e.g., from Air hot reload)
│   ├── ⚙️ .air.toml                   # Air configuration (hot reload)
│   ├── 🔒 .env 🚫 (auto-hidden)       # Environment file (production/secret)
│   ├── 📄 .env.local 🚫 (auto-hidden) # Local environment file (development)
│   ├── 📄 DockerFile.local            # Dockerfile for local development
│   ├── ⚙️ docker-compose-local.yaml   # Docker Compose (Go + PostgreSQL)
│   ├── 🐹 go.mod                      # Go module definition
│   └── 🐹 go.sum                      # Dependency checksums
├── 🚫 .gitignore                      # Files ignored by Git
└── 📖 README.md                       # Main documentation (project overview)
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
