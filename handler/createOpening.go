package handler

import (
	"net/http"

	"github.com/BrunoNeves1995/api-portunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpiningrequest true "Request body"
// @Success 200 {object} SuccessCreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpiningHandler(ctx *gin.Context) {
	request := CreateOpiningrequest{}

	//Validation JSON Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// Retorna erro de validação
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

	//Initializate new Struct
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	//Return successfull
	sendSuccess(ctx, "create-opening", opening)

}
