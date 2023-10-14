package book

type Repository interface {
	Update(Book) error
	GetAll() ([]Book, error)
}
