package services

import (
	"github.com/juanmabaracat/books-challenge/internal/domain/book"
	"log/slog"
)

type BookService interface {
	Update(book.Book) error
	Get(string) (*book.Book, error)
	GetAll() ([]book.Book, error)
}

func NewBookService(repository book.Repository) BookService {
	return &bookService{repository: repository}
}

type bookService struct {
	repository book.Repository
}

func (s *bookService) Update(book book.Book) error {
	err := s.repository.Update(book)
	if err != nil {
		slog.Error("error updating book", err)
		return err
	}

	return nil
}

func (s *bookService) Get(name string) (*book.Book, error) {
	b, err := s.repository.Get(name)
	if err != nil {
		slog.Error("error getting book", err)
		return nil, err
	}

	return b, nil
}

func (s *bookService) GetAll() ([]book.Book, error) {
	books, err := s.repository.GetAll()
	if err != nil {
		slog.Error("error getting book", err)
		return nil, err
	}

	return books, nil
}
