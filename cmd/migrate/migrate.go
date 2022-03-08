package main

/*
	AINDA TA EM PERIODO DE TESTE! NÃO ESTÁ FUNCIONANDO AINDA

*/

import (
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
			Name:   "migrate",
			Usage:  "Criar tebelas/deletar tabelas",
			Flags:  flag,
			Action: migration,
		},
	}

	return app
}

// função que executa o migrate
func migration(c *cli.Context) {

}

func main() {
	aplicação := gerarAPP()
	err := aplicação.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
