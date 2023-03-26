package rotas

import (
	"api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct { //Struct para criação de rotas da API
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router { //Coloca as rotas no router
	rotas := rotasMusicas

	for _, rota := range rotas { //itera pelas rotas de usuários e joga dentro do router por meio do HandleFunc

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
			//r.HandleFunc(rota.Uri, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo) //verifica se está autenticado, caso positivo chama o método
		} else {
			r.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	return r
}
