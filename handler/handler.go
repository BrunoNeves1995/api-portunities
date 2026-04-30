package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpiningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get opining",
	})
}
func GetOpiningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get opining",
	})
}
func DeleteOpiningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get opining",
	})
}
func UpdateOpiningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get opining",
	})
}
func ListOpiningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get opining",
	})
}
