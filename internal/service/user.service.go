package service

import (
	"base_go_be/internal/model"
	"base_go_be/internal/repo"
	"base_go_be/pkg/response"
	"fmt"
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
	// GetUserByID  Register(email string, password string) int
	GetUserByID(id int) *model.User
	CreateUser(email string, username string, role string) (int, error)
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

func (us *userService) GetUserByID(id int) *model.User {
	result := us.userRepo.GetUserByID(id)
	if result == nil {
		return nil
	}
	return result
}

func (us *userService) CreateUser(email string, username string, role string) (int, error) {
	if us.userRepo.GetUserByEmail(email) {
		return 0, fmt.Errorf("user already exists: %d", response.ErrCodeUserHasExists)
	}

	user := &model.User{
		Email:    email,
		Username: username,
		Role:     role,
	}

	return us.userRepo.CreateUser(user)
}
