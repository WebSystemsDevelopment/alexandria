package service

import (
	model "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/models"
	repository "github.com/WebSystemsDevelopment/alexandria/rest-api/internal/repositories"
)

type BookService interface {
	CreateBook(book *model.Book) error
	GetAllBooks() ([]model.Book, error)
	GetBookByID(id int) (*model.Book, error)
	UpdateBook(id int, book *model.Book) error
	DeleteBook(id int) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(book *model.Book) error {
	return s.repo.CreateBook(book)
}

func (s *bookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAllBooks()
}

func (s *bookService) GetBookByID(id int) (*model.Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *bookService) UpdateBook(id int, book *model.Book) error {
	return s.repo.UpdateBook(id, book)
}

func (s *bookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}
