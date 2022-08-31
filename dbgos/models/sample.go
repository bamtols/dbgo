package models

import "time"

type Sample struct {
	Id        uint
	Password  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
