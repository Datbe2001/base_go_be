package service

import (
	"base_go_be/internal/repo"
	"base_go_be/pkg/response"
)

//type UserService struct {
//	userRepo *repo.UserRepo
//}
//
//func NewUserService() *UserService {
//	return &UserService{
//		userRepo: repo.NewUserRepo(),
//	}
//}
//
//func (us *UserService) GetMe(name string) string {
//	return us.userRepo.GetMe(name)
//}

type IUserService interface {
	Register(email string, password string) int
	//	...
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) Register(email string, password string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}
