package main

import (
	"context"
	"github.com/tandonraghav/go-logging/config"
	logger "github.com/tandonraghav/go-logging/logging"
	"github.com/tandonraghav/go-logging/web/mux"
	_ "go.uber.org/zap"
)


var defaultContext = context.Background()

func init() {
	config.InitializeConfig()
	logger.InitializeLogger()
	logger.GetLogger(defaultContext).Info( "Initialized Logger config")

	logger.GetLogger(defaultContext).Info( "Init Completed")

}

func main() {
	go mux.InitilizeMux()
	select {

	}
}
