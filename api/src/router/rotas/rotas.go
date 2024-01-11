package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota reprenta todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.funcao).Methods(rota.Metodo)
	}

	return r
}
