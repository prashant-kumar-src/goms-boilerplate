package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prashant-kumar-src/goms-boilerplate/services/users/config"
	"github.com/prashant-kumar-src/goms-boilerplate/utils/logger"
)

type routeRegisterer func(r *gin.Engine) *gin.Engine

func InitHttpServer(config *config.Config, logger logger.ILogger,
	routeRegisterer routeRegisterer,
	middlewares ...gin.HandlerFunc,
) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("pong %s", config.Name),
		})
	})
	//  configure middlewares
	r.Use(middlewares...)

	// configure routes
	routeRegisterer(r)

	return r.Run(":8080")
}
