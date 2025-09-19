package main

import (
	"log"

	configs "book-management/configs"
	"book-management/internal/handlers"
	"book-management/internal/repositories"
	router "book-management/internal/routers"
	"book-management/internal/services"
	database "book-management/pkg/databases"
)

func main() {
	// 1. Load config
	cfg := configs.LoadConfig()

	// 2. Connect DB
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
	}
	log.Println("✅ Database connected and migrated successfully")

	// 3. Init repository, service, handler
	authorRepo := repositories.NewAuthorRepository()
	authorService := services.NewAuthorService(authorRepo, db)
	authorHandler := handlers.NewAuthorHandler(authorService)

	// 4. Setup Gin router
	server := router.NewRouter(authorHandler)

	// 5. Start server
	port := cfg.HTTPPort
	log.Printf("🚀 Server running on http://localhost:%s\n", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
