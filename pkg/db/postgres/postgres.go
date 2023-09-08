package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"test/config"

	_ "github.com/lib/pq"
)

const schema = `CREATE TABLE IF NOT EXISTS author(
		author_id SERIAL PRIMARY KEY,	
		full_name VARCHAR NOT NULL,
		nick VARCHAR NOT NULL,
		speciality VARCHAR NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS book(
		book_id SERIAL PRIMARY KEY,
		author_id INT NOT NULL,
		name VARCHAR NOT NULL,
		genre VARCHAR NOT NULL,
		code_isbn VARCHAR NOT NULL,

		FOREIGN KEY (author_id) REFERENCES author(author_id)
	);
	
	CREATE TABLE IF NOT EXISTS reader(
		reader_id SERIAL PRIMARY KEY,
		full_name VARCHAR NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS reader_book(
		reader_id INT NOT NULL,
		book_id INT NOT NULL,

		FOREIGN KEY(reader_id) REFERENCES reader(reader_id),
		FOREIGN KEY(book_id) REFERENCES book(book_id)
	);`

func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	dsn1 := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	// dsn2 := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode)

	// db, err := sqlx.Open("postgres", dsn2)

	fmt.Println(dsn1)
	db, err := sqlx.Open("postgres", dsn1)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		// fmt.Println(22222)
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
