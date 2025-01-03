package controller

import (
	"base_go_be/internal/service"
	"base_go_be/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetMe(c *gin.Context) {
	name := c.Param("name")
	//uid := c.Query("uid")
	result := uc.userService.GetMe(name)
	response.SuccessResponse(c, 2001, result)
	response.FailResponse(c, 2002)

}
