package main

import (
	"API-RS-TOUKIO/configs"
	"API-RS-TOUKIO/internal/interf/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	_, port := configs.Load() // pegando a porta do servidor
	r := routers.Generate()   // gerando as rotas da API

	fmt.Printf("Escutando na porta %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r)) //inciando o servidor

}
