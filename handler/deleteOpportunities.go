package handler

import (
	"fmt"
	"net/http"

	"github.com/costa-everton/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpportunitiesHandler(ctx *gin.Context){
	id := ctx.Query("id")
	if id == "" {
		logger.Errorf("error id is required")
		sendError(ctx, errParamIsRequired("id", "queryParameter").Error(), http.StatusBadRequest)
		return
	}

	opportunities := schemas.Opportunities{}

	err := db.First(&opportunities, id).Error
	if err != nil {
		logger.Errorf("error id %s not found: %v",id, err.Error())
		sendError(ctx, fmt.Sprintf("opportunities with id %s not found", id), http.StatusNotFound)
		return
	}

	if err := db.Delete(&opportunities).Error; err != nil{
		logger.Errorf("error opportunities not deleted: %v", err.Error())
		sendError(ctx, fmt.Sprintf("error deleting opportunities with id: %s", id), http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "delete-opportunities", opportunities)
}