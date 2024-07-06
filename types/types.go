package types

type BookStore interface {
	NewBook(book Book) error
	GetBook(id int) (*Book, error)
	GetBooks()([]Book, error)
	UpdateBook(book Book) error
	DeleteBook(id int) error
}

type Book struct {
	ID					int `json:"id"`
	ISBN				string `json:"isbn"`
	Author			string `json:"author"`
	Pages				int			`json:"pages"`
}

type BookPayload struct {
	ISBN				string `json:"isbn"`
	Author			string `json:"author"`
	Pages				int			`json:"pages"`
}