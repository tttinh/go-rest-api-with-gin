package group

import (
	"example.com/demo/repository"
)

type Service interface {
	GetGroup(req GetGroupRequest) (GetGroupResponse, error)
}

type service struct {
	groupRepository repository.GroupRepository
}

func NewService(groupRepository repository.GroupRepository) Service {
	return &service{
		groupRepository: groupRepository,
	}
}

func (s *service) GetGroup(req GetGroupRequest) (GetGroupResponse, error) {
	return GetGroupResponse{Description: "Haha, nice to see you."}, nil
}
