package main

import (
	"fmt"
	"net/http"
	"os"
)

// Cores ANSI para o terminal
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println(Yellow + "\n📜 Exibindo Logs..." + Reset)
		case 0:
			fmt.Println(Red + "\nSaindo do Programa..." + Reset)
			os.Exit(0)
		default:
			fmt.Println(Red + "\n❌ Não conheço este comando!" + Reset)
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Gabriel"
	versao := 1.1

	fmt.Println(Cyan + "=====================================" + Reset)
	fmt.Println(Magenta + "      🖥️  Monitor de Sites v1.1" + Reset)
	fmt.Println(Cyan + "=====================================" + Reset)
	fmt.Println("Olá, sr.", Green+nome+Reset)
	fmt.Println("Versão do programa:", Yellow, versao, Reset)
	fmt.Println()
}

func exibeMenu() {
	fmt.Println(Blue + "=======================" + Reset)
	fmt.Println(White + "      MENU PRINCIPAL" + Reset)
	fmt.Println(Blue + "=======================" + Reset)
	fmt.Println(Cyan + "1" + Reset + " - Iniciar Monitoramento")
	fmt.Println(Cyan + "2" + Reset + " - Exibir Logs")
	fmt.Println(Cyan + "0" + Reset + " - Sair do Programa")
	fmt.Println()
}

func lerComando() int {
	var comandoLido int
	fmt.Print(Yellow + "Digite o comando desejado: " + Reset)
	fmt.Scan(&comandoLido)
	fmt.Println(Gray, "O endereço da minha variável comando é", &comandoLido, Reset)
	fmt.Println(Green, "O comando escolhido foi", comandoLido, Reset)
	fmt.Println()
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println(Magenta + "\n🚀 Iniciando monitoramento..." + Reset)
	fmt.Println(Cyan + "---------------------------------" + Reset)

	sites := [4]string{
		"https://www.alura.com.br",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.udemy.com",
	}

	for i, site := range sites {
		fmt.Println(Blue, "\nTestando site", i+1, ":", site, Reset)
		testaSite(site)
	}

	fmt.Println(Cyan + "\n---------------------------------" + Reset)
	fmt.Println(Green + "✅ Monitoramento finalizado!\n" + Reset)
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println(Green, "✔️  Site:", site, "foi carregado com sucesso!", Reset)
	} else {
		fmt.Println(Red, "⚠️  Site:", site, "está com problemas. Status Code:", resp.StatusCode, Reset)
	}
}
