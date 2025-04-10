package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func DataDetailResponse(c *gin.Context, statusCode int, code int, data interface{}) {
	c.JSON(statusCode, Response{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func ErrorResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}
