package httproutes

import (
	"github.com/gin-gonic/gin"
	httphandlers "github.com/prashant-kumar-src/goms-boilerplate/services/users/server/http-handlers"
)

func v1(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1")
	v1UserRoutes(v1)
	return r
}

func v1UserRoutes(r *gin.RouterGroup) {
	userRoutes := r.Group("/users")
	userRoutes.GET("", httphandlers.HttpUsersHandler.GetAllUsers)
	userRoutes.GET("/:id", httphandlers.HttpUsersHandler.GetUserById)
	userRoutes.POST("", httphandlers.HttpUsersHandler.CreateUser)
	userRoutes.PUT("/:id", httphandlers.HttpUsersHandler.UpdateUser)
	userRoutes.DELETE("/:id", httphandlers.HttpUsersHandler.DeleteUser)
}
