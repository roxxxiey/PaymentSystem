package main

import (
	"PaymentSystem/db"
	"PaymentSystem/internal/handlers"
	"PaymentSystem/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database := db.CreatePostgresDataBase("localhost", "5432", "postgresPayment")

	err := database.Connect("postgres", "gintas2003")
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
