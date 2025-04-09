package controller

import (
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

//func (uc *UserController) Register(c *gin.Context) {
//	result := uc.userService.Register("", "")
//	response.SuccessResponse(c, result, nil)
//}

func (uc *UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrResponse(c, 422, "Invalid user ID")
		return
	}

	user := uc.userService.GetUserByID(id)
	response.DataResponse(c, user)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		response.ErrResponse(c, 400, "Invalid request payload")
		return
	}

	userID, err := uc.userService.CreateUser(userRequest.Email, userRequest.Username, userRequest.Role)
	if err != nil {
		response.ErrResponse(c, 500, err.Error())
		return
	}

	response.DataResponse(c, gin.H{"user_id": userID})
}
