package main

import (
	"github.com/BrunoNeves1995/api-portunities/config"
	"github.com/BrunoNeves1995/api-portunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Config
	err := config.InitDB()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	//Initialize router
	router.Initialize()

}
