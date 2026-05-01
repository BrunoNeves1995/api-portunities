package router

import (
	"github.com/BrunoNeves1995/api-portunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()

	// Initialize router
	v1 := router.Group("/api/v1")
	{
		//opining: vaga
		v1.GET("/opining", handler.GetOpiningHandler)
		v1.POST("/opining", handler.CreateOpiningHandler)
		v1.DELETE("/opining", handler.DeleteOpiningHandler)
		v1.PUT("/opining", handler.UpdateOpiningHandler)
		v1.GET("/opinings", handler.ListOpiningsHandler)
	}
}
