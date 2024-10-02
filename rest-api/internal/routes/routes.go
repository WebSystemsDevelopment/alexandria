package routes

import (
	handler "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/handlers"
	repository "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/repositories"
	service "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/services"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	e.POST("/books", bookHandler.CreateBook)
	e.GET("/books", bookHandler.GetAllBooks)
	e.GET("/books/:id", bookHandler.GetBookByID)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)
}
