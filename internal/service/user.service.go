package service

import (
	"base_go_be/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetMe(name string) string {
	return us.userRepo.GetMe(name)
}
