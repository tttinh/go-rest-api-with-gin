package group

import (
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/infra/errcode"
	"github.com/tttinh/go-rest-api-with-gin/infra/transport/http"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetGroup(c *gin.Context) {
	req := GetGroupRequest{
		ID: c.Param("id"),
	}

	res, err := ctrl.service.GetGroup(req)

	if err != nil {
		http.Response400(c, errcode.InvalidInput, nil)
	} else {
		http.Response200(c, res)
	}
}
