package domain

import (
	"time"

	"github.com/google/uuid"
)

type Borrow struct {
	Id        uuid.UUID
	PatronId  uuid.UUID
	BookId    int
	DueDate   time.Time
	IsActive  bool
}

type BorrowRequest struct {
	PatronId  uuid.UUID  `json:"patronId"`
	BookId    int        `json:"bookId"`
	DueDate   string  `json:"dueDate"`
}
