package main

import (
	"github.com/HTM1000/go-oportunities/config"
	"github.com/HTM1000/go-oportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize routes
	router.Initialize()
}
