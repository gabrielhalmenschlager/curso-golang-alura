package main

import (
	"fmt"
	"log"
	"myapi/internal/config"
	"myapi/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()

	r := mux.NewRouter()

	// Endpoint raiz
	r.HandleFunc("/api", indexHandler)

	// Endpoints para Itens
	r.HandleFunc("/itens", handlers.ListItensHandler)                // GET para listar todos os itens
	r.HandleFunc("/itens/get", handlers.GetItenHandler)              // GET para buscar um item (espera id via query: ?id=1)
	r.HandleFunc("/itens/get-code", handlers.GetItenByCodigoHandler) // get-code?codigo=TEC001
	r.HandleFunc("/itens/create", handlers.CreateItenHandler)        // POST para criar um item
	r.HandleFunc("/itens/update", handlers.UpdateItenHandler)        // PUT para atualizar um item (JSON com id)
	r.HandleFunc("/itens/delete", handlers.DeleteItenHandler)        // DELETE para deletar um item (espera id via query: ?id=1)

	// Endpoints para Categorias
	r.HandleFunc("/categorias", handlers.ListCategoriasHandler)         // GET para listar todas as categorias
	r.HandleFunc("/categorias/get", handlers.GetCategoriaHandler)       // GET para buscar uma categoria (espera id via query)
	r.HandleFunc("/categorias/create", handlers.CreateCategoriaHandler) // POST para criar uma categoria
	r.HandleFunc("/categorias/update", handlers.UpdateCategoriaHandler) // PUT para atualizar uma categoria (JSON com id)
	r.HandleFunc("/categorias/delete", handlers.DeleteCategoriaHandler) // DELETE para deletar uma categoria (espera id via query)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Handler raiz
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API Go!")
}
