package service

import (
	"fmt"
	"test/internal/domain/entity"
	"test/internal/store"
)

type ReaderService struct {
	s store.Reader
}

func NewReaderService(s store.Reader) *ReaderService {
	return &ReaderService{
		s: s,
	}
}

func (r *ReaderService) GetReaderById(id int) (entity.Reader, error) {
	return r.s.GetReaderById(id)
}

func (r *ReaderService) CreateReader(reader entity.Reader) error {
	if err := validateReader(reader); err != nil {
		return fmt.Errorf("service: CreateAuthor(): %w", err)
	}

	return r.s.CreateReader(reader)
}

func (r *ReaderService) UpdateReader(reader entity.Reader) error {
	return r.s.UpdateReader(reader)
}

func (r *ReaderService) DeleteReader(id int) error {
	return r.s.DeleteReader(id)
}

func (r *ReaderService) TakeBook(readerId, bookId int) error {
	return r.s.TakeBook(readerId, bookId)
}

func (r *ReaderService) GetReaderBookList(readerId int) ([]entity.ReaderBookList, error) {
	return r.s.GetReaderBookList(readerId)
}
