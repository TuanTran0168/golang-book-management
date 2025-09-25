package databases

import (
	"fmt"
	"log"
	"os"

	config "book-management/configs"
	"book-management/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	// Global logger
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // output to console
		logger.Config{
			// SlowThreshold:             time.Second, // log slow queries
			LogLevel:                  logger.Info, // log all queries
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	log.Println("✅ Database connected")

	// Auto migrate models
	if err := db.AutoMigrate(
		&models.Author{},
		&models.Book{},
		&models.User{},
		&models.Genre{},
	); err != nil {
		return nil, err
	}
	return db, nil
}
