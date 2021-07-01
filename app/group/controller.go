package group

import (
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/infra/errcode"
	"github.com/tttinh/go-rest-api-with-gin/infra/transport/http"
)

type controllerImpl struct {
	service Service
}

func (ctrl *controllerImpl) GetGroup(c *gin.Context) {
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

func (ctrl *controllerImpl) CreateGroup(c *gin.Context) {
	var err error
	var req CreateGroupRequest

	err = http.BindAndValid(c, &req)
	if err != nil {
		http.Response400(c, errcode.InvalidInput, nil)
	}

	err = ctrl.service.CreateGroup(req)

	if err != nil {
		http.Response400(c, errcode.InvalidInput, nil)
	} else {
		http.Response200(c, nil)
	}
}
