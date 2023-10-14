package memory

import (
	"github.com/juanmabaracat/books-challenge/internal/domain/book"
)

type Repository struct {
	books map[string]book.Book
}

func (r *Repository) Update(book book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetAll() ([]book.Book, error) {
	//TODO implement me
	panic("implement me")
}
