package pgclient

import (
	"API-RS-TOUKIO/internal/domain/users"
	"database/sql"
	"errors"
)

// objeto que vai servi para a implementação dos metodos do repositorio de users
type userRepositoryImpl struct{}

// scan -> Serve para verificar se os dados do request estão de acordo com os seus tipos da tabela users
func (userImpl *userRepositoryImpl) scan(tipo string, rows *sql.Rows) (*users.Entity, error) {
	id := sql.NullInt64{}
	name := sql.NullString{}
	nick := sql.NullString{}
	email := sql.NullString{}
	password := sql.NullString{}
	create_at := sql.NullTime{}

	if tipo == "" {
		return nil, errors.New("não foi definido o tipo de scan")
	}

	if tipo == "login" {
		err := rows.Scan(
			&id,
			&password,
		)

		if err != nil {
			return nil, err
		}
	}

	if tipo == "listALL" {
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

// CreateUser -> metodo de criar um usuario, executando o comando sql de criação de dados
func (userImpl *userRepositoryImpl) CreateUser(e *users.Entity) error {
	db, err := Connectar() // abre a conexão
	if err != nil {
		return err //trata o erro
	}
	defer db.Close() // fecha a conexão por ultimo

	sqlText := "insert into users (name,nick,email,password) values (?,?,?,?)" // comanddo sql
	statement, err := db.Prepare(sqlText)                                      // prepara um statement para a execução do sql
	if err != nil {
		return err // trata o erro
	}
	defer statement.Close() //fecha o statemant por ultimo

	result, err := statement.Exec(e.Name, e.Nick, e.Email, e.Password) // executa o sql e retorna os resultados
	if err != nil {
		return err //trata o erro
	}

	rows, err := result.RowsAffected() // retorna o numero de linhas afetadas
	if err != nil {
		return err // trata o erro
	}

	if rows != 1 { // verifica se tem linhas afetadas
		return errors.New("erro ao cadastrar usuarios") // não tem linhas afetadas, logo não foi possivel criar usuario
	}

	return nil // retorna nil se tudo ocorreu bem
}

// ListALLUser -> lista todos os usuarios cadastrados , exectutando o comando sql de listar dados
func (userImpl *userRepositoryImpl) ListALLUser() ([]users.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlText := "select * from users"
	rows, err := db.Query(sqlText) // Consulta o banco de dados, e retorna um array com os valores encontrados pelo sql
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]users.Entity, 0) //  slice da entidade

	for rows.Next() { // percorre todos valores
		ent, err := userImpl.scan("listALL", rows) // usa o scan  para verificar se os valores estão de acordo e retorna uma entidade populada
		if err != nil {
			return nil, err
		}

		result = append(result, *ent) // adiciona a entidade populada no slice de entidades
	}

	return result, nil // retorna o slice de entidades populado
}

// ListByNameOrNick -> faz uma listagem de usuarios com o nome ou o nick
func (userImpl *userRepositoryImpl) ListByNameOrNickUsers(NameOrNick string) ([]users.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// id, name, nick, email, create_at
	sqlText := "select * from users where name like ? or nick like ?" // cada "?" representa um valor a ser adiconado
	rows, err := db.Query(sqlText, NameOrNick)                        // atibuindo valor ao "?", e retornado os valores
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]users.Entity, 0)

	for rows.Next() {
		ent, err := userImpl.scan("listALL", rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *ent)
	}

	return result, nil
}

// FindUser -> busca um unico usuario, através do id
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

	if row.Next() { // não precisa percorrer com o fro, pois é apenas um usuario
		return userImpl.scan("listALL", row) // verificar com o scan
	}

	return nil, errors.New("Usuário não foi encontrado!")

}

// UpdateUser -> atualiza dados de um usuario
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

// DeletarUser -> deleta um usuario
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

func (userImpl *userRepositoryImpl) SearchforEmail(email string) (*users.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlText := "select id, password from users where email = ?"

	row, err := db.Query(sqlText, email)
	if err != nil {
		return nil, err
	}

	if row.Next() {
		return userImpl.scan("login", row)
	}

	return nil, errors.New("Email não foi encontrado!")

}

// função reponsavel por Retornar todos os metodos do repositorio de users
func NewUserRepository() users.Repository {
	return &userRepositoryImpl{}
}
