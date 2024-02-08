package httphandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prashant-kumar-src/goms-boilerplate/services/users/services"
)

type IHttpUsersHandler interface {
	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	GetUserById(*gin.Context)
	GetAllUsers(*gin.Context)
}

type httpUsersHandler struct {
	UserSvc *services.IUsersService
}

// CreateUser implements IHttpUsersHandler.
func (httpUsersHandler) CreateUser(*gin.Context) {
	panic("unimplemented")
}

// DeleteUser implements IHttpUsersHandler.
func (httpUsersHandler) DeleteUser(*gin.Context) {
	panic("unimplemented")
}

// GetAllUsers implements IHttpUsersHandler.
func (httpUsersHandler) GetAllUsers(*gin.Context) {
	panic("unimplemented")
}

// GetUserById implements IHttpUsersHandler.
func (httpUsersHandler) GetUserById(*gin.Context) {
	panic("unimplemented")
}

// UpdateUser implements IHttpUsersHandler.
func (httpUsersHandler) UpdateUser(*gin.Context) {
	panic("unimplemented")
}

func NewHttpUsersHandler(userSvc *services.IUsersService) IHttpUsersHandler {
	return httpUsersHandler{
		UserSvc: userSvc,
	}
}
