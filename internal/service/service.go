package service

import (
	"context"
	"test/internal/entity"
	"test/internal/store"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go

type Author interface {
	GetAuthors() ([]entity.Author, error)
	GetAuthorById(ctx context.Context, id int) (entity.Author, error)
	CreateAuthor(author entity.Author) error
	DeleteAuthor(id int) error
	UpdateAuthor(author entity.Author) error
	GetAuthorBooks(authorID int) ([]entity.Book, error)
}

type Book interface {
	GetBookById(id int) (entity.Book, error)
	CreateBook(book entity.Book) error
	UpdateBook(book entity.Book) error
	DeleteBook(authorID, bookID int) error
}

type Reader interface {
	GetReaderById(id int) (entity.Reader, error)
	CreateReader(reader entity.Reader) error
	UpdateReader(reader entity.Reader) error
	DeleteReader(readerID int) error
	TakeBook(readerId, bookId int) error
	GetReaderBookList(readerId int) ([]entity.ReaderBookList, error)
}

type Service struct {
	Author
	Book
	Reader
}

func NewService(s *store.Store) *Service {
	return &Service{
		Author: NewAuthorService(s.Author),
		Book:   NewBookService(s.Book),
		Reader: NewReaderService(s.Reader),
	}
}
