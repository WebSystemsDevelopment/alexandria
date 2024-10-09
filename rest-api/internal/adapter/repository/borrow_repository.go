package repository

import (
	"time"

	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/out"
	"github.com/google/uuid"
)

type BorrowRepository struct {
	borrows map[uuid.UUID]domain.Borrow
}

func NewBorrowRepository() (out.BorrowRepository){
	return &BorrowRepository{
		borrows: make(map[uuid.UUID]domain.Borrow),
	}
}

func (b *BorrowRepository) CreateBorrow(borrow domain.BorrowRequest) error {
	id := uuid.New()
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, borrow.DueDate)
	if err != nil {
		return err
	}

	lborrow := domain.Borrow{
		Id: id,
		PatronId: borrow.PatronId,
		BookId:   borrow.BookId,
		DueDate:  parsedDate,
		IsActive: true,
	}

	b.borrows[id] = lborrow
	return nil
}

func (b *BorrowRepository) GetAllBorrows() ([]domain.Borrow, error) {
	var items []domain.Borrow
	for _, item := range b.borrows {
		items = append(items, item)
	}
	return items, nil
}

