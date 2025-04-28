package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpportunitiesHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "PUT opening updeted",
	})
}