package model

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	UserId       uuid.UUID
	FirstName    string
	LastName     string
	Username     string
	Password     string
	Email        string
	InterestId   uuid.UUID
	UserInterest Interest
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
