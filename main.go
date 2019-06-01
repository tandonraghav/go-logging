package main

import (
	"github.com/go-logging/web/app"
	"github.com/go-logging/web/config"
	"github.com/go-logging/web/logging"
	logger "github.com/go-logging/web/logging"

	"context"
	_ "go.uber.org/zap"
)


var defaultContext = context.Background()

func init() {

	config.InitializeConfig()
	logging.InitializeLogger()
	logger.GetLogger(defaultContext).Info("This is actual line")
	logger.GetLogger(defaultContext).Info( "Initialized Logger config")
	go app.InitilizeMux()
	logger.GetLogger(defaultContext).Info( "Init Completed")
	//logging.InitLoggers()
	//appLogger.Debug(req.Context(),"Performing Actions",zap.Reflect("action",action))

}

func main() {
	select {

	}
}
