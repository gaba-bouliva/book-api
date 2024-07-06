package book

import (
	"book-api/types"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db,
	}
}

func (s *Store) GetBooks() ([]types.Book, error) {

	return []types.Book{}, nil
}

func (s *Store) GetBook(id int) (*types.Book, error) {

	return &types.Book{}, nil
}

func (s *Store) UpdateBook(book types.Book) error {

	return nil
}

func (s *Store) NewBook(book types.Book) error {

	return nil
}

func (s *Store) DeleteBook(id int) error {

	return nil
}