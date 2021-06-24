package group

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Endpoints struct {
	GetGroup gin.HandlerFunc
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		GetGroup: makeGetGroupEndpoint(s),
	}
}

func makeGetGroupEndpoint(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := GetGroupRequest{}
		res, _ := s.GetGroup(req)
		c.JSON(http.StatusOK, res)
	}
}
