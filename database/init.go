package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() error {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`CREATE TABLE playerinfo (
        ip VARCHAR(128) PRIMARY KEY UNIQUE,
        words_guessed INTEGER NULL,
        losses INTEGER NULL
    )`)
	if err != nil {
		return err
	}
	stmt.Exec()
	db.Close()

	return nil
}
