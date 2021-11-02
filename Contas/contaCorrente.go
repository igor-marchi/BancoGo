package contas

import (
	"curso-02/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumAg, NumAcc int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.saldo
	} else {
		return "saldo insuficiente", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "não é possível depositar valores menores que 0", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorTrasferencia float64, contaDestino *ContaCorrente) bool {
	if valorTrasferencia < c.saldo && valorTrasferencia > 0 {
		c.saldo -= valorTrasferencia
		contaDestino.Depositar(valorTrasferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObsterSaldo() float64 {
	return c.saldo
}
