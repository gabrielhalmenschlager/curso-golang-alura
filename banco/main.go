package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaCorrente1 := ContaCorrente{
		titular:       "Gabriel",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	}

	fmt.Println(contaCorrente1)
}
