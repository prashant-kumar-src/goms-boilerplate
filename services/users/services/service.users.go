package services

import userRepo "github.com/prashant-kumar-src/goms-boilerplate/services/users/repositories/user"

type IUsersService interface {
}

type usersService struct {
	userRepo *userRepo.IUserRepository
}
