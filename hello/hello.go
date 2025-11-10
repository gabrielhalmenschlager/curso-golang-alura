package main

import "fmt"

func main() {
	nome := "Gabriel"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minhha variável comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitoramento...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do Programa...")
	} else {
		fmt.Println("Não conheco este comando!")
	}
}
