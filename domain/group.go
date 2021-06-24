package domain

import "time"

type Group struct {
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
}
