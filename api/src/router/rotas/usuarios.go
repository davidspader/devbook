package rotas

import "net/http"

var rotasUsuarios = []Rota{
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
}
