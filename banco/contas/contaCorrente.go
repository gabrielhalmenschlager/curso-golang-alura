package contas

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor do deposito menor que zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTrasferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTrasferencia < c.Saldo && valorDaTrasferencia > 0 {
		c.Saldo -= valorDaTrasferencia
		contaDestino.Depositar(valorDaTrasferencia)
		return true
	} else {
		return false
	}
}
