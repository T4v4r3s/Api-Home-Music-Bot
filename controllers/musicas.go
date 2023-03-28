package controllers

import (
	"api/modelos"
	"api/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

var queue []string
var cmd *exec.Cmd

func CriarMusica(w http.ResponseWriter, r *http.Request) { //Criar um usuário
	corpoRequest, erro := io.ReadAll(r.Body) //Faz a leitura do corpo da requisição
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var musica modelos.Musica //Cria uma variável do tipo modelos.usuario para poder receber as informações do corpo da requisição

	if erro = json.Unmarshal(corpoRequest, &musica); erro != nil { // Converte o corpo da requisição de JSON e joga dentro da struct usuario
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = musica.Preparar("cadastro"); erro != nil { //Realiza os tratamentos e verificação dos dados para serem colocados na struct
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	fmt.Println(musica.URL)
	//cmd := exec.Command("/home/tavares/Área de Trabalho/Api-Home-Music-Bot/baixar.sh", musica.URL)

	respostas.JSON(w, http.StatusCreated, musica)

	go fila(musica.URL)
}

func fila(URL string) {

	if len(queue) < 1 {
		fmt.Print("Fila atual ")
		queue = append(queue, URL)
		fmt.Print(queue)
		Reproduzir(URL)
	} else {
		queue = append(queue, URL)
		fmt.Print("Fila atual ")
		fmt.Print(queue)
	}

}

func Reproduzir(URL string) {

	for len(queue) != 0 {
		cmd = exec.Command("yt-dlp", queue[0], "--exec", "ffplay -nodisp -autoexit", "filename", "--exec", "rm", "filename")
		// Execute o comando e capture a saída
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))

		queue = queue[1:]

		fmt.Println(queue)
		fmt.Println("acabou a reprodução")
	}
}

func Pular(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("pkill", "ffplay")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error stopping music: %v", err)
	}
}
