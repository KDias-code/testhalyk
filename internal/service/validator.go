package service

import (
	"strings"
	"test/internal/domain/entity"
	customError "test/pkg/error"
)

func validateAuthor(author entity.Author) error {

	if author.FullName == "" || strings.TrimSpace(author.FullName) == "" {
		return customError.ErrEmptyFullName
	}

	if author.Nick == "" || strings.TrimSpace(author.Nick) == "" {
		return customError.ErrEmptyFields
	}

	if author.Speciality == "" || strings.TrimSpace(author.Speciality) == "" {
		return customError.ErrEmptyFields
	}

	return nil
}

func validateBook(book entity.Book) error {

	if book.Name == "" || strings.TrimSpace(book.Name) == "" {
		return customError.ErrEmptyFields
	}

	if book.Genre == "" || strings.TrimSpace(book.Genre) == "" {
		return customError.ErrEmptyFields
	}

	if book.CodeISBN == "" || strings.TrimSpace(book.CodeISBN) == "" {
		return customError.ErrEmptyFields
	}

	return nil
}

func validateReader(reader entity.Reader) error {
	if reader.FullName == "" || strings.TrimSpace(reader.FullName) == "" {
		return customError.ErrEmptyFields
	}

	return nil
}
