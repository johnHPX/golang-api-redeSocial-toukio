package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type MysqlDB struct {
	DB_UserName string
	DB_Password string
	DB_Name     string
}

var port = 0

func Load() (string, int) {
	databaseConfig := MysqlDB{
		DB_UserName: os.Getenv("DB_USER"),
		DB_Password: os.Getenv(""),
		DB_Name:     os.Getenv("DB_NOME"),
	}
	err := godotenv.Load() // verificar se existem um arquivo com variaveis de ambiente
	if err != nil {
		log.Fatal(err)
	}

	port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal("Porta do servidor não definida")
	}

	StringConnctBD := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", databaseConfig.DB_UserName, databaseConfig.DB_Password, databaseConfig.DB_Name)

	return StringConnctBD, port
}
