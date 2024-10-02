package service

import (
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/in"
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/port/out"
)

type bookService struct {
	repo out.BookRepository
}

func NewBookService(repo out.BookRepository) in.BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(book *domain.Book) error {
	return s.repo.CreateBook(book)
}

func (s *bookService) GetAllBooks() ([]domain.Book, error) {
	return s.repo.GetAllBooks()
}

func (s *bookService) GetBookByID(id int) (*domain.Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *bookService) UpdateBook(id int, book *domain.Book) error {
	return s.repo.UpdateBook(id, book)
}

func (s *bookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}
