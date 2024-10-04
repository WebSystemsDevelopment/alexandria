package repository

import (
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port"
	"github.com/google/uuid"
)

type PatronRepository struct {
    Database port.DatabasePort
    Collection map[uuid.UUID]domain.Patron
}

func NewPatronRepository() (repository PatronRepository) {
    repository := PatronRepository {

    }
    return
}

func (r *PatronRepository) create(name string, email string) {
    uuid := uuid.New()
    patron := domain.Patron {
        Name: name,
        MembershipNumber: uuid,
        Email: email,
    }

    r.Collection[uuid] = patron
    r.Database.createPatron(patron)
}

