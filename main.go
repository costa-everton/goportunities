package main

import (
	"github.com/costa-everton/goportunities/config"
	"github.com/costa-everton/goportunities/router"
)

var (
	logger *config.Logger
)

func main(){
	logger = config.GetLogger("main")

	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// initialize router
	router.Initalize()
} 