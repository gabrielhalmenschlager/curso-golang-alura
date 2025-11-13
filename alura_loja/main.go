package main

import (
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/alura_loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
