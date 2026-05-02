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
		// opening: vaga
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	//Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
