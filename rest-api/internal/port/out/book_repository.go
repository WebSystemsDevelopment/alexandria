package out

import "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"

type BookRepository interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id int) (*domain.Book, error)
	CreateBook(book *domain.Book) error
	UpdateBook(id int, book *domain.Book) error
	DeleteBook(id int) error
}
