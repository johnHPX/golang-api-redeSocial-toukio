package pgclient

import (
	"API-RS-TOUKIO/internal/domain/users"
	"database/sql"
	"errors"
)

type userRepositoryImpl struct{}

func (userImpl *userRepositoryImpl) scanIterator(rows *sql.Rows) (*users.Entity, error) {
	id := sql.NullInt64{}
	name := sql.NullString{}
	nick := sql.NullString{}
	email := sql.NullString{}
	password := sql.NullString{}

	err := rows.Scan(
		&id,
		&name,
		&nick,
		&email,
		&password,
	)

	if err != nil {
		return nil, err
	}

	ent := new(users.Entity)
	if id.Valid {
		ent.ID = id.Int64
	}

	if name.Valid {
		ent.Name = name.String
	}

	if nick.Valid {
		ent.Nick = nick.String
	}

	if email.Valid {
		ent.Email = email.String
	}

	if password.Valid {
		ent.Password = password.String
	}

	return ent, nil

}

func (uuserImpl *userRepositoryImpl) CreateUser(e *users.Entity) error {
	db, err := Connectar()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlText := "insert into users (name,nick,email,password) values (?,?,?,?)"
	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(e.Name, e.Nick, e.Email, e.Password)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("erro ao cadastrar usuarios")
	}

	return nil
}

func NewUserRepository() users.Repository {
	return &userRepositoryImpl{}
}
