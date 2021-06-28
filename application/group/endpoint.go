package group

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func makeGetGroupEndpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := GetGroupRequest{}
		res, _ := s.GetGroup(req)
		c.JSON(http.StatusOK, res)
	}
}
