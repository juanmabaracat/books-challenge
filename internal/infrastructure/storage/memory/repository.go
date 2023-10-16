package memory

import (
	"github.com/juanmabaracat/books-challenge/internal/domain/book"
)

func NewRepository() book.Repository {
	return &Repository{make(map[string]book.Book)}
}

type Repository struct {
	books map[string]book.Book
}

func (r *Repository) Update(book book.Book) error {
	r.books[book.Name] = book
	return nil
}

func (r *Repository) GetAll() ([]book.Book, error) {
	books := make([]book.Book, 0)
	for _, v := range r.books {
		books = append(books, v)
	}

	return books, nil
}

func (r *Repository) Get(name string) (*book.Book, error) {
	b, found := r.books[name]
	if !found {
		return nil, nil
	}

	return &b, nil
}
