package repository

import (
	"github.com/google/uuid"
	"github.com/tttinh/go-rest-api-with-gin/entity"
	"gorm.io/gorm"
)

type GroupRepository interface {
	FindGroup(id string) (*entity.Group, error)
	AddGroup(group *entity.Group) error
	UpdateGroup(group *entity.Group) error
	DeleteGroup(group *entity.Group) error
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

// FindGroup finds a single group based on id
func (r *groupRepositoryImpl) FindGroup(id string) (*entity.Group, error) {
	var group entity.Group
	err := r.db.Where("id = ? AND deleted = 0", id).First(&group).Error
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// AddGroup adds a new group
func (r *groupRepositoryImpl) AddGroup(group *entity.Group) error {
	group.ID = "group-" + uuid.New().String()
	if err := r.db.Create(group).Error; err != nil {
		return err
	}

	return nil
}

// UpdateGroup updates a group
func (r *groupRepositoryImpl) UpdateGroup(group *entity.Group) error {
	if err := r.db.Save(group).Error; err != nil {
		return err
	}

	return nil
}

// DeleteGroup deletes a group
func (r *groupRepositoryImpl) DeleteGroup(group *entity.Group) error {
	if err := r.db.Delete(group).Error; err != nil {
		return err
	}

	return nil
}
