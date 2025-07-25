package main

import (
	"PaymentSystem/db"
	"PaymentSystem/internal/handlers"
	"PaymentSystem/router"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	database := db.CreatePostgresDataBase(dbHost, dbPort, dbName)

	err := database.Connect(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		panic(err)
	}

	handlers.SetDB(database.GetDB())

	err = handlers.AutoMigrateAndInit()
	if err != nil {
		panic(err)
	}

	r := router.NewRouter(router.Gin)
	r.SetupRoutes()

	engine := r.GetHandler().(*gin.Engine)
	log.Println("Сервер запущен на :8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
