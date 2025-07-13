package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"messenger-api/internal/db"
	"messenger-api/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	db.Connect()
	router := gin.Default()
	routes.SetupRoutes(router)
	log.Fatal(router.Run(":3000"))
}