package controller

import (
	"base_go_be/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{
		"name": result,
	})

}
