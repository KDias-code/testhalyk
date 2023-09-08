package service

import (
	"context"
	"fmt"
	"test/internal/domain/entity"
	"test/internal/store"

	_ "github.com/lib/pq"
)

type AuthorService struct {
	s store.Author
}

func NewAuthorService(s store.Author) *AuthorService {
	return &AuthorService{
		s: s,
	}
}

func (a *AuthorService) GetAuthors() ([]entity.Author, error) {
	return a.s.GetAuthors()
}

func (a *AuthorService) GetAuthorById(ctx context.Context, id int) (entity.Author, error) {
	return a.s.GetAuthorById(ctx, id)
}

func (a *AuthorService) CreateAuthor(author entity.Author) error {
	if err := validateAuthor(author); err != nil {
		return fmt.Errorf("service: CreateAuthor(): %w", err)
	}

	return a.s.CreateAuthor(author)
}

func (a *AuthorService) DeleteAuthor(id int) error {
	return a.s.DeleteAuthor(id)
}

func (a *AuthorService) UpdateAuthor(author entity.Author) error {
	return a.s.UpdateAuthor(author)
}

func (a *AuthorService) GetAuthorBooks(authorID int) ([]entity.Book, error) {
	return a.s.GetAuthorBooks(authorID)
}
