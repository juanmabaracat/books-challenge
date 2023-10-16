package http

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/juanmabaracat/books-challenge/internal/app/services"
	"github.com/juanmabaracat/books-challenge/internal/infrastructure/http/book"
	"log"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	services services.BookService
	router   chi.Router
}

func NewServer(services services.BookService) Server {
	httpServer := Server{services: services}
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(commonMiddleware)
	httpServer.router = r
	httpServer.addBookHttpRoutes()

	return httpServer
}

func (server *Server) addBookHttpRoutes() {
	handler := book.NewHandler(server.services)

	server.router.Route("/books", func(r chi.Router) {
		r.Get("/", handler.GetAll)
	})

	server.router.Route("/books/{"+book.NameParam+"}", func(r chi.Router) {
		r.Use(BookContext)
		r.Get("/", handler.Get)
		r.Put("/", handler.Update)
	})
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}

func BookContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		bookName := chi.URLParam(request, book.NameParam)
		if strings.TrimSpace(bookName) == "" {
			http.Error(writer, "book name cannot be empty", http.StatusBadRequest)
		}

		ctx := context.WithValue(request.Context(), book.NameParam, bookName)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func (server *Server) Run(port string) error {
	log.Println("Listening on http://localhost:8080")
	return http.ListenAndServe(port, server.router)
}
