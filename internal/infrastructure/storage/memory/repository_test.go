package memory

import (
	"github.com/juanmabaracat/books-challenge/internal/domain/book"
	"testing"
)

func TestNewRepository(t *testing.T) {
	t.Run("create repository without error", func(t *testing.T) {
		repository := NewRepository()
		if repository == nil {
			t.Errorf("error creating repository")
		}
	})
}

func TestRepository_GetAll(t *testing.T) {
	t.Run("return all books without error", func(t *testing.T) {
		repository := NewRepository()
		book1 := book.Book{
			Name:        "The Alchemist",
			ReleaseDate: "24-01-1988",
		}

		err := repository.Update(book1)
		if err != nil {
			t.Fatalf("error updating %v", err)
		}

		got, _ := repository.GetAll()

		if got[0] != book1 {
			t.Fatalf("GOT=%v, EXPECTED=%v", got, book1)
		}
	})
}

func TestRepository_Update(t *testing.T) {
	t.Run("update book without error", func(t *testing.T) {
		repository := NewRepository()
		book1 := book.Book{
			Name:        "The Alchemist",
			ReleaseDate: "24-01-1988",
		}

		err := repository.Update(book1)
		if err != nil {
			t.Fatalf("error updating %v", err)
		}
	})
}
