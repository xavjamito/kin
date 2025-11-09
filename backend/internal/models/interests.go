package model

import (
	"time"

	"github.com/google/uuid"
)

type Interest struct {
	InterestID  uuid.UUID
	Name        string
	Description string
	IconURL     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
}
