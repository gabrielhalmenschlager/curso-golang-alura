package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGabriel := ContaCorrente{
		titular:       "Gabriel",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	}

	fmt.Println(contaDoGabriel)

	// contaDaBruna := ContaCorrente{"Bruna", 172, 654321, 150.50}
	// fmt.Println(contaDaBruna)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"
	// contaDaCris.saldo = 500

	// fmt.Println(*contaDaCris)
}
