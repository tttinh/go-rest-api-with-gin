package httperror

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ErrCode    Code        `json:"code"`
	ErrMessage string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Response400 returns bad request error.
func Response400(c *gin.Context, errCode Code, data interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		ErrCode:    errCode,
		ErrMessage: getErrorMessage(errCode),
		Data:       data,
	})
}
