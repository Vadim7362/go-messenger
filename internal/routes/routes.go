package routes

import (
	"messenger-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Привет! Сервер работает через Gin.")
    })

    r.POST("/register", handlers.Register)
		r.POST("/login", handlers.Login)
}