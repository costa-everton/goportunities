package handler

import (
	"net/http"

	"github.com/costa-everton/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpportunitiesHandler(ctx *gin.Context){
	request := CreateOpportunitiesRequest{}

	ctx.BindJSON(&request)
	if err:= request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	opportunities := schemas.Opportunities{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	err := db.Create(&opportunities).Error
	if err != nil {
		logger.Errorf("error creating opportunities: %v", err.Error())
		sendError(ctx, "error creating opportunities on database", http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "create-opportuninies", opportunities)
}