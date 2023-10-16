package book

import (
	"encoding/json"
	"fmt"
	"github.com/juanmabaracat/books-challenge/internal/app/services"
	domain "github.com/juanmabaracat/books-challenge/internal/domain/book"
	"log/slog"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	dateFormat      = "02-01-2006"
	NameParam       = "bookName"
	booksOrderParam = "order"
)

type UpdateBookRequest struct {
	ReleaseDate string `json:"release_date"`
}

func NewHandler(bookServices services.BookService) Handler {
	return Handler{bookServices}
}

type Handler struct {
	BookService services.BookService
}

func (h *Handler) Update(writer http.ResponseWriter, request *http.Request) {
	updateRequest := UpdateBookRequest{}
	decodeErr := json.NewDecoder(request.Body).Decode(&updateRequest)
	if decodeErr != nil {
		slog.Info("couldn't decode update request body", decodeErr)
		handleError(writer, http.StatusBadRequest, "wrong body")
		return
	}

	updateRequest.ReleaseDate = strings.TrimSpace(updateRequest.ReleaseDate)
	if updateRequest.ReleaseDate == "" {
		handleError(writer, http.StatusBadRequest, "release date cannot be empty")
		return
	}

	formattedReleaseDate, parseErr := time.Parse(dateFormat, updateRequest.ReleaseDate)
	if parseErr != nil {
		slog.Error("error parsing date", "error", parseErr.Error())
		handleError(writer, http.StatusBadRequest, "wrong date format")
		return
	}

	if time.Now().Before(formattedReleaseDate) {
		handleError(writer, http.StatusBadRequest, "release date cannot be greater than current date")
		return
	}

	bookName := request.Context().Value(NameParam).(string)
	bookFound, err := h.BookService.Get(bookName)
	if err != nil {
		handleError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	book := domain.Book{
		Name:        bookName,
		ReleaseDate: formattedReleaseDate.Format(dateFormat),
	}

	updateErr := h.BookService.Update(book)
	if updateErr != nil {
		handleError(writer, http.StatusInternalServerError, updateErr.Error())
		return
	}

	if bookFound == nil {
		slog.Info("Book Created", "name", book.Name)
		writer.WriteHeader(http.StatusCreated)
		return
	}

	slog.Info("Book Updated", "name", book.Name)
	writer.WriteHeader(http.StatusNoContent)
	return
}

func (h *Handler) Get(writer http.ResponseWriter, request *http.Request) {
	bookName := request.Context().Value(NameParam).(string)
	book, err := h.BookService.Get(bookName)
	if err != nil {
		handleError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	encodeErr := json.NewEncoder(writer).Encode(book)
	if encodeErr != nil {
		slog.Error("error encoding book", "book", book)
		handleError(writer, http.StatusInternalServerError, encodeErr.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	return
}

func (h *Handler) GetAll(writer http.ResponseWriter, request *http.Request) {
	order := request.URL.Query().Get(booksOrderParam)
	order = strings.TrimSpace(order)
	if order != "" && order != "asc" && order != "desc" {
		handleError(writer, http.StatusBadRequest, "invalid order param")
	}

	books, err := h.BookService.GetAll()
	if err != nil {
		handleError(writer, http.StatusInternalServerError, err.Error())
	}

	if order != "" {
		sort.Slice(books, func(i, j int) bool {
			dateI, _ := time.Parse(dateFormat, books[i].ReleaseDate)
			dateJ, _ := time.Parse(dateFormat, books[j].ReleaseDate)
			if order == "asc" {
				return dateI.Before(dateJ)
			}
			return dateI.After(dateJ)
		})
	}

	encodeErr := json.NewEncoder(writer).Encode(books)
	if encodeErr != nil {
		slog.Error("error encoding book", "books", books)
		handleError(writer, http.StatusInternalServerError, encodeErr.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	return
}

func handleError(writer http.ResponseWriter, statusCode int, msg string) {
	writer.WriteHeader(statusCode)
	_, err := fmt.Fprintf(writer, msg)
	if err != nil {
		slog.Error("error writing response", err)
	}
}
