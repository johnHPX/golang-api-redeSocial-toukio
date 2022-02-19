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
	create_at := sql.NullTime{}

	err := rows.Scan(
		&id,
		&name,
		&nick,
		&email,
		&password,
		&create_at,
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

	if create_at.Valid {
		ent.Create_at = create_at.Time
	}

	return ent, nil

}

func (userImpl *userRepositoryImpl) CreateUser(e *users.Entity) error {
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

func (userImpl *userRepositoryImpl) ListALLUser() ([]users.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlText := "select * from users"
	rows, err := db.Query(sqlText)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]users.Entity, 0)

	for rows.Next() {
		ent, err := userImpl.scanIterator(rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *ent)
	}

	return result, nil
}

func (userImpl *userRepositoryImpl) ListByNameOrNickUsers(NameOrNick string) ([]users.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// id, name, nick, email, create_at
	sqlText := "select * from users where name like ? or nick like ?"
	rows, err := db.Query(sqlText, NameOrNick, NameOrNick)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]users.Entity, 0)

	for rows.Next() {
		ent, err := userImpl.scanIterator(rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *ent)
	}

	return result, nil
}

func (userImpl *userRepositoryImpl) FindUser(id int64) (*users.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlText := "select * from users where id = ?"
	row, err := db.Query(sqlText, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		return userImpl.scanIterator(row)
	}

	return nil, errors.New("Usuário não foi encontrado!")

}

func (userImpl *userRepositoryImpl) UpdateUser(e *users.Entity) error {
	db, err := Connectar()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlText := "update users set name = ?, nick = ?, email = ? where id = ?"
	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.Name, e.Nick, e.Email, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (userImpl *userRepositoryImpl) DeleteUser(id int64) error {
	db, err := Connectar()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlText := "delete from users where id = ?"
	statement, err := db.Prepare(sqlText)
	if err != nil {
		return nil
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository() users.Repository {
	return &userRepositoryImpl{}
}
