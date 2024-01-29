package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	funcao:             controllers.Login,
	RequerAutenticacao: false,
}
