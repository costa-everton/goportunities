package handler

import (
	"net/http"

	"github.com/costa-everton/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpportunitiesHandler(ctx *gin.Context){
	opportunities := []schemas.Opportunities{}

	err := db.Find(&opportunities).Error

	if  err != nil {
		logger.Errorf("error listing opportunities: %v", err.Error())
		sendError(ctx, "error listing opportunities", http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "list-oportunities", opportunities)
}