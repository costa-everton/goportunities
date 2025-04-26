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

func ShowOpportunitiesHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opportunities opening",
	})
}

func ListOpportunitiesHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET list opportunities",
	})
}

func UpdateOpportunitiesHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "PUT opening updeted",
	})
}

func DeleteOpportunitiesHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"DELETE opportunities",			
	})
}