package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"reflect"
	"strings"
	"test/internal/domain/entity"
	customError "test/pkg/error"
	"time"
)

type Author interface {
	GetAuthors() ([]entity.Author, error)
	GetAuthorById(ctx context.Context, id int) (entity.Author, error)
	CreateAuthor(author entity.Author) error
	DeleteAuthor(authorID int) error
	UpdateAuthor(author entity.Author) error
	GetAuthorBooks(authorID int) ([]entity.Book, error)
}

type AuthorStore struct {
	db *sqlx.DB
}

func NewAuthorStore(db *sqlx.DB) *AuthorStore {
	return &AuthorStore{
		db: db,
	}
}

func (a *AuthorStore) GetAuthors() ([]entity.Author, error) {
	query := `SELECT author_id, full_name, nick, speciality
	FROM author`

	var authors []entity.Author

	if err := a.db.Select(&authors, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return authors, fmt.Errorf("store: GetAutors(): %w", customError.ErrNothingToFound)
		}
		return nil, fmt.Errorf("store: GetAutors(): %w", err)
	}

	return authors, nil
}

func (a *AuthorStore) GetAuthorById(c context.Context, id int) (entity.Author, error) {

	ctx, cancel := context.WithTimeout(c, time.Second*3)

	defer cancel()

	query := `SELECT author_id, full_name, nick, speciality
	FROM author
	WHERE author_id = $1`

	var author entity.Author

	err := a.db.GetContext(ctx, &author, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return author, fmt.Errorf("store: GetAutors(): %w", customError.ErrNothingToFound)
		}
		return author, fmt.Errorf("store: GetAutors(): %w", err)
	}

	return author, nil
}

func (a *AuthorStore) CreateAuthor(author entity.Author) error {
	query := `INSERT INTO author (full_name, nick, speciality) 
	VALUES (:full_name, :nick, :speciality)`

	_, err := a.db.NamedExec(query, author)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Printf("postgres error: %v\n", pqErr)
			return err
		}
		return fmt.Errorf("store: CreateAuthor(): %w", err)
	}

	return nil
}

func (a *AuthorStore) DeleteAuthor(id int) error {
	query := `DELETE FROM author WHERE author_id = $1`

	_, err := a.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("store: DeleteAuthor(): %w", err)
	}

	return nil
}

func (a *AuthorStore) UpdateAuthor(author entity.Author) error {
	if err := checkForExists(author.ID, authorTable, a.db); err != nil {
		return fmt.Errorf("store: UpdateAuthor(): %w", customError.ErrNothingToFound)
	}

	updates := make([]string, 0, reflect.TypeOf(author).NumField())

	if author.FullName != "" || strings.TrimSpace(author.FullName) != "" {
		updates = append(updates, fmt.Sprintf(`full_name = '%s'`, author.FullName))
	}

	if author.Nick != "" || strings.TrimSpace(author.Nick) != "" {
		updates = append(updates, fmt.Sprintf(`nick = '%s'`, author.Nick))
	}

	if author.Speciality != "" || strings.TrimSpace(author.Speciality) != "" {
		updates = append(updates, fmt.Sprintf(`speciality = '%s'`, author.Speciality))
	}

	if len(updates) == 0 {
		return fmt.Errorf("store: UpdateAuthor(): %w", customError.ErrNothingToUpdate)
	}

	values := strings.Join(updates, ", ")
	log.Println(values)

	query := fmt.Sprintf(`UPDATE author SET %s WHERE author_id = $1`, values)

	log.Println(query)

	_, err := a.db.Exec(query, author.ID)
	if err != nil {
		return fmt.Errorf("store: UpdateAuthor(): %w", err)
	}

	// rows, err := res.RowsAffected()
	// if err != nil {

	// }
	// log.Printf("rows updated: %d\n", rows)

	return nil
}

func (a *AuthorStore) GetAuthorBooks(authorID int) ([]entity.Book, error) {
	var books []entity.Book

	query := `SELECT book_id, name, genre, code_isbn
	FROM book
	WHERE author_id = $1`

	if err := a.db.Select(&books, query, authorID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return books, fmt.Errorf("store: GetAuthorBooks(): %w", customError.ErrNothingToFound)
		}
		return books, fmt.Errorf("store: GetAuthorBooks(): %w", err)
	}

	return books, nil
}
