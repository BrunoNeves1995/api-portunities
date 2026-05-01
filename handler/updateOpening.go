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

func UpdateOpiningHandler(ctx *gin.Context) {
	request := UpdateOpiningrequest{}

	//Validation JSON Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("JSON invalid: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//Validation
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if strings.TrimSpace(id) == "" {
		logger.Errorf("validation error: %s queryParameter is required", id)
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

	//Update Opening
	applyOpeningUpdate(&opening, request)

	//Save Opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	//Return successfull
	sendSuccess(ctx, "update-opening", opening)

}
