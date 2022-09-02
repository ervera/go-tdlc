package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection() *sql.DB {
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	// SIN EL var err error, no funciona.
	DSN := os.Getenv("PLANET_SCALE_MYSQL")
	StorageDB, err := sql.Open("mysql", DSN)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")
	return StorageDB
}
