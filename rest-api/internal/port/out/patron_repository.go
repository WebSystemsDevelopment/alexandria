package out

import "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"

type PatronRepository interface {
	GetAllPatrons() ([]domain.Patron, error)
	CreatePatron(book *domain.Patron) error
}
