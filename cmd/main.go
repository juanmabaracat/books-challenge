package main

import (
	"github.com/juanmabaracat/books-challenge/internal/app/services"
	"github.com/juanmabaracat/books-challenge/internal/infrastructure/http"
	"github.com/juanmabaracat/books-challenge/internal/infrastructure/storage/memory"
	"log"
)

func main() {
	bookRepository := memory.NewRepository()
	bookServices := services.NewBookService(bookRepository)
	server := http.NewServer(bookServices)

	log.Fatal(server.Run(":8080"))
}
