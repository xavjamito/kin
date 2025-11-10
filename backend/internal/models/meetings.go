package model

import (
	"time"

	"github.com/google/uuid"
)

type Meetings struct {
	MeetingId   uuid.UUID
	ClubID      uuid.UUID
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Location    string
	IsRecurring bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
