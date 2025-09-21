# 📚 Book Management Backend

**Book Management Backend** is a **RESTful API** built with **Golang** following **Clean Architecture**.  
It provides book, author, and user management with **JWT authentication**, uses **PostgreSQL** for data storage,  
supports **Cloudinary** for image uploads, and includes **Swagger** for API documentation.

---

## 🚀 Demo & Swagger UI

You can see the live deployed version of the API and explore its endpoints on Swagger UI:

📖 [Book Management Backend Swagger API](https://golang-book-management-h5pt.onrender.com/swagger/index.html)

Here you’ll find full API documentation, try out endpoints, check request/response schemas, etc.

---

## Project Structure
```
├── 📁 .git/ 🚫 (auto-hidden)              # Git metadata, do not touch
├── 📁 book-management-backend/            # Main project source
│   ├── 📁 cmd/                            # Application entry point
│   │   └── 🔵 main.go                     # Main file, starts the server
│   ├── 📁 configs/                        # Configuration management, loads from .env
│   │   └── 🔵 config_env.go               # Loads env variables and configs
│   ├── 📁 docs/                           # Auto-generated Swagger docs (do not edit manually)
│   │   ├── 🔵 docs.go
│   │   ├── 📄 swagger.json
│   │   └── ⚙️ swagger.yaml
│   ├── 📁 internal/                       # Business logic (Clean Architecture)
│   │   ├── 📁 handlers/                   # Controllers: handle requests → call services
│   │   │   ├── 🔵 auth_handler.go
│   │   │   ├── 🔵 author_handler.go
│   │   │   └── 🔵 book_handler.go
│   │   ├── 📁 middlewares/                # Middleware (auth, logging, CORS, etc.)
│   │   │   ├── 🔵 auth_middleware.go
│   │   │   ├── 🔵 cors_middleware.go
│   │   │   └── 🔵 ip_middleware.go
│   │   ├── 📁 models/                     # Entities / structs mapping to DB
│   │   │   ├── 🔵 author.go
│   │   │   ├── 🔵 book.go
│   │   │   └── 🔵 user.go
│   │   ├── 📁 repositories/               # Repository layer: DB queries
│   │   │   ├── 🔵 author_repository.go
│   │   │   ├── 🔵 book_repository.go
│   │   │   └── 🔵 user_repository.go
│   │   ├── 📁 routers/                    # HTTP route definitions
│   │   │   ├── 🔵 auth_routes.go
│   │   │   ├── 🔵 author_routes.go
│   │   │   ├── 🔵 book_routes.go
│   │   │   └── 🔵 router.go
│   │   ├── 📁 services/                   # Service layer: business logic
│   │   │   ├── 🔵 author_service.go
│   │   │   ├── 🔵 book_service.go
│   │   │   └── 🔵 user_service.go
│   │   └── 📁 wire/                       # Dependency injection (Google Wire / manual DI)
│   ├── 📁 notes/                          # Development notes (internal docs)
│   │   ├── 📝 RUN.md                      # How to run the project
│   │   ├── 📄 init.txt
│   │   ├── 📄 lib.txt
│   │   ├── 📄 run.txt
│   │   └── 📄 structure.txt
│   ├── 📁 pkg/                            # Reusable packages (utils, db, etc.)
│   │   ├── 📁 databases/
│   │   │   └── 🔵 postgresql.go           # PostgreSQL connection & migrations
│   │   └── 📁 utils/
│   │       ├── 🔵 cloudinary.go           # Cloudinary image upload helper
│   │       └── 🔵 jwt.go                  # JWT helper functions
│   ├── 📁 tmp/ 🚫 (auto-hidden)           # Temporary files (e.g., from Air hot reload)
│   ├── ⚙️ .air.toml                       # Air configuration (hot reload)
│   ├── 🔒 .env 🚫 (auto-hidden)           # Environment file (production/secret)
│   ├── 📄 .env.local 🚫 (auto-hidden)     # Local environment file (development)
│   ├── 📄 DockerFile.local                # Dockerfile for local development
│   ├── ⚙️ docker-compose-local.yaml       # Docker Compose (Go + PostgreSQL)
│   ├── 🔵 go.mod                          # Go module definition
│   └── 🔵 go.sum                          # Dependency checksums
├── 🚫 .gitignore                          # Files ignored by Git
└── 📖 README.md                           # Main documentation (project overview)

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

---

### 🔧 Tech Stack
- **Golang** 🟦
- **Gin** ⚡ (HTTP web framework)
- **GORM** 📦 (ORM for database)
- **PostgreSQL** 🐘
- **JWT Authentication** 🔑
- **Cloudinary** ☁️ (image upload)
- **Swagger** 📑 (API docs)
- **Docker** 🐳 (containerization)