package router

import (
	"api/router/rotas"

	"github.com/gorilla/mux"
)

// Rertorna um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
