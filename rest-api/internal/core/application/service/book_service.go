package service

import (
	domain "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/core/domain"
	repository "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/ports"
)

type BookService interface {
	CreateBook(book *domain.Book) error
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id int) (*domain.Book, error)
	UpdateBook(id int, book *domain.Book) error
	DeleteBook(id int) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
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
