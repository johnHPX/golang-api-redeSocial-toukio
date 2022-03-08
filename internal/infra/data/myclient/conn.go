package myclient

import (
	"API-RS-TOUKIO/configs"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect -> faz a conexão com o banco de dados e a retorna
func Connect() (*sql.DB, error) {
	strDb, _ := configs.LoadOS()        //pegando a string de endereco do mysql
	db, err := sql.Open("mysql", strDb) //abrindo a conexão
	if err != nil {
		return nil, err
	}

	err = db.Ping() // verificando se a conexão foi bem sucedida
	if err != nil {
		return nil, err
	}

	return db, nil //retornando a conexão

}
