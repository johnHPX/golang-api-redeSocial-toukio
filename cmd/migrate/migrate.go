package main

/*
	AINDA TA EM PERIODO DE TESTE! NÃO ESTÁ FUNCIONANDO AINDA

*/

import (
	"API-RS-TOUKIO/internal/infra/data/pgclient"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

// GerarAPP vai retorna uma aplicação de linha de comando pronto para ser exercutada
func gerarAPP() *cli.App {
	app := cli.NewApp()
	app.Name = "Migration amador"
	app.Usage = "Manipula os dados dos banco de dados, podendo criar e deletar tabelas"
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "m",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "up",
			Usage:  "Criar tebelas",
			Flags:  flag,
			Action: migrationUP,
		},
		{
			Name:   "down",
			Usage:  "deleta tabelas",
			Flags:  flag,
			Action: migrationDown,
		},
	}

	return app
}

func migrationUP(c *cli.Context) {
	valorPassado := c.String("m")

	content, err := ioutil.ReadFile("migrations/" + valorPassado)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	fmt.Println(text)

	db, err := pgclient.Connectar()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(">>>>>>>>")

	_, err = db.Exec(text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("migration executed successfully!")

}

func migrationDown(c *cli.Context) {}

func main() {
	aplicação := gerarAPP()
	err := aplicação.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
