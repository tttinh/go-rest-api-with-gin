package repository

import (
	"github.com/tttinh/go-rest-api-with-gin/domain"
	"gorm.io/gorm"
)

type GroupRepository interface {
	GetGroup(id string) (*domain.Group, error)
}

//var RepoErr = errors.New("unable to handle repository request")

type groupRepositoryImpl struct {
}

func makeGroupRepository() *groupRepositoryImpl {
	return &groupRepositoryImpl{}
}

// GetGroup Get a single group based on id
func (r *groupRepositoryImpl) GetGroup(id string) (*domain.Group, error) {
	var group domain.Group
	err := db.Where("id = ? AND deleted = 0", id).First(&group).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &group, nil
}
