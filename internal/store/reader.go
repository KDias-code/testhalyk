package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"test/internal/entity"
	customError "test/pkg/error"

	"github.com/jmoiron/sqlx"
)

type Reader interface {
	GetReaderById(id int) (entity.Reader, error)
	CreateReader(reader entity.Reader) error
	UpdateReader(reader entity.Reader) error
	DeleteReader(readerID int) error
	TakeBook(readerId, bookId int) error
	GetReaderBookList(readerId int) ([]entity.ReaderBookList, error)
}

type ReaderStore struct {
	db *sqlx.DB
}

func NewReaderStore(db *sqlx.DB) *ReaderStore {
	return &ReaderStore{
		db: db,
	}
}

func (r *ReaderStore) GetReaderById(id int) (entity.Reader, error) {

	var reader entity.Reader

	query := `SELECT reader_id, full_name
	FROM reader
	WHERE reader_id = $1`

	if err := r.db.Get(&reader, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return reader, fmt.Errorf("store: GetReaderById(): %w", customError.ErrNothingToFound)
		}
		return reader, fmt.Errorf("store: GetReaderById(): %w", err)
	}

	return reader, nil
}

func (r *ReaderStore) CreateReader(reader entity.Reader) error {
	query := `INSERT INTO reader (full_name)
	VALUES (:full_name)`

	_, err := r.db.NamedExec(query, reader)
	if err != nil {
		return fmt.Errorf("store: CreateReader(): %w", err)
	}

	return nil
}

func (r *ReaderStore) UpdateReader(reader entity.Reader) error {
	if err := checkForExists(reader.ID, "reader", r.db); err != nil {
		return fmt.Errorf("store: UpdateReader(): %w", customError.ErrNothingToFound)
	}

	if reader.FullName == "" || strings.TrimSpace(reader.FullName) == "" {
		return fmt.Errorf("store: UpdateReader(): %w", customError.ErrNothingToUpdate)
	}

	query := `UPDATE reader SET full_name = $1 WHERE reader_id = $2`

	_, err := r.db.Exec(query, reader.FullName, reader.ID)
	if err != nil {
		return fmt.Errorf("store: UpdateReader(): %w", err)
	}

	return nil
}

func (r *ReaderStore) DeleteReader(id int) error {
	query := `DELETE FROM reader WHERE reader_id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("store: DeleteReader(): %w", err)
	}

	return nil
}

func (r *ReaderStore) TakeBook(readerId, bookId int) error {

	if err := checkForExists(readerId, "reader", r.db); err != nil {
		return fmt.Errorf("store: TakeBook(): %w", customError.ErrNothingToFound)
	}
	if err := checkForExists(bookId, "book", r.db); err != nil {
		return fmt.Errorf("store: TakeBook(): %w", customError.ErrNothingToFound)
	}

	query := `INSERT INTO reader_book(reader_id, book_id) 
	VALUES($1, $2)`

	_, err := r.db.Exec(query, readerId, bookId)
	if err != nil {
		return fmt.Errorf("store: TakeBook(): %w", err)
	}

	return nil
}

func (r *ReaderStore) GetReaderBookList(readerId int) ([]entity.ReaderBookList, error) {
	query := `SELECT name, genre, code_isbn, full_name, nick, speciality
	FROM reader_book rb
	JOIN book b
	USING (book_id)
	JOIN author a
	USING (author_id)
	WHERE rb.reader_id = $1`

	var readerBooks []entity.TempReaderBookList

	if err := r.db.Select(&readerBooks, query, readerId); err != nil {
		return nil, fmt.Errorf("store: GetReaderBookList(): %w", err)
	}

	if len(readerBooks) == 0 {
		return nil, fmt.Errorf("store: GetReaderBookList(): %w", customError.ErrNothingToFound)
	}

	bookList := make([]entity.ReaderBookList, 0, len(readerBooks))

	for _, book := range readerBooks {
		bookList = append(bookList, book.ToReaderBookList())
	}

	return bookList, nil
}
