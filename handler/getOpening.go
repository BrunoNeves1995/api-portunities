package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpiningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get opining",
	})
}
