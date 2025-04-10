package controller

import (
	"base_go_be/internal/dto"
	"base_go_be/internal/service"
	"base_go_be/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.DataDetailResponse(c, 422, response.ErrCodeInvalidParams, nil)
		return
	}

	user := uc.userService.GetUserByID(id)
	response.SuccessResponse(c, user)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	userRequest := dto.UserRequestDto{}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		response.ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	userID, err := uc.userService.CreateUser(userRequest.Email, userRequest.Username, userRequest.Role)
	if err != nil {
		response.ErrorResponse(c, 500, err.Error())
		return
	}

	response.SuccessResponse(c, gin.H{"user_id": userID})

}
