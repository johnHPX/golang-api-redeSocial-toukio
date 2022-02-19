package pgclient

import (
	"API-RS-TOUKIO/configs"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connectar() (*sql.DB, error) {
	strDb, _ := configs.Load()
	db, err := sql.Open("mysql", strDb)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}
