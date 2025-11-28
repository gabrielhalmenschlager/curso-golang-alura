package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("pizzaria/dados/pizzas.json")
	if err != nil {
		fmt.Println("Error File:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Println("Error Decoding JSON:", err)
	}
}

func SavePizza() {
	file, err := os.Create("pizzaria/dados/pizzas.json")
	if err != nil {
		fmt.Println("Error File:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Pizzas); err != nil {
		fmt.Println("Error Enconding JSON:", err)
	}
}
