package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize router
	v1 := router.Group("/api/v1")
	{
		//opining: vaga
		v1.GET("/opining", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "Get opining",
			})
		})

		v1.POST("/opining", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "POST opining",
			})
		})

		v1.DELETE("/opining", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "DELETE opining",
			})
		})

		v1.PUT("/opining", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "PUT opining",
			})
		})

		v1.GET("/opinings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "Get opining",
			})
		})
	}
}
