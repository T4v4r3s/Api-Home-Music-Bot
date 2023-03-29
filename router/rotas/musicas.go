package rotas

import (
	"api/controllers"
	"net/http"
)

var rotasMusicas = []Rota{ //Slice de struct rota para criar todas as rotas que envolvem usuário
	{
		Uri:                "/musicas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarMusica,
		RequerAutenticacao: false,
	}, // Rota para criar usuário
	{
		Uri:                "/musicas/pular",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Pular,
		RequerAutenticacao: true,
	}, //rota para buscar todos os usuários
	
		{
			Uri:                "/musicas/fila",
			Metodo:             http.MethodGet,
			Funcao:             controllers.Fila,
			RequerAutenticacao: false,
		},
	/*
		{
			Uri:                "/usuarios/{usuarioid}",
			Metodo:             http.MethodPut,
			Funcao:             controllers.EditarUsuario,
			RequerAutenticacao: false,
		}, //Rota para editar um usuário
		{
			Uri:                "/usuarios/{usuarioid}",
			Metodo:             http.MethodDelete,
			Funcao:             controllers.DeletarUsuario,
			RequerAutenticacao: false,
		}, //Rota para deletar um usuário */
}
