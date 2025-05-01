package router

import (
	"github.com/costa-everton/goportunities/handler"
	"github.com/gin-gonic/gin"
)


func initializeRoutes(router *gin.Engine){
	// initialize handler
	handler.InitializeHandler()

	v1 := router.Group( "/api/v1")
	{
		v1.POST("/create-opportunities", handler.CreateOpportunitiesHandler)
		v1.GET( "/opening-opportunities", handler.ShowOpportunitiesHandler)
		v1.GET( "/list-opportunities", handler.ListOpportunitiesHandler)
		v1.PUT("/update-opportunities", handler.UpdateOpportunitiesHandler)
		v1.DELETE("/delete-opportunities", handler.DeleteOpportunitiesHandler)
	}
}