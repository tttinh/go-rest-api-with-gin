package group

import (
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/infra/log"
	"github.com/tttinh/go-rest-api-with-gin/repository"
)

type Service interface {
	GetGroup(requesterID string, groupID string) (*GetGroupResponse, error)
	CreateGroup(requesterID string, req CreateGroupRequest) error
	UpdateGroup(requesterID string, groupID string, req UpdateGroupRequest) error
	DeleteGroup(requesterID string, groupID string) error
}

type Controller interface {
	GetGroup(c *gin.Context)
	CreateGroup(c *gin.Context)
	UpdateGroup(c *gin.Context)
	DeleteGroup(c *gin.Context)
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
