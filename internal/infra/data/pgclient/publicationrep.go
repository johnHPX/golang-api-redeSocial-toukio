package pgclient

import (
	"API-RS-TOUKIO/internal/domain/publication"
	"database/sql"
	"errors"
)

type publicationRepositoryImpl struct{} /// objeto que vai servi para a implementação dos metodos do repositorio de publication

// verifica se os tipos dos dados do request estão de acordo com os da tabela publication
func (publicationImpl *publicationRepositoryImpl) scan(rows *sql.Rows) (*publication.Entity, error) {
	id := sql.NullInt64{}
	title := sql.NullString{}
	Content := sql.NullString{}
	AuthorID := sql.NullInt64{}
	Likes := sql.NullInt64{}
	create_at := sql.NullTime{}
	Author_nick := sql.NullString{}

	err := rows.Scan(
		&id,
		&title,
		&Content,
		&AuthorID,
		&Likes,
		&create_at,
		&Author_nick,
	)

	if err != nil {
		return nil, err
	}

	ent := new(publication.Entity)
	if id.Valid {
		ent.ID = id.Int64
	}

	if title.Valid {
		ent.Title = title.String
	}

	if Content.Valid {
		ent.Content = Content.String
	}

	if AuthorID.Valid {
		ent.AuthorID = AuthorID.Int64
	}

	if Likes.Valid {
		ent.Likes = Likes.Int64
	}

	if create_at.Valid {
		ent.Create_at = create_at.Time
	}

	if Author_nick.Valid {
		ent.AuthorNick = Author_nick.String
	}

	return ent, nil

}

// criar uma publicação
func (publicationImpl *publicationRepositoryImpl) CreatePublication(e *publication.Entity) (int64, error) {
	db, err := Connectar() // abrir conexão com o banco
	if err != nil {
		return 0, err // trata o erro
	}

	defer db.Close() // fecha a conexão

	sqlText := "INSERT INTO publication (title, content, author_id) VALUES (?,?,?)" // comando sql

	statement, err := db.Prepare(sqlText) // cria um statement
	if err != nil {
		return 0, err // trata o erro
	}

	result, err := statement.Exec(e.Title, e.Content, e.AuthorID) // executa o statement
	if err != nil {
		return 0, err //trata o erro
	}

	defer statement.Close() // fecha o statement

	lastIDInsert, err := result.LastInsertId() // pega o ultimo id inserido
	if err != nil {
		return 0, err // trata o erro
	}

	return lastIDInsert, nil // retorna o ultimo id
}

// lista as pulicações de um usuario e as de outros usuarios que esse usuario está seguindo
func (publicationImpl *publicationRepositoryImpl) ListAllPublication(userID int64) ([]publication.Entity, error) {
	db, err := Connectar() //abre a conexão
	if err != nil {
		return nil, err // trata o erro
	}

	db.Close() // fecha a conexão

	sqlText := "select distinct p.*, u.nick from publication p inner join users u on u.id = p.author_id inner join followers s on p.author_id = s.user_id where u.id = ? or s.follower_id = ? order by 1 desc" // comando sql

	rows, err := db.Query(sqlText, userID, userID) // query retorna todas as linhas/dados validos pelo comando sql
	if err != nil {
		return nil, err // trata o erro
	}

	defer rows.Close() // fecha as linhas

	var result []publication.Entity // variavel do tipo da entidade de publicação
	for rows.Next() {               // percorre as linhas/dados validos
		pub, err := publicationImpl.scan(rows) // usa o scan para verificar se os dados estão ok e os atruibui a uma variavel
		if err != nil {
			return nil, err // trata o erro
		}

		result = append(result, *pub) // pegar a variavel da entidade e adicona nela a variavel do scan
	}

	return result, nil // retorna uma variavel do tipo entidade populada
}

// procurar uma publicação atravez do id
func (publicationImpl *publicationRepositoryImpl) FindByIDPublication(publicationID int64) (*publication.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlText := "select p.*, u.nick from publication p inner join users u on u.id = p.author_id where p.id = ?"

	row, err := db.Query(sqlText, publicationID)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		return publicationImpl.scan(row)
	}

	return nil, errors.New("Usuário não foi encontrado!")

}

// atualiza uma publicação de um usuario
func (publicationImpl *publicationRepositoryImpl) UpdatePublication(publicationID int64, e *publication.Entity) error {
	db, err := Connectar()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlText := "update publication set title = ?, content = ? where id = ?"

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Title, e.Content, publicationID)
	if err != nil {
		return err
	}

	return nil
}

// deleta uma publicação de um usuario
func (publicationImpl *publicationRepositoryImpl) DeletePublication(publicationID int64) error {
	db, err := Connectar()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlText := "delete from publication where id = ?"

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(publicationID)
	if err != nil {
		return err
	}

	return nil
}

// traz todas as publicações de um usuario
func (publicationImpl *publicationRepositoryImpl) ListByIDUserPublication(userID int64) ([]publication.Entity, error) {
	db, err := Connectar()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlText := "select p.*, u.nick from publication p join users u on u.id = p.author_id where p.author_id = ?"

	rows, err := db.Query(sqlText, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []publication.Entity
	for rows.Next() {
		ent, err := publicationImpl.scan(rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *ent)
	}

	return result, nil
}

// permite que um usuario curta uma publicação
func (publicationImpl *publicationRepositoryImpl) LikePublication(publicationID int64) error {
	db, err := Connectar()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlText := "update publication set likes = likes + 1 where id = ?"

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

// permite que um usuario descurta uma publicação
func (publicationImpl *publicationRepositoryImpl) DeslikePublication(publicationID int64) error {
	db, err := Connectar()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlText := "update publication set likes = CASE WHEN likes > 0 THEN likes - 1 ELSE 0 END where id = ?"

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

// função responsvavel por retornar todos os metodos implementados
func NewPublicationRepository() publication.Repository {
	return &publicationRepositoryImpl{}
}
