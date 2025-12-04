package main

import (
	"log"
	"myapi/internal/config"
	"myapi/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/api/itens", handlers.ListItensHandler).Methods("GET")
	r.HandleFunc("/api/itens/{id}", handlers.GetItenHandler).Methods("GET")
	r.HandleFunc("/api/itens/codigo/{codigo}", handlers.GetItenByCodigoHandler).Methods("GET")
	r.HandleFunc("/api/itens", handlers.CreateItenHandler).Methods("POST")
	r.HandleFunc("/api/itens", handlers.UpdateItenHandler).Methods("PUT")
	r.HandleFunc("/api/itens", handlers.DeleteItenHandler).Methods("DELETE")

	// Endpoints para Categorias
	r.HandleFunc("/categorias", handlers.ListCategoriasHandler).Methods("GET")
	r.HandleFunc("/categorias/get", handlers.GetCategoriaHandler).Methods("GET")
	r.HandleFunc("/categorias/create", handlers.CreateCategoriaHandler)
	r.HandleFunc("/categorias/update", handlers.UpdateCategoriaHandler)
	r.HandleFunc("/categorias/delete", handlers.DeleteCategoriaHandler)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
