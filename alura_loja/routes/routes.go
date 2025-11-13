package routes

import (
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/alura_loja/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)

}
