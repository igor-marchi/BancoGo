package main

import (
	"curso-02/clientes"
	"curso-02/contas"
	"fmt"
)

func main() {
	clienteIgor := clientes.Titular{
		Nome:      "Igor",
		CPF:       "123.123.666.88",
		Profissao: "Developer",
	}

	contaCorrenteDoIgor := contas.ContaCorrente{
		Titular: clienteIgor,
		NumAg:   123,
		NumAcc:  123,
	}

	contaPoupancaDoIgor := contas.ContaPoupanca{
		Titular:  clienteIgor,
		NumAg:    123,
		NumAcc:   123,
		Operacao: 12,
	}

	contaCorrenteDoIgor.Depositar(100)
	contaPoupancaDoIgor.Depositar(100)

	PagarBoleto(&contaCorrenteDoIgor, 60)
	PagarBoleto(&contaPoupancaDoIgor, 60)

	fmt.Println(contaCorrenteDoIgor)
	fmt.Println(contaPoupancaDoIgor)
	fmt.Println(contaCorrenteDoIgor.ObsterSaldo())
	fmt.Println(contaPoupancaDoIgor.ObsterSaldo())
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}
