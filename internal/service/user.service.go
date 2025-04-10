package service

import (
	"base_go_be/internal/dto"
	"base_go_be/internal/model"
	"base_go_be/internal/repo"
	"base_go_be/pkg/response"
	"fmt"
)

type IUserService interface {
	GetUserByID(id uint) *dto.UserResponseDto
	GetListUser() []*dto.UserResponseDto
	CreateUser(email string, username string, role string) (int, error)
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) GetUserByID(id uint) *dto.UserResponseDto {
	result := us.userRepo.GetUserByID(id)
	if result == nil {
		return nil
	}
	userResponse := dto.UserResponseDto{
		Id:       result.ID,
		Email:    result.Email,
		Username: result.Username,
		Role:     result.Role,
	}
	return &userResponse
}

func (us *userService) GetListUser() []*dto.UserResponseDto {
	results := us.userRepo.GetListUser()
	users := make([]*dto.UserResponseDto, 0, len(results))

	for _, u := range results {
		users = append(users, &dto.UserResponseDto{
			Id:       u.ID,
			Email:    u.Email,
			Username: u.Username,
			Role:     u.Role,
		})
	}
	return users
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
