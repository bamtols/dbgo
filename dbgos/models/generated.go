package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserSession struct {
	SessionId *string
	UserId    uint
	CreatedAt time.Time
	DeletedAt *time.Time
	User      *User
}
type User struct {
	Password  string
	Point     decimal.Decimal
	Username  string
	JoinAt    time.Time
	Id        *uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//67065e84-08b5-4a2a-a66d-c840fe9a2096
