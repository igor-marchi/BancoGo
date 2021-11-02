package main

import (
	"fmt"
)

func main() {
	contaIgor := ContaCorrente{
		Titular: "Igor",
		NumAg:   589,
		NumAcc:  123456,
		Saldo:   100.00,
	}

	contaSilvia := ContaCorrente{
		Titular: "Silvia",
		NumAg:   589,
		NumAcc:  123456,
		Saldo:   500.00,
	}

	fmt.Println(contaIgor.Saldo)
	status := contaSilvia.transferir(300, &contaIgor)
	fmt.Println(status)
	fmt.Println(contaIgor, contaSilvia)
}
