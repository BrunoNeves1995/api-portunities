package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	FullPath  string `json:"fullPath"`
	URI       string `json:"uri"`
	Timestamp string `json:"timestamp"`
}

func sendError(ctx *gin.Context, code int, msg string) {
	resp := ErrorResponse{
		Status:    code,
		Message:   msg,
		FullPath:  ctx.FullPath(),
		URI:       ctx.Request.RequestURI,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	ctx.JSON(code, resp)
}

type SuccessCreateOpeningResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	success := SuccessCreateOpeningResponse{
		Message: fmt.Sprintf("Operation from handler: %s successfull", operation),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, success)
}
