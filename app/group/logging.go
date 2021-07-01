package group

import (
	"github.com/tttinh/go-rest-api-with-gin/infra/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func (s *loggingService) GetGroup(requesterID string, groupID string) (res *GetGroupResponse, err error) {
	defer func(begin time.Time) {
		s.logger.Infow("get_group",
			"requester", requesterID,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetGroup(requesterID, groupID)
}

func (s *loggingService) CreateGroup(requesterID string, req CreateGroupRequest) (err error) {
	defer func(begin time.Time) {
		s.logger.Infow("create_group",
			"requester", requesterID,
			"req", req,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CreateGroup(requesterID, req)
}

func (s *loggingService) UpdateGroup(requesterID string, groupID string, req UpdateGroupRequest) (err error) {
	defer func(begin time.Time) {
		s.logger.Infow("update_group",
			"requester", requesterID,
			"group", groupID,
			"req", req,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.UpdateGroup(requesterID, groupID, req)
}

func (s *loggingService) DeleteGroup(requesterID string, groupID string) (err error) {
	defer func(begin time.Time) {
		s.logger.Infow("delete_group",
			"requester", requesterID,
			"group", groupID,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteGroup(requesterID, groupID)
}
