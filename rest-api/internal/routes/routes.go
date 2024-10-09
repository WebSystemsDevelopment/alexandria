package routes

import (
	adapter "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/adapter/http"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/adapter/repository"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/application/service"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookHandler := adapter.NewBookHandler(bookService)

	patronRepo := repository.NewPatronRepository()
	patronService := service.NewPatronService(patronRepo)
	patronHandler := adapter.NewPatronHandler(patronService)

	borrowRepo := repository.NewBorrowRepository()
	borrowService := service.NewBorrowService(borrowRepo)
	borrowHandler := adapter.NewBorrowHandler(borrowService)

	e.POST("/books", bookHandler.CreateBook)
	e.GET("/books", bookHandler.GetAllBooks)
	e.GET("/books/:id", bookHandler.GetBookByID)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)

	e.POST("/patron", patronHandler.CreatePatron)
	e.GET("/patron", patronHandler.GetAllPatrons)

	e.POST("/borrow", borrowHandler.CreateBorrow)
	e.GET("/borrow", borrowHandler.GetAllBorrows)
	e.GET("/borrow/availability/:id", borrowHandler.CheckBorrow)
	e.GET("/borrow/return/:id", borrowHandler.HandleReturn)
}

