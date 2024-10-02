package repository

import (
	model "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/models"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetAllBooks() ([]model.Book, error)
	GetBookByID(id int) (*model.Book, error)
	UpdateBook(id int, book *model.Book) error
	DeleteBook(id int) error
}

type bookRepository struct {
	books []model.Book
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		books: []model.Book{},
	}
}

func (r *bookRepository) CreateBook(book *model.Book) error {
	book.ID = len(r.books) + 1 // Auto-increment ID for now
	r.books = append(r.books, *book)
	return nil
}

func (r *bookRepository) GetAllBooks() ([]model.Book, error) {
	return r.books, nil
}

func (r *bookRepository) GetBookByID(id int) (*model.Book, error) {
	for _, book := range r.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil
}

func (r *bookRepository) UpdateBook(id int, book *model.Book) error {
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
