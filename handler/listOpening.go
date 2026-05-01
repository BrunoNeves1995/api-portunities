package handler

import (
	"net/http"

	"github.com/BrunoNeves1995/api-portunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpiningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	//Find all openings
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
