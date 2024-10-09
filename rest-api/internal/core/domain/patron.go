package domain

import "github.com/google/uuid"

type Patron struct {
	Name              string
	MembershipNumber  uuid.UUID
	Email             string
}

type PatronRequest struct {
	Name              string  `json:"name"`
	Email             string  `json:"email"`
}
