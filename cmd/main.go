package main

import (
	"log"

	"github.com/gin-contrib/cors"
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
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.SetupRoutes(app)
	log.Fatal(app.Run(":3000"))
}