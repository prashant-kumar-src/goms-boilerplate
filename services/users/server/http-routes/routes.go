package httproutes

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine) *gin.Engine {
	// configure v1 routes
	v1(r)
	return r
}
