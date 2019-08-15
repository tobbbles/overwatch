package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Store provides a naive implementation of an SQLite3 database. This store doesn't use any ORM; however
// it is highly recommended to use an ORM if you're dealing with any more than two tables.
type Store struct {
	db *sql.DB
}

func New(path string) (*Store, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Commit migration for schema
	if _, err := db.Exec(UpHeros); err != nil {
		return nil, err
	}

	if _, err := db.Exec(UpAbilities); err != nil {
		return nil, err
	}

	s := &Store{
		db: db,
	}

	return s, nil
}

func (s *Store) Close() {
	s.db.Close()
}
