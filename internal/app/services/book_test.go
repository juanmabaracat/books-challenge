package services

import (
	"errors"
	"github.com/juanmabaracat/books-challenge/internal/domain/book"
	"reflect"
	"testing"
)

type mockRepository struct {
	books []book.Book
}

func (m mockRepository) Update(book book.Book) error {
	if book.Name == "error" {
		return errors.New("repository error")
	}

	return nil
}

func (m mockRepository) Get(name string) (*book.Book, error) {
	if name == "error" {
		return nil, errors.New("repository error")
	}

	return &book.Book{
		Name:        name,
		ReleaseDate: "18-12-2022",
	}, nil
}

func (m mockRepository) GetAll() ([]book.Book, error) {
	if m.books[0].Name == "error" {
		return nil, errors.New("repository error")
	}

	return m.books, nil
}

func Test_bookService_Get(t *testing.T) {
	tests := []struct {
		name       string
		repository mockRepository
		bookName   string
		want       *book.Book
		wantErr    bool
	}{
		{
			name:       "return a book without error",
			repository: mockRepository{},
			bookName:   "Test Book",
			want:       &book.Book{Name: "Test Book", ReleaseDate: "18-12-2022"},
			wantErr:    false,
		},
		{
			name:       "return an error when the repository fails",
			repository: mockRepository{},
			bookName:   "error",
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookService{
				repository: tt.repository,
			}
			got, err := s.Get(tt.bookName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_GetAll(t *testing.T) {
	tests := []struct {
		name       string
		repository mockRepository
		want       []book.Book
		wantErr    bool
	}{
		{
			name:       "return error when repository fails",
			repository: mockRepository{createBookList("error")},
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "return the list of book without error",
			repository: mockRepository{createBookList("Test")},
			want:       createBookList("Test"),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookService{
				repository: tt.repository,
			}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_Update(t *testing.T) {
	tests := []struct {
		name       string
		repository mockRepository
		book       book.Book
		wantErr    bool
	}{
		{
			name:       "return error when repository fails",
			repository: mockRepository{},
			book:       book.Book{Name: "error"},
			wantErr:    true,
		},
		{
			name:       "update book without error",
			repository: mockRepository{},
			book:       book.Book{Name: "test", ReleaseDate: "18-12-2022"},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookService{
				repository: tt.repository,
			}
			if err := s.Update(tt.book); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createBookList(firstBook string) []book.Book {
	return []book.Book{{
		Name:        firstBook,
		ReleaseDate: "18-12-2022",
	}, {
		Name:        "Test book 2",
		ReleaseDate: "11-12-2022",
	}}
}
