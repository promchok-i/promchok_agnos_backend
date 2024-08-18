package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/promchok-i/promchok_agnos_backend/middleware"
	"github.com/promchok-i/promchok_agnos_backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate your models
	err = database.AutoMigrate(&models.RequestLog{})
	if err != nil {
		return nil, err
	}

	return database, nil
}

func main() {
	var err error
	db, err := initDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	log.Printf("Connect to Database Successfully: %v", db)

	r := gin.Default()
	r.Use(middleware.LogRequestResponseMiddleware(db))

	// Group routes under /api
	api := r.Group("/api")
	{
		api.POST("/strong_password_steps", checkPasswordHandler)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
