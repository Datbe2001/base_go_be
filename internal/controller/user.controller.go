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
// @Summary Run a task in the background
// @Description Starts a long-running task in the background and returns immediately
// @Tags user
// @Accept json
// @Produce json
// @Param data query string true "Data to process in the background"
// @Success 200 {object} response.Response "Background task started"
// @Failure 400 {object} response.Response "Missing data parameter"
// @Router /user/register [get]
func (uc *UserController) Register(c *gin.Context) {
	data := c.Query("data")
	if data == "" {
		c.JSON(400, gin.H{"error": "Missing data"})
		return
	}

	go longRunningTask(data)
	response.SuccessResponse(c, gin.H{"message": "🚀 Task đã được gửi xử lý background!"})
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
// @Success 200 {object} response.Response{data=[]dto.UserResponseDto} "List of users"
// @Router /user/list_user [get]
func (uc *UserController) GetListUser(c *gin.Context) {
	users := uc.userService.GetListUser()
	response.SuccessResponse(c, users)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	userRequest := dto.UserRequestDto{}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		response.ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	userID, err := uc.userService.CreateUser(userRequest.Email, userRequest.Username, userRequest.Role)
	if err != nil {
		response.ErrorResponse(c, 405, err.Error())
		return
	}

	response.SuccessResponse(c, gin.H{"user_id": userID})

}
