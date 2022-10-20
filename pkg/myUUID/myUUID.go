package myuuid

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func LastInsertUUID(db *sql.DB, tableName string, ID int64, ctx context.Context) (uuid.UUID, error) {
	query := "SELECT uuid FROM " + tableName + " WHERE id = ?"
	row := db.QueryRow(query, ID)
	var UUID uuid.UUID
	err := row.Scan(&UUID)
	if err != nil {
		fmt.Println(err.Error())
		return uuid.UUID{}, err
	}
	return UUID, nil
}
