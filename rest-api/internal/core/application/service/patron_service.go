package service

import (
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/in"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/out"
)

type PatronService struct {
	repo out.PatronRepository
}

func NewPatronService(repo out.PatronRepository) in.PatronService {
	return &PatronService{repo: repo}
}

func (p *PatronService) GetAllPatrons() ([]domain.Patron, error) {
	return p.repo.GetAllPatrons()
}

func (p *PatronService) CreatePatron(patron *domain.PatronRequest) error {
	return p.repo.CreatePatron(patron)
}

