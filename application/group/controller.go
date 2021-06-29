package group

import (
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/pkg/errcode"
	"github.com/tttinh/go-rest-api-with-gin/pkg/httptransport"
	"net/http"
)

type Controller interface {
	SetRoutes(router *gin.RouterGroup)
	GetGroup(c *gin.Context)
}

type controllerImpl struct {
	service Service
}

func NewController(service Service) Controller {
	return &controllerImpl{
		service: service,
	}
}

func (ctrl *controllerImpl) SetRoutes(router *gin.RouterGroup) {
	router.GET("/:id", ctrl.GetGroup)
}

func (ctrl *controllerImpl) GetGroup(c *gin.Context) {
	req := GetGroupRequest{
		ID: c.Param("id"),
	}

	res, err := ctrl.service.GetGroup(req)

	if err != nil {
		httptransport.Response400(c, errcode.ErrInvalidInput, nil)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
