package service

import (
	"fmt"
	"test/internal/entity"
	"test/internal/store"
)

type BookService struct {
	s store.Book
}

func NewBookService(s store.Book) *BookService {
	return &BookService{
		s: s,
	}
}

func (b *BookService) GetBookById(id int) (entity.Book, error) {
	return b.s.GetBookById(id)
}
func (b *BookService) CreateBook(book entity.Book) error {
	if err := validateBook(book); err != nil {
		return fmt.Errorf("service: CreateBook(): %w", err)
	}

	return b.s.CreateBook(book)
}
func (b *BookService) UpdateBook(book entity.Book) error {
	return b.s.UpdateBook(book)
}
func (b *BookService) DeleteBook(authorID, bookID int) error {
	return b.s.DeleteBook(bookID)
}
