package model

import "time"

type Users struct {
	UserId    int32
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
