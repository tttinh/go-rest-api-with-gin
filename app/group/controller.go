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

	groupID := c.Param("id")
	res, err := ctrl.service.GetGroup("", groupID)

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

	err = ctrl.service.CreateGroup("", req)

	if err != nil {
		http.Response400(c, errcode.InvalidInput, nil)
	} else {
		http.Response200(c, nil)
	}
}

func (ctrl *controllerImpl) UpdateGroup(c *gin.Context) {
	var err error
	var req UpdateGroupRequest

	err = http.BindAndValid(c, &req)
	if err != nil {
		http.Response400(c, errcode.InvalidInput, nil)
	}

	groupID := c.Param("id")
	err = ctrl.service.UpdateGroup("", groupID, req)

	if err != nil {
		http.Response400(c, errcode.InvalidInput, nil)
	} else {
		http.Response200(c, nil)
	}
}

func (ctrl *controllerImpl) DeleteGroup(c *gin.Context) {
	var err error
	groupID := c.Param("id")
	err = ctrl.service.DeleteGroup("", groupID)

	if err != nil {
		http.Response400(c, errcode.InvalidInput, nil)
	} else {
		http.Response200(c, nil)
	}
}
