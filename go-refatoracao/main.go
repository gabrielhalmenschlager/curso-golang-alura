package main

import (
	"fmt"
	"log"
	"myapi/internal/config"
	"myapi/internal/handlers"
	"net/http"
)

func main() {
	config.ConnectDatabase()

	// Endpoint raiz
	http.HandleFunc("/api", indexHandler)

	// Endpoints para Itens
	http.HandleFunc("/itens", handlers.ListItensHandler)                // GET para listar todos os itens
	http.HandleFunc("/itens/get", handlers.GetItenHandler)              // GET para buscar um item (espera id via query: ?id=1)
	http.HandleFunc("/itens/get-code", handlers.GetItenByCodigoHandler) // get-code?codigo=TEC001
	http.HandleFunc("/itens/create", handlers.CreateItenHandler)        // POST para criar um item
	http.HandleFunc("/itens/update", handlers.UpdateItenHandler)        // PUT para atualizar um item (JSON com id)
	http.HandleFunc("/itens/delete", handlers.DeleteItenHandler)        // DELETE para deletar um item (espera id via query: ?id=1)

	// Endpoints para Categorias
	http.HandleFunc("/categorias", handlers.ListCategoriasHandler)         // GET para listar todas as categorias
	http.HandleFunc("/categorias/get", handlers.GetCategoriaHandler)       // GET para buscar uma categoria (espera id via query)
	http.HandleFunc("/categorias/create", handlers.CreateCategoriaHandler) // POST para criar uma categoria
	http.HandleFunc("/categorias/update", handlers.UpdateCategoriaHandler) // PUT para atualizar uma categoria (JSON com id)
	http.HandleFunc("/categorias/delete", handlers.DeleteCategoriaHandler) // DELETE para deletar uma categoria (espera id via query)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler raiz
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API Go!")
}
