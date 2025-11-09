package model

import "github.com/google/uuid"

type Memberships struct {
	MembershipId uuid.UUID
	UserId       uuid.UUID
	ClubId       uuid.UUID
	Role         string
	Status       string
}
