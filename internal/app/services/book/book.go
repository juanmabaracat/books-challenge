package book

import (
	"github.com/juanmabaracat/books-challenge/internal/domain/book"
	"log/slog"
)

type Service struct {
	repository book.Repository
}

func (s *Service) Update(book book.Book) error {
	err := s.repository.Update(book)
	if err != nil {
		slog.Error("error updating book", err)
		return err
	}

	return nil
}
