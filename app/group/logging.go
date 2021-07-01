package group

import (
	"github.com/tttinh/go-rest-api-with-gin/infra/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func (s *loggingService) GetGroup(req GetGroupRequest) (res *GetGroupResponse, err error) {
	defer func(begin time.Time) {
		s.logger.Infow("GetGroup",
			"req", req,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetGroup(req)
}

func (s *loggingService) CreateGroup(req CreateGroupRequest) (err error) {
	defer func(begin time.Time) {
		s.logger.Infow("CreateGroup",
			"req", req,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CreateGroup(req)
}
