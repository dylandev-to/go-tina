package queries

import (
	"context"
	"go-tina/internal/database"
	"go-tina/internal/models"
	"strconv"
)

func GetUser(userID string, username string) (models.User, error) {
	var user models.User

	idParsed, err := strconv.Atoi(userID)
	if err != nil {
		return user, nil
	}

	row := database.DB.QueryRowContext(context.Background(),
		"SELECT * FROM users WHERE id=?", idParsed)

	err = row.Scan(&user.ID, &user.Username, &user.LastInteraction)
	if err != nil {
		_, _ = database.DB.ExecContext(
			context.Background(),
			`INSERT INTO users (id, username, last_interaction) VALUES (?,?,?);`, userID, username, "2005-09-02 00:00:00",
		)
		return user, err
	}

	return user, nil
}
