package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/itens", handlers.ListItens).Methods("GET")
	r.HandleFunc("/api/itens/{id}", handlers.GetItem).Methods("GET")
	r.HandleFunc("/api/itens/codigo/{codigo}", handlers.GetItenByCode).Methods("GET")
	r.HandleFunc("/api/itens", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/api/itens/", handlers.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/itens/{id}", handlers.DeleteItem).Methods("DELETE")

	// Endpoints para Categorias
	r.HandleFunc("/categorias", handlers.ListCategoriasHandler).Methods("GET")
	r.HandleFunc("/categorias/get", handlers.GetCategoriaHandler).Methods("GET")
	r.HandleFunc("/categorias/create", handlers.CreateCategoriaHandler)
	r.HandleFunc("/categorias/update", handlers.UpdateCategoriaHandler)
	r.HandleFunc("/categorias/delete", handlers.DeleteCategoriaHandler)

	return r
}
