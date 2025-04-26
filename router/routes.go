package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func initializeRoutes(router *gin.Engine){

	v1 := router.Group( "/api/v1")
	{
		v1.POST("/create-opportunities", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "POST opportunities create",
			})
		})
		v1.GET( "/opening-opportunities", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET opportunities opening",
			})
		})

		v1.PUT("/update-opportunities", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message" : "PUT opening updeted",
			})
		})

		v1.DELETE("/delete-opportunities", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message":"DELETE opportunities",			
			})
		})
	}
}