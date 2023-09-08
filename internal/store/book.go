package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"test/internal/entity"
	customError "test/pkg/error"

	"github.com/jmoiron/sqlx"
)

type Book interface {
	GetBookById(id int) (entity.Book, error)
	CreateBook(book entity.Book) error
	UpdateBook(book entity.Book) error
	DeleteBook(bookID int) error
}

type BookStore struct {
	db *sqlx.DB
}

func NewBookStore(db *sqlx.DB) *BookStore {
	return &BookStore{
		db: db,
	}
}

func (b *BookStore) GetBookById(id int) (entity.Book, error) {
	query := `SELECT book_id, name, genre, code_isbn, author_id, full_name, nick, speciality
	FROM book
	JOIN author 
	USING (author_id)
	WHERE book_id = $1`

	var book entity.Book

	if err := b.db.Get(&book, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return book, fmt.Errorf("store: GetBookById(): %w", customError.ErrNothingToFound)
		}
		return book, err
	}

	log.Println("OK", book)

	return book, nil
}

func (b *BookStore) CreateBook(book entity.Book) error {
	query := `INSERT INTO book (author_id, name, genre, code_isbn) 
	VALUES (:author_id, :name, :genre, :code_isbn)`

	_, err := b.db.NamedExec(query, book)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookStore) UpdateBook(book entity.Book) error {
	if err := checkForExists(book.ID, "book", b.db); err != nil {
		return fmt.Errorf("store: UpdateBook(): %w", customError.ErrNothingToFound)
	}

	updates := []string{}

	if book.Name != "" || strings.TrimSpace(book.Name) != "" {
		updates = append(updates, fmt.Sprintf(`name = '%s'`, book.Name))
	}

	if book.Genre != "" || strings.TrimSpace(book.Genre) != "" {
		updates = append(updates, fmt.Sprintf(`genre = '%s'`, book.Genre))
	}

	if book.CodeISBN != "" || strings.TrimSpace(book.CodeISBN) != "" {
		updates = append(updates, fmt.Sprintf(`code_isbn = '%s'`, book.CodeISBN))
	}

	if len(updates) == 0 {
		return fmt.Errorf("store: UpdateBook(): %w", customError.ErrNothingToUpdate)
	}

	values := strings.Join(updates, ", ")
	log.Println(values)

	query := fmt.Sprintf(`UPDATE book SET %s WHERE book_id = $1`, values)

	if _, err := b.db.Exec(query, book.ID); err != nil {
		return fmt.Errorf("store: UpdateBook(): %w", err)
	}

	return nil
}

func (b *BookStore) DeleteBook(bookID int) error {
	query := `DELETE FROM book WHERE book_id = $1`

	_, err := b.db.Exec(query, bookID)
	if err != nil {
		return fmt.Errorf("store: DeleteBook(): %w", err)
	}

	return nil
}
