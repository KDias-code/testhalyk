package store

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)


const (
	authorTable = "author"
)
type Store struct {
	Author
	Book
	Reader
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		Author: NewAuthorStore(db),
		Book:   NewBookStore(db),
		Reader: NewReaderStore(db),
	}
}

func checkForExists(id int, tableName string, db *sqlx.DB) error {
	var existsID int

	query := fmt.Sprintf(`SELECT %s_id FROM %s WHERE %s_id = $1`, tableName, tableName, tableName)
	// fmt.Println(query)
	if err := db.Get(&existsID, query, id); err != nil {
		log.Printf("store: checkForExists(): %s table: %s", tableName, err.Error())
		return err
	}
	return nil
}
