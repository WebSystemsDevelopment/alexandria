package routes

import (
	adapter "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/adapters"
	service "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/services"
	ports "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/ports"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	bookRepo := ports.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookHandler := adapter.NewBookHandler(bookService)

	e.POST("/books", bookHandler.CreateBook)
	e.GET("/books", bookHandler.GetAllBooks)
	e.GET("/books/:id", bookHandler.GetBookByID)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)
}
