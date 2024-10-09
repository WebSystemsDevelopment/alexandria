package repository

import (
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/google/uuid"
)

type PatronRepository struct {
	patrons map[uuid.UUID]domain.Patron
}

func NewPatronRepository() ( PatronRepository) {
	return PatronRepository{
		patrons: make(map[uuid.UUID]domain.Patron),
	}
}

func (r PatronRepository) CreatePatron(patronRequest *domain.PatronRequest) error {
	id := uuid.New()
	lpatron := domain.Patron{
		Name:             patronRequest.Name,
		MembershipNumber: id,
		Email:            patronRequest.Email,
	}

	r.patrons[id] = lpatron
	return nil
}

func (r PatronRepository) GetAllPatrons() ([]domain.Patron, error) {
	var patrons []domain.Patron
	for _, patron := range r.patrons {
		patrons = append(patrons, patron)
	}
	return patrons, nil
}

