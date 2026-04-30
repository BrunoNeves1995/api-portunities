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
