package port

import "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"

type DatabasePort interface {
    createPatron(patron domain.Patron)
}

