package service

import (
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/in"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/out"
)

type BorrowService struct {
	repo out.BorrowRepository
}

func NewBorrowService(r out.BorrowRepository) in.BorrowService {
	return &BorrowService{
		repo: r,
	}
}

func (b *BorrowService) CreateBorrow(borrow domain.BorrowRequest) error {
	return b.repo.CreateBorrow(borrow)
}

func (b *BorrowService) CheckBorrow(bookId int) (int, error) {
	borrows, err := b.repo.GetAllBorrows()
	if err != nil {
		return 0, err
	}

	for index := range borrows {
		book := borrows[index]
		if bookId == book.BookId && book.IsActive {
			return 1, nil
		}
	}

	return 0, nil
}

func (b *BorrowService) GetAllBorrows() ([]domain.Borrow, error) {
	return b.repo.GetAllBorrows()
}

func (b *BorrowService) HandleReturn(id int) error {
	borrows, err := b.repo.GetAllBorrows()
	if err != nil {
		return err
	}

	for index := range borrows {
		book := borrows[index]
		if id == book.BookId && book.IsActive {
			book.IsActive = false
			break
		}
	}

	return nil
}

