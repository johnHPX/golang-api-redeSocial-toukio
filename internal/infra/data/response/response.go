package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON para requisição
func JSON(w http.ResponseWriter, statusCode int, date interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if date != nil {
		erro := json.NewEncoder(w).Encode(date)
		if erro != nil {
			log.Fatal(erro)
		}
	}

}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
