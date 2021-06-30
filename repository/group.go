package repository

import (
	"github.com/tttinh/go-rest-api-with-gin/entity"
	"gorm.io/gorm"
)

type GroupRepository interface {
	GetGroup(id string) (*entity.Group, error)
}

//var RepoErr = errors.New("unable to handle repository request")

type groupRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *groupRepositoryImpl {
	return &groupRepositoryImpl{
		db: db,
	}
}

// GetGroup Get a single group based on id
func (r *groupRepositoryImpl) GetGroup(id string) (*entity.Group, error) {
	var group entity.Group
	err := r.db.Where("id = ? AND deleted = 0", id).First(&group).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &group, nil
}
