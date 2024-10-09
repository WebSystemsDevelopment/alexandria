package out

import "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"

type BorrowRepository interface {
	GetAllBorrows() ([]domain.Borrow, error)
	CreateBorrow(domain.BorrowRequest) error
}
