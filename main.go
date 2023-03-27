package main

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := router.Gerar()

	porta := 8888

	fmt.Printf("Escutando na porta %d...", porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", porta), r))
}
