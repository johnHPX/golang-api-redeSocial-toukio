package main

import (
	"API-RS-TOUKIO/configs"
	"API-RS-TOUKIO/internal/interf/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	_, port := configs.Load()
	r := routers.Generate()

	fmt.Printf("Escutando na porta %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))

}
