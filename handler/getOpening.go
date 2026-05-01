package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/BrunoNeves1995/api-portunities/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOpiningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if strings.TrimSpace(id) == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}

	//Initializate new Struct
	opening := schemas.Opening{}

	//Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		//List is empty
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Infof("opening is not found: %v", err.Error())
			sendError(ctx, http.StatusNotFound, "opening is not found")
			return
		}
		logger.Errorf("error listing opening with id: %s, details %v", id, err.Error())
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("error listing opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "get-opening", opening)
}
