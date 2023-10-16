package services

import (
	"github.com/juanmabaracat/books-challenge/internal/domain/book"
	"reflect"
	"testing"
)

func TestNewBookService(t *testing.T) {
	type args struct {
		repository book.Repository
	}
	tests := []struct {
		name string
		args args
		want BookService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_Get(t *testing.T) {
	type fields struct {
		repository book.Repository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *book.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookService{
				repository: tt.fields.repository,
			}
			got, err := s.Get(tt.args.name)
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
	type fields struct {
		repository book.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []book.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookService{
				repository: tt.fields.repository,
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
	type fields struct {
		repository book.Repository
	}
	type args struct {
		book book.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookService{
				repository: tt.fields.repository,
			}
			if err := s.Update(tt.args.book); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
