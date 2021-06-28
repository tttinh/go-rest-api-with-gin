package repository

import (
	"github.com/jinzhu/gorm"
)

type GroupRepository interface {
	GetGroup(id string) (string, error)
}

//var RepoErr = errors.New("unable to handle repository request")

type groupRepositoryImpl struct {
	db *gorm.DB
}

func makeGroupRepository(db *gorm.DB) *groupRepositoryImpl {
	return &groupRepositoryImpl{db: db}
}

func (r *groupRepositoryImpl) GetGroup(id string) (string, error) {
	var description string
	//err := r.db.QueryRow("SELECT description FROM 'group' WHERE id=$1", id).Scan(&description)
	//if err != nil {
	//	return "", RepoErr
	//}

	return description, nil
}
