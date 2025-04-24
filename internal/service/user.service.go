package service

import (
	"base_go_be/internal/dto"
	"base_go_be/internal/model"
	"base_go_be/internal/repo"
	"base_go_be/pkg/config"
	"base_go_be/pkg/jwt"
	"base_go_be/pkg/response"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetUserByID(id uint) *dto.UserResponseDto
	GetListUser() []*dto.UserResponseDto
	CreateUser(email string, username string, password string, role string) (uint, error)
	Login(email string, password string) (*dto.AuthResponseDto, error)
	Register(registerDto dto.RegisterRequestDto) (*dto.AuthResponseDto, error)
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

func (us *userService) CreateUser(email string, username string, password string, role string) (uint, error) {
	if us.userRepo.GetUserByEmail(email) {
		return 0, fmt.Errorf("user already exists: %d", response.ErrCodeUserHasExists)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := &model.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	return us.userRepo.CreateUser(user)
}

func (us *userService) Login(email string, password string) (*dto.AuthResponseDto, error) {
	user := us.userRepo.GetUserByEmailFull(email)
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(user.ID, user.Email, user.Role, config.JWT.SecretKey, config.JWT.TokenExpiry)
	if err != nil {
		return nil, err
	}

	// Generate refresh token with longer expiry
	refreshToken, err := jwt.GenerateToken(user.ID, user.Email, user.Role, config.JWT.SecretKey, config.JWT.RefreshExpiry)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponseDto{
		Token:        token,
		RefreshToken: refreshToken,
		User: dto.UserResponseDto{
			Id:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Role:     user.Role,
		},
	}, nil
}

func (us *userService) Register(registerDto dto.RegisterRequestDto) (*dto.AuthResponseDto, error) {
	userID, err := us.CreateUser(registerDto.Email, registerDto.Username, registerDto.Password, registerDto.Role)
	if err != nil {
		return nil, err
	}

	user := us.userRepo.GetUserByID(userID)
	if user == nil {
		return nil, errors.New("failed to retrieve created user")
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(user.ID, user.Email, user.Role, config.JWT.SecretKey, config.JWT.TokenExpiry)
	if err != nil {
		return nil, err
	}

	// Generate refresh token with longer expiry
	refreshToken, err := jwt.GenerateToken(user.ID, user.Email, user.Role, config.JWT.SecretKey, config.JWT.RefreshExpiry)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponseDto{
		Token:        token,
		RefreshToken: refreshToken,
		User: dto.UserResponseDto{
			Id:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Role:     user.Role,
		},
	}, nil
}
