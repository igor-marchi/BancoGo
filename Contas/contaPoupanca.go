package contas

import "curso-02/clientes"

type ContaPoupanca struct {
	Titular                 clientes.Titular
	NumAg, NumAcc, Operacao int
	saldo                   float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.saldo
	} else {
		return "saldo insuficiente", c.saldo
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "não é possível depositar valores menores que 0", c.saldo
	}
}

func (c *ContaPoupanca) ObsterSaldo() float64 {
	return c.saldo
}
