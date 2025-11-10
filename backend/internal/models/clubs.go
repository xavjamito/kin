package model

import "github.com/google/uuid"

type Clubs struct {
	ClubsId     uuid.UUID
	ClubName    string
	Description string
}
