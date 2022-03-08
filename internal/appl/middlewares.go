package appl

import (
	"API-RS-TOUKIO/internal/infra/data/response"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Loggar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validateToken(r)
		if err != nil {
			response.Err(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
