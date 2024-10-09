package in

import (
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
)

type BorrowService interface {
    CreateBorrow(domain.BorrowRequest) error
    CheckBorrow(id int) (int, error)
    GetAllBorrows() ([]domain.Borrow, error)
    HandleReturn(id int) error
}

