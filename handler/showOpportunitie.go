package handler

import (
	"net/http"

	"github.com/costa-everton/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpportunitieHandler(ctx *gin.Context){
	id := ctx.Query("id")

	if id == "" {
		logger.Errorf("error id is required")
		sendError(ctx, errParamIsRequired("id", "queryParameter").Error(), http.StatusBadRequest)
		return
	}

	opportunitie := schemas.Opportunities{}
	err := db.First(&opportunitie, id).Error

	if err != nil {
		logger.Errorf("error id %s not found: %v",id, err.Error())
		sendError(ctx, "opportunities not found", http.StatusNotFound)
		return
	}

	sendSuccess(ctx, "show-opportunities", opportunitie)
}