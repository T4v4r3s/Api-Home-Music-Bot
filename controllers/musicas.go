package controllers

import (
	"api/modelos"
	"api/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
		cmd = exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "restrict-filenames", "-o", "%(title)s.%(ext)s", "ytsearch:" + queue[0])
		_, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		cmd = exec.Command("yt-dlp", "--get-filename", "--restrict-filenames","-o","%(title)s","ytsearch:" + queue[0])
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		// Obtém o nome do arquivo baixado
		filename := string(out)
		filename = strings.TrimSpace(filename)
		filename = filepath.Join(os.Getenv("PWD"), filename)
		filename = filename + ".mp3"

		// Executa o ffplay para reproduzir o vídeo
		cmd = exec.Command("ffplay", "-nodisp", "-autoexit", filename)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		// Deleta o arquivo baixado
		err = os.Remove(filename)
		if err != nil {
			fmt.Println(err)
		}

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
