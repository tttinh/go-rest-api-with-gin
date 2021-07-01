package group

import (
	"github.com/tttinh/go-rest-api-with-gin/repository"
)

type serviceImpl struct {
	groupRepository repository.GroupRepository
}

func (s *serviceImpl) GetGroup(req GetGroupRequest) (*GetGroupResponse, error) {
	group, _ := s.groupRepository.GetGroup(req.ID)

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

func (s *serviceImpl) CreateGroup(req CreateGroupRequest) error {
	return nil
}
