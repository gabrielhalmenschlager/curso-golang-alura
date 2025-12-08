package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func CategoriaRoutes(r *mux.Router) {
	r.HandleFunc("/categorias", handlers.ListCategoriasHandler).Methods("GET")
	r.HandleFunc("/categorias/get", handlers.GetCategoriaHandler).Methods("GET")
	r.HandleFunc("/categorias/create", handlers.CreateCategoriaHandler)
	r.HandleFunc("/categorias/update", handlers.UpdateCategoriaHandler)
	r.HandleFunc("/categorias/delete", handlers.DeleteCategoriaHandler)
}
