package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func StartDatabase(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("error starting database %v", err)
	}

	err = createTables()
	if err != nil {
		return fmt.Errorf("error creating the tables %v", err)
	}

	return nil
}

func createTables() error {
	_, err := DB.ExecContext(
		context.Background(),
		"CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, username VARCHAR(50) NOT NULL, last_interaction DATETIME NOT NULL)",
	)
	if err != nil {
		return err
	}

	_, err = DB.ExecContext(
		context.Background(),
		"CREATE TABLE IF NOT EXISTS interactions(id INTEGER PRIMARY KEY AUTOINCREMENT, interaction_type VARCHAR(50) NOT NULL, interaction_to VARCHAR(50), interaction_time DATETIME NOT NULL, user_id INTEGER, FOREIGN KEY (user_id) REFERENCES users(id))",
	)
	if err != nil {
		return err
	}

	return nil
}
