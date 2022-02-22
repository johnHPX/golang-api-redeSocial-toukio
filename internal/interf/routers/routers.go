package routers

import (
	"API-RS-TOUKIO/internal/appl"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Path              string
	Method            string
	Handler           http.HandlerFunc
	ReqAuthentication bool
}

// configura as rotas que precisam de autenticação com o token e as que não
func configurar(r *mux.Router) *mux.Router {
	routers := routerUsers
	routers = append(routers, routerPublication...)

	for _, rota := range routers {
		if rota.ReqAuthentication {
			r.HandleFunc(rota.Path, appl.Loggar(appl.Autenticar(rota.Handler))).Methods(rota.Method)
		}
		r.HandleFunc(rota.Path, appl.Loggar(rota.Handler)).Methods(rota.Method)
	}

	return r
}

// gera as rotas
func Generate() *mux.Router {
	r := mux.NewRouter()
	return configurar(r)
}
