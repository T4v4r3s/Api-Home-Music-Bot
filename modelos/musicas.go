package modelos

// Pacote que guarda structs e métodos para usuários.

import (
	"errors"
	"strings"
)

type Musica struct { // Struct para usuários com sua referência em JSON
	URL  string `json:"url,omitempty"`
	Path string `json:"path,omitempty"`
}

// Prepara o usuário verificando se os campos estão preenchidos e tirando os espaços deles

func (musica *Musica) Preparar(etapa string) error {
	if erro := musica.validar(etapa); erro != nil {
		return erro
	}

	if erro := musica.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

//Verifica se os campos do usuário estão ou não vazios

func (musica *Musica) validar(etapa string) error {
	if musica.URL == "" {
		return errors.New("o campo obrigatório URL está em branco")
	}

	return nil

}

//Formata os espaços em Branco dos campos do usuário

func (musica *Musica) formatar(etapa string) error {
	musica.URL = strings.TrimSpace(musica.URL)
	return nil

}
