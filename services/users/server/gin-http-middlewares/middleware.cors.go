package ginmiddlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prashant-kumar-src/goms-boilerplate/services/users/config"
)

func GinCorsMiddleware(config *config.Config) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowOrigins:     config.AllowedOrigins,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	})
}
