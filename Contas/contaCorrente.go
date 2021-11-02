package contas

type ContaCorrente struct {
	Titular string
	NumAg   int
	NumAcc  int
	Saldo   float64
}

func (c *ContaCorrente) sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.Saldo
	} else {
		return "Saldo insuficiente", c.Saldo
	}
}

func (c *ContaCorrente) depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "não é possível depositar valores menores que 0", c.Saldo
	}
}

func (c *ContaCorrente) transferir(valorTrasferencia float64, contaDestino *ContaCorrente) bool {
	if valorTrasferencia < c.Saldo && valorTrasferencia > 0 {
		c.Saldo -= valorTrasferencia
		contaDestino.depositar(valorTrasferencia)
		return true
	} else {
		return false
	}
}
