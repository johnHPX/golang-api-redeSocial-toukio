package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Path              string
	Method            string
	Handler           func(http.ResponseWriter, *http.Request)
	ReqAuthentication bool
}

func configurar(r *mux.Router) *mux.Router {
	routers := routerUsers

	for _, rota := range routers {

		r.HandleFunc(rota.Path, rota.Handler).Methods(rota.Method)
	}

	return r
}

func Generate() *mux.Router {
	r := mux.NewRouter()
	return configurar(r)
}
