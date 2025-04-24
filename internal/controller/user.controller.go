package controller

import (
	"base_go_be/internal/dto"
	"base_go_be/internal/service"
	"base_go_be/pkg/response"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func longRunningTask(data string) {
	// Giả lập tác vụ lâu (5s)
	fmt.Printf("⏳ Bắt đầu xử lý: %s\n", data)
	time.Sleep(5 * time.Second)
	fmt.Printf("✅ Xong xử lý: %s\n", data)
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequestDto true "User Registration Data"
// @Success 200 {object} response.Response{data=dto.AuthResponseDto} "Registration successful"
// @Failure 400 {object} response.Response "Invalid request data"
// @Failure 422 {object} response.Response "User already exists"
// @Router /user/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequestDto
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		response.ErrorResponse(c, 400, "Invalid request data")
		return
	}

	authResponse, err := uc.userService.Register(registerRequest)
	if err != nil {
		response.ErrorResponse(c, 422, err.Error())
		return
	}

	response.SuccessResponse(c, authResponse)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.LoginRequestDto true "User Login Data"
// @Success 200 {object} response.Response{data=dto.AuthResponseDto} "Login successful"
// @Failure 400 {object} response.Response "Invalid request data"
// @Failure 401 {object} response.Response "Invalid credentials"
// @Router /user/login [post]
func (uc *UserController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequestDto
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		response.ErrorResponse(c, 400, "Invalid request data")
		return
	}

	authResponse, err := uc.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response.ErrorResponse(c, 401, "Invalid email or password")
		return
	}

	response.SuccessResponse(c, authResponse)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieves a user by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.Response{data=dto.UserResponseDto} "User details"
// @Failure 422 {object} response.Response "Invalid user ID"
// @Router /user/get_user/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	idUint64, err := strconv.ParseUint(idParam, 10, 0)
	id := uint(idUint64)
	if err != nil {
		response.DataDetailResponse(c, 422, response.ErrCodeInvalidParams, nil)
		return
	}

	user := uc.userService.GetUserByID(id)
	response.SuccessResponse(c, user)
}

// GetListUser godoc
// @Summary Get all users
// @Description Returns a list of all users
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=[]dto.UserResponseDto} "List of users"
// @Failure 401 {object} response.Response "Unauthorized"
// @Router /user/list_user [get]
func (uc *UserController) GetListUser(c *gin.Context) {
	users := uc.userService.GetListUser()
	response.SuccessResponse(c, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Creates a new user with the provided information
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body dto.UserRequestDto true "User Information"
// @Success 200 {object} response.Response{data=map[string]interface{}} "User created successfully"
// @Failure 400 {object} response.Response "Invalid request payload"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 405 {object} response.Response "User already exists"
// @Router /user/create_user [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	userRequest := dto.UserRequestDto{}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		response.ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	userID, err := uc.userService.CreateUser(userRequest.Email, userRequest.Username, userRequest.Password, userRequest.Role)
	if err != nil {
		response.ErrorResponse(c, 405, err.Error())
		return
	}

	response.SuccessResponse(c, gin.H{"user_id": userID})
}

// GetCurrentUser godoc
// @Summary Get current user
// @Description Get the currently logged in user's information
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=dto.UserResponseDto} "Current user"
// @Failure 401 {object} response.Response "Unauthorized"
// @Router /user/me [get]
func (uc *UserController) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.ErrorResponse(c, 401, "Unauthorized")
		return
	}

	user := uc.userService.GetUserByID(userID.(uint))
	response.SuccessResponse(c, user)
}
