package router

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	// Confia apenas em localhost
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal(err)
	}
	// Configuração básica
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:8000"}, // frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	initializeRoutes(r)

	r.Run(":8000")
}
