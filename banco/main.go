package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/banco/clientes"
	"github.com/gabrielhalmenschlager/curso-golang-alura/banco/contas"
)

func main() {
	clienteGabriel := clientes.Titular{
		Nome:      "Gabriel",
		CPF:       "123.111.222.333.12",
		Profissao: "Desenvolvedor",
	}

	contaDoGabriel := contas.ContaCorrente{
		Titular:       clienteGabriel,
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}

	contaDoPedro := contas.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "Pedro",
			CPF:       "321.333.222.111.32",
			Profissao: "Analista de Dados",
		},
		NumeroAgencia: 172,
		NumeroConta:   654321,
	}

	fmt.Println(contaDoGabriel)
	fmt.Println(contaDoPedro)
}
