package configs

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Configuraçõs do banco de dados: mysql
type Mysqldb struct {
}

var (
	db_UserName string //nome de usuario
	db_Password string //senha de usuario
	db_Name     string //nome do banco

	port      = 0    // porta do servidor, inicialmente está com o valor zero
	SecretKey []byte // chave secreta, que será gerada pela api
)

// gera um valor para SecretKey, que será usada para asinar o token
func init() {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	SecretKey = []byte(stringBase64)
}

// Load -> vai configurar a string de endereço do mysql e irá definir a porta do servidor,
// os valores vão ser coletados do arquivo ".env",
// que contem as variaveis de ambiente que podem ser modificadas depedendo da necessidada.
func LoadOS() (string, int) {
	err := godotenv.Load() // vai verificar se existe um arquivo com variaveis de ambiente
	if err != nil {
		log.Fatal(err)
	}

	db_Name = os.Getenv("DB_NAME")
	db_UserName = os.Getenv("DB_USER")
	db_Password = os.Getenv("DB_PASSWORD")

	port, err = strconv.Atoi(os.Getenv("API_PORT")) // atribuindo valor na porta do servidor
	if err != nil {
		log.Fatal("Porta do servidor não definida")
	}

	StringConnctBD := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", db_UserName, db_Password, db_Name) // string de endereço do banco mysql

	return StringConnctBD, port //retorna a strig e a porta
}
