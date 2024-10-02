package handler

import (
	"net/http"
	"strconv"

	model "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/models"
	service "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/services"

	"github.com/labstack/echo/v4"
)

// BookHandler handles book-related HTTP requests
type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the library
// @Tags books
// @Accept json
// @Produce json
// @Param book body model.Book true "Book details"
// @Success 201 {object} model.Book
// @Failure 400 {object} string "Bad Request"
// @Router /books [post]
func (h *BookHandler) CreateBook(c echo.Context) error {
	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		return err
	}
	err := h.service.CreateBook(book)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, book)
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Retrieve a list of all books in the library
// @Tags books
// @Produce json
// @Success 200 {array} model.Book
// @Failure 500 {object} string "Internal Server Error"
// @Router /books [get]
func (h *BookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.service.GetAllBooks()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, books)
}

// GetBookByID godoc
// @Summary Get a book by ID
// @Description Retrieve a specific book using its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} model.Book
// @Failure 404 {object} string "Book not found"
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.service.GetBookByID(id)
	if err != nil || book == nil {
		return c.JSON(http.StatusNotFound, "Book not found")
	}
	return c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update details of an existing book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body model.Book true "Updated book details"
// @Success 200 {object} model.Book
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Book not found"
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		return err
	}
	err := h.service.UpdateBook(id, book)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a specific book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} string "Book deleted"
// @Failure 404 {object} string "Book not found"
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteBook(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Book deleted")
}
