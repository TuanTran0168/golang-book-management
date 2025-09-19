# Running the Book Management Project

This document explains how to start and run the Book Management project using Docker and execute the Go application inside the container.

---

## 1. Build and Start Containers (First Time)

Run the following command to build the Docker images and start containers:

```bash
docker compose -f 'book-management-backend\docker-compose-local.yaml' up -d --build
```

* This command builds the `book_app:dev` image and starts the app and PostgreSQL containers.
* Use `-d` to run in detached mode.

---

## 2. Start Containers (Subsequent Runs)

If containers have already been built, you can simply start them without rebuilding:

```bash
docker compose -f 'book-management-backend\docker-compose-local.yaml' up -d
```

---

## 3. Access the App Container

To enter the running container for debugging or running Go commands:

```bash
docker exec -it book_app bash
```

* This opens a shell inside the container.
* You can navigate to the app directory and run Go commands manually.

---

## 4. Run the Go Application

Inside the container shell, run the main Go application:

```bash
go run cmd/main.go
```

* This will start the Book Management application.
* Ensure that the PostgreSQL container is running and accessible (`postgres_db` service).

---

## 5. Stop Containers

To stop the running containers and remove them:

```bash
docker compose -f 'book-management-backend\docker-compose-local.yaml' down
```

* This will stop both the app and PostgreSQL containers.
* Use this when you want to completely shut down the environment.

---

## Notes

* The application reads database configuration from environment variables (`.env` or Docker env injection).
* App HTTP server will run on port `8080` by default.
* For development with hot reload, Air can be used inside the container if filesystem events are working properly, otherwise run Air on the host.
