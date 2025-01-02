package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", Pong)
	}
	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "dat")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message":   "pong...." + name,
		"uid":       uid,
		"list_user": []string{uid, name},
	})
}
