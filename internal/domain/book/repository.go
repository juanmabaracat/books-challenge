package book

type Repository interface {
	Update(Book) error
	Get(string) (*Book, error)
	GetAll() ([]Book, error)
}
