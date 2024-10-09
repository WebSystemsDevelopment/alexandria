package adapter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/in"
	"github.com/labstack/echo/v4"
)

type BorrowHandler struct {
	service in.BorrowService
}

func NewBorrowHandler(service in.BorrowService) *BorrowHandler {
	return &BorrowHandler{
		service: service,
	}
}

func (b *BorrowHandler) CreateBorrow(context echo.Context) error {
	borrow := new(domain.BorrowRequest)
	if err := context.Bind(borrow); err != nil {
		return err
	}

	isBorrowed, err  := b.service.CheckBorrow(borrow.BookId)
	if err != nil {
		return err
	}

	if isBorrowed == 1 {
		return context.JSON(http.StatusBadRequest, fmt.Sprintf("The book %d is not available for lending", borrow.BookId))
	}

	if err := b.service.CreateBorrow(*borrow); err != nil {
		return err
	}
	
	return context.JSON(http.StatusCreated, borrow)
}

func (b *BorrowHandler) CheckBorrow(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	result, err := b.service.CheckBorrow(id)
	if err != nil {
		return err
	}
	
	return	c.JSON(http.StatusOK, result)
}

func (b *BorrowHandler) GetAllBorrows(c echo.Context) error {
	borrows, err := b.service.GetAllBorrows()
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, borrows)
}

func (b *BorrowHandler) HandleReturn(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if b.service.HandleReturn(id) != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Could not return book %d: %v", id, err))
	}
	
	return	c.JSON(http.StatusOK, "Success")
}

