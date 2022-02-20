package main

import (
	"bufio"
	"fmt"
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
			Usage:  "Cria tebelas",
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
	valorPassado := c.String("migration")
	fmt.Println(valorPassado)
	readFile, err := os.Open(fmt.Sprintf(".../.../migrations/%s", valorPassado))

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	readFile.Close()

}

func migrationDown(c *cli.Context) {}

func main() {
	aplicação := gerarAPP()
	err := aplicação.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
