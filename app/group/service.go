package group

import (
	"github.com/tttinh/go-rest-api-with-gin/entity"
	"github.com/tttinh/go-rest-api-with-gin/repository"
)

type serviceImpl struct {
	groupRepository repository.GroupRepository
}

func (s *serviceImpl) GetGroup(requesterID string, groupID string) (*GetGroupResponse, error) {
	group, err := s.groupRepository.FindGroup(groupID)

	if err != nil {
		return nil, err
	}

	res := &GetGroupResponse{
		ID:            group.ID,
		Privacy:       group.Privacy,
		OwnerID:       group.OwnerID,
		Name:          group.Name,
		Category:      group.Category,
		Location:      group.Location,
		Avatar:        group.Avatar,
		Cover:         group.Cover,
		Description:   group.Description,
		Terms:         group.Terms,
		MemberCount:   group.MemberCount,
		Deleted:       bool(group.Deleted),
		JoinByDefault: bool(group.JoinByDefault),
		CreatedAt:     group.CreatedAt.Unix() * 1000,
		UpdatedAt:     group.UpdatedAt.Unix() * 1000,
	}

	return res, nil
}

func (s *serviceImpl) CreateGroup(requesterID string, req CreateGroupRequest) error {
	group := &entity.Group{
		Privacy:       req.Privacy,
		OwnerID:       requesterID,
		Name:          req.Name,
		Category:      req.Category,
		Location:      req.Location,
		Avatar:        req.Avatar,
		Cover:         req.Cover,
		Description:   req.Description,
		Terms:         req.Terms,
		MemberCount:   0,
		Deleted:       false,
		JoinByDefault: false,
	}

	return s.groupRepository.AddGroup(group)
}

func (s *serviceImpl) UpdateGroup(requesterID string, groupID string, req UpdateGroupRequest) error {
	group, err := s.groupRepository.FindGroup(groupID)

	if err != nil {
		return err
	}

	group.Privacy = req.Privacy
	group.Name = req.Name
	group.Category = req.Category
	group.Location = req.Location
	group.Avatar = req.Avatar
	group.Cover = req.Cover
	group.Description = req.Description
	group.Terms = req.Terms

	return s.groupRepository.UpdateGroup(group)
}

func (s *serviceImpl) DeleteGroup(requesterID string, groupID string) error {
	group, err := s.groupRepository.FindGroup(groupID)

	if err != nil {
		return err
	}

	return s.groupRepository.DeleteGroup(group)
}
