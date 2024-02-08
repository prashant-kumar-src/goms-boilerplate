package main

import (
	"os"

	"github.com/prashant-kumar-src/goms-boilerplate/services/users/config"
	"github.com/prashant-kumar-src/goms-boilerplate/services/users/server"
	ginmiddlewares "github.com/prashant-kumar-src/goms-boilerplate/services/users/server/gin-http-middlewares"
	httproutes "github.com/prashant-kumar-src/goms-boilerplate/services/users/server/http-routes"
	"github.com/prashant-kumar-src/goms-boilerplate/utils/logger"
)

func main() {
	config := config.GetConfig()
	logger := logger.GetLogger(logger.LoggerConfig{
		LoggerClient: config.LoggerClient,
	})
	logger.Infof("logger configured for %v", config.LoggerClient)
	logger.Infof("Starting %v", config.Name)

	if err := server.InitHttpServer(config, logger,
		httproutes.ConfigureRoutes,
		ginmiddlewares.GinCorsMiddleware(config)); err != nil {
		logger.Fatalf("failed to start http server: %v", err)
		os.Exit(1)
	}

	logger.Infof("Started %v", config.Name)
}
