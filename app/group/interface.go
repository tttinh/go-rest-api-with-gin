package group

import (
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/infra/log"
	"github.com/tttinh/go-rest-api-with-gin/repository"
)

type Service interface {
	GetGroup(req GetGroupRequest) (*GetGroupResponse, error)
	CreateGroup(req CreateGroupRequest) error
}

type Controller interface {
	GetGroup(c *gin.Context)
	CreateGroup(c *gin.Context)
}

func NewService(repo repository.GroupRepository) Service {
	return &serviceImpl{repo}
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func NewController(service Service) *controllerImpl {
	return &controllerImpl{
		service: service,
	}
}
