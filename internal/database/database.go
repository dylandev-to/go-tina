package database

import (
	"database/sql"
	"fmt"
	"go-tina/pkg/utils"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func StartDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path.Join(utils.GetCwd(), "db", "bot.db"))
	if err != nil {
		return nil, fmt.Errorf("error starting database %v", err)
	}

	return db, nil
}
