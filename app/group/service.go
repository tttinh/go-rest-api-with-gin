package group

import (
	"github.com/tttinh/go-rest-api-with-gin/repository"
)

type Service struct {
	groupRepository repository.GroupRepository
}

func NewService(groupRepository repository.GroupRepository) *Service {
	return &Service{
		groupRepository: groupRepository,
	}
}

func (s *Service) GetGroup(req GetGroupRequest) (*GetGroupResponse, error) {
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
