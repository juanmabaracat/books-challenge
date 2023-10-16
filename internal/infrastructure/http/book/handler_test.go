package book

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/juanmabaracat/books-challenge/internal/domain/book"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	Book book.Book
}

func (m *mockService) Update(book book.Book) error {
	if book.Name == "error" {
		return errors.New("service error")
	}
	m.Book = book
	return nil
}

func (m *mockService) Get(name string) (*book.Book, error) {
	if name == "error" {
		return nil, errors.New("service error")
	}

	if name == "empty" {
		return nil, nil
	}

	aBook := book.Book{
		Name:        "TestBook",
		ReleaseDate: "14-10-1990",
	}

	return &aBook, nil
}

func (m *mockService) GetAll() ([]book.Book, error) {
	return nil, nil
}

func TestHandler_Get(t *testing.T) {
	tests := []struct {
		name        string
		BookService mockService
		bookName    string
		request     *http.Request
		wantErr     bool
		want        string
		statusCode  int
	}{
		{
			name:        "return 204 No Content when update the book",
			BookService: mockService{},
			bookName:    "TestBook",
			request:     httptest.NewRequest(http.MethodGet, "/books/TestBook", nil),
			wantErr:     false,
			want:        `{"Name":"TestBook","ReleaseDate":"14-10-1990"}` + "\n",
			statusCode:  200,
		},
		{
			name:        "return an error when book service fails",
			BookService: mockService{},
			bookName:    "error",
			request:     httptest.NewRequest(http.MethodGet, "/books/error", nil),
			wantErr:     true,
			want:        "service error",
			statusCode:  500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{&tt.BookService}
			ctx := context.WithValue(tt.request.Context(), NameParam, tt.bookName)
			tt.request = tt.request.WithContext(ctx)
			response := httptest.NewRecorder()

			h.Get(response, tt.request)
			got := response.Body.String()
			if got != tt.want {
				t.Fatalf("\n GOT=%v\nWANT=%v", got, tt.want)
			}

			if response.Code != tt.statusCode {
				t.Fatalf("\n GOT=%v\nWANT=%v", response.Code, tt.statusCode)
			}
		})
	}
}

func TestHandler_Update(t *testing.T) {
	tests := []struct {
		name        string
		BookService *mockService
		bookName    string
		request     *http.Request
		statusCode  int
		want        *book.Book
		wantErr     bool
	}{
		{
			name:        "return server error when the service fails",
			BookService: &mockService{},
			bookName:    "error",
			request:     createRequest(http.MethodPut, UpdateBookRequest{"14-10-2010"}),
			statusCode:  500,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "return 201 Created when the book doesn't exists",
			BookService: &mockService{},
			bookName:    "empty",
			request:     createRequest(http.MethodPut, UpdateBookRequest{"14-10-1990"}),
			statusCode:  201,
			want: &book.Book{
				Name:        "empty",
				ReleaseDate: "14-10-1990",
			},
			wantErr: false,
		},
		{
			name:        "return 204 No Content when update a book without error",
			BookService: &mockService{},
			bookName:    "TestBook",
			request:     createRequest(http.MethodPut, UpdateBookRequest{"14-10-1990"}),
			statusCode:  204,
			want: &book.Book{
				Name:        "TestBook",
				ReleaseDate: "14-10-1990",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				BookService: tt.BookService,
			}

			ctx := context.WithValue(tt.request.Context(), NameParam, tt.bookName)
			tt.request = tt.request.WithContext(ctx)
			response := httptest.NewRecorder()
			h.Update(response, tt.request)
			if !tt.wantErr && tt.BookService.Book != *tt.want {
				t.Fatalf("\n GOT=%v\nWANT=%v", tt.BookService.Book.Name, tt.want)
			}

			if response.Code != tt.statusCode {
				t.Fatalf("\n GOT=%v\nWANT=%v", response.Code, tt.statusCode)
			}
		})
	}
}

func createRequest(method string, data interface{}) *http.Request {
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(data)
	req := httptest.NewRequest(method, "/books/TestBook", body)
	return req
}
