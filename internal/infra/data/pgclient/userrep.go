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

	// para rota de login
	if tipo == "login" {
		err := rows.Scan(
			&id,
			&password,
		)

		if err != nil {
			return nil, err
		}
	}
	// para rota de listarALL
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
	// para outras rotas de listar
	if tipo == "outersList" {
		err := rows.Scan(
			&id,
			&name,
			&nick,
			&email,
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
	// cada "?" representa um valor a ser adiconado
	sqlText := "select id, name, nick, email, create_at from users where name like ? or nick like ?"
	rows, err := db.Query(sqlText, NameOrNick, NameOrNick) // atibuindo valor ao "?", e retornado os valores
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]users.Entity, 0)

	for rows.Next() {
		ent, err := userImpl.scan("outersList", rows)
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

	sqlText := "select id, name, nick, email, create_at from users where id = ?"
	row, err := db.Query(sqlText, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() { // não precisa percorrer com o for, pois é apenas um usuario
		return userImpl.scan("otherlists", row) // verificar com o scan
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

// busca uma senha de um usuario salvo no banco atraves do seu email
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

// Permite que um usuario siga o outro
func (userImpl *userRepositoryImpl) FollowUser(userID, followerID int64) error {
	db, err := Connectar()
	if err != nil {
		return nil
	}

	sqlText := "INSERT IGNORE INTO followers(user_id, follower_id) VALUES (?,?)"
	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

// permite que um usuario pare de seguir outro
func (userImpl *userRepositoryImpl) StopFollowing(userID, followerID int64) error {
	db, err := Connectar()
	if err != nil {
		return nil
	}

	sqlText := "DELETE FROM followers WHERE user_id = ? AND follower_id = ?"

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

// traz todos os seguidores de um usuario
func (userImpl *userRepositoryImpl) SearchFollowers(userID int64) ([]users.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}

	sqlText := "select u.id, u.name, u.nick, u.email, u.create_at from users u inner join followers s on u.id = s.follower_id where s.user_id = ?"

	rows, err := db.Query(sqlText, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []users.Entity
	for rows.Next() {
		ent, err := userImpl.scan("outersList", rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *ent)
	}

	return result, nil
}

// traz todos os usuarios de um determinado usuario está seguindo
func (userImpl *userRepositoryImpl) SearchFollowing(userID int64) ([]users.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}

	sqlText := "select u.id, u.name, u.nick, u.email, u.create_at from users u inner join followers s on u.id = s.user_id where s.follower_id = ?"

	rows, err := db.Query(sqlText, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []users.Entity
	for rows.Next() {
		ent, err := userImpl.scan("outersList", rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *ent)
	}

	return result, nil
}

func (userImpl *userRepositoryImpl) SearchPassword(userID int64) (string, error) {
	db, err := Connectar()
	if err != nil {
		return "", nil
	}

	sqlText := "select password from users where id = ?"

	rows, err := db.Query(sqlText, userID)
	if err != nil {
		return "", nil
	}

	defer rows.Close()

	var user users.Entity
	for rows.Next() {
		err := rows.Scan(&user.Password)
		if err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (userImpl *userRepositoryImpl) UpdatePassword(userID int64, password string) error {
	db, err := Connectar()
	if err != nil {
		return err
	}

	sqlText := "update users set password = ? where id = ?"
	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(password, userID)
	if err != nil {
		return err
	}

	return nil
}

// função reponsavel por Retornar todos os metodos do repositorio de users
func NewUserRepository() users.Repository {
	return &userRepositoryImpl{}
}
