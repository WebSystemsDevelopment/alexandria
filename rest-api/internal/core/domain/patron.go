package domain

import "github.com/google/uuid"

type Patron struct {
    Name string
    MembershipNumber uuid.UUID
    Email string
}

