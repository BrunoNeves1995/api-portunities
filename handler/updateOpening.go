package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpiningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get opining",
	})
}
