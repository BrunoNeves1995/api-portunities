package router

import (
	"github.com/BrunoNeves1995/api-portunities/docs"
	"github.com/BrunoNeves1995/api-portunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	// Initialize Swagger info
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Initialize router
	v1 := router.Group(basePath)
	{
		//opining: vaga
		v1.GET("/opining", handler.GetOpiningHandler)
		v1.POST("/opining", handler.CreateOpiningHandler)
		v1.DELETE("/opining", handler.DeleteOpiningHandler)
		v1.PUT("/opining", handler.UpdateOpiningHandler)
		v1.GET("/opinings", handler.ListOpiningsHandler)
	}

	//Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
