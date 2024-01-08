package rotas

import "net/http"

type Rota struct {
	URI                string
	Metodo             string
	funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}