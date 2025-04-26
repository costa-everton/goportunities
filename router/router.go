package router

import "github.com/gin-gonic/gin"

func Initalize(){
	router := gin.Default()
	initializeRoutes(router)
	router.Run()
}
	
