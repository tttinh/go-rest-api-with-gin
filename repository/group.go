package repository

import (
	"context"
	"database/sql"
	"errors"
)

type GroupRepository interface {
	GetGroup(ctx context.Context, id string) (string, error)
}

var RepoErr = errors.New("unable to handle repository request")

type groupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *groupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) GetGroup(ctx context.Context, id string) (string, error) {
	var description string
	err := r.db.QueryRow("SELECT description FROM 'group' WHERE id=$1", id).Scan(&description)
	if err != nil {
		return "", RepoErr
	}

	return description, nil
}
