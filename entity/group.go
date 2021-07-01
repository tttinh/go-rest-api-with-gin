package entity

import (
	"github.com/tttinh/go-rest-api-with-gin/infra/common"
	"time"
)

type Group struct {
	ID            string
	Privacy       string
	OwnerID       string
	Name          string
	Category      string
	Location      string
	Avatar        string
	Cover         string
	Description   string
	Terms         string
	MemberCount   int
	Deleted       common.BitBool
	JoinByDefault common.BitBool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TableName sets the table name for the entity.
func (Group) TableName() string {
	return "group"
}
