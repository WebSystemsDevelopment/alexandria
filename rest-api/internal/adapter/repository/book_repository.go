package repository

import (
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/out"
)

type bookRepository struct {
	books []domain.Book
}

func NewBookRepository() out.BookRepository {
	return &bookRepository{books: []domain.Book{}}
}

func (r *bookRepository) CreateBook(book *domain.Book) error {
	book.ID = len(r.books) + 1
	r.books = append(r.books, *book)
	return nil
}

func (r *bookRepository) GetAllBooks() ([]domain.Book, error) {
	return r.books, nil
}

func (r *bookRepository) GetBookByID(id int) (*domain.Book, error) {
	for _, book := range r.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil
}

func (r *bookRepository) UpdateBook(id int, book *domain.Book) error {
	for i, b := range r.books {
		if b.ID == id {
			r.books[i] = *book
			r.books[i].ID = id
			return nil
		}
	}
	return nil
}

func (r *bookRepository) DeleteBook(id int) error {
	for i, b := range r.books {
		if b.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return nil
}
