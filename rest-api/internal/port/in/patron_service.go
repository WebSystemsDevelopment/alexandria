package in

import "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"

type PatronService interface {
	CreatePatron(patron *domain.Patron) error
}
