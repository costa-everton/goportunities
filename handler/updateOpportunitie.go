package handler

import (
	"net/http"

	"github.com/costa-everton/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpportunitieHandler(ctx *gin.Context){
	request := UpdateOpportunitiesRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	opportunities := schemas.Opportunities{}

	if err := db.First(&opportunities, request.ID).Error; err != nil {
		logger.Errorf("error id %s not found: %v",request.ID, err.Error())
		sendError(ctx, "error not fount", http.StatusNotFound)
		return
	}

	if request.Role != "" {
		opportunities.Role = request.Role
	}
	if request.Company != "" {
		opportunities.Company = request.Company
	}
	if request.Location != "" {
		opportunities.Location = request.Location
	}
	if request.Remote != nil {
		opportunities.Remote = *request.Remote
	}
	if request.Link != "" {
		opportunities.Link = request.Link
	}
	if request.Salary > 0 {
		opportunities.Salary = request.Salary
	}

	if err := db.Save(&opportunities).Error; err != nil {
		logger.Errorf("error updatining opportunitie %v", err.Error())
		sendError(ctx, "error updatining opportunitie", http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "update-opportuninies", opportunities)
}