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

	fmt.Println(contaCorrenteDoIgor)
	fmt.Println(contaPoupancaDoIgor)
	fmt.Println(contaCorrenteDoIgor.ObsterSaldo())
	fmt.Println(contaPoupancaDoIgor.ObsterSaldo())
}
