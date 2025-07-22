package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"messenger-api/internal/db"
	"messenger-api/internal/metrics"
	"messenger-api/internal/rabbitmq"
	"messenger-api/internal/redisdb"
	"messenger-api/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	db.Connect()
	metrics.InitMetrics()
	redisdb.InitRedis("redis:6379", "", 0)

	go func() {
		time.Sleep(3 * time.Second)
		err := rabbitmq.PublishMessage("👋 Привет от Go-продюсера через RabbitMQ")
		if err != nil {
			log.Printf("❌ Ошибка при отправке в RabbitMQ: %v", err)
		}
	}()
	go rabbitmq.StartConsumer()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Метрики доступны на http://localhost:2112/metrics")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	app.OPTIONS("/*cors", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})
	routes.SetupRoutes(app)
	log.Fatal(app.Run(":3000"))
}