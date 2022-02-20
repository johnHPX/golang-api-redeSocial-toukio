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

func configurar(r *mux.Router) *mux.Router {
	routers := routerUsers

	for _, rota := range routers {
		if rota.ReqAuthentication {
			r.HandleFunc(rota.Path, appl.Autenticar(rota.Handler)).Methods(rota.Method)
		}
		r.HandleFunc(rota.Path, appl.Loggar(rota.Handler)).Methods(rota.Method)
	}

	return r
}

func Generate() *mux.Router {
	r := mux.NewRouter()
	return configurar(r)
}
