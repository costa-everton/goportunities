package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpportunitiesHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST opportunities create",
	})
}