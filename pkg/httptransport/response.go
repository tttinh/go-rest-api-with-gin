package httptransport

import (
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/pkg/errcode"
	"net/http"
)

type Response struct {
	Status     string       `json:"status"`
	ErrCode    errcode.Code `json:"code"`
	ErrMessage string       `json:"message"`
	Data       interface{}  `json:"data"`
}

// Response400 returns bad request error.
func Response400(c *gin.Context, errCode errcode.Code, data interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Status:     "ERROR",
		ErrCode:    errCode,
		ErrMessage: errcode.GetMessage(errCode),
		Data:       data,
	})
}

// Response200 returns success.
func Response200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:     "SUCCESS",
		ErrCode:    errcode.OK,
		ErrMessage: errcode.GetMessage(errcode.OK),
		Data:       data,
	})
}
