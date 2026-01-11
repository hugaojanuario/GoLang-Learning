package contas

import (
	"fmt"
	"banco/titular"
)

type ContaCorrente struct{
	Titular titular.Titular
	NumeroConta int
	saldo float64
}
func (c *ContaCorrente) Transferir (valorTransferencia float64, contaDestino *ContaCorrente) string{
	if valorTransferencia <= c.saldo && valorTransferencia > 0{
		c.saldo -= valorTransferencia
		contaDestino.saldo += valorTransferencia
		return "Transferencia feita com sucesso !"
	}else{
		return "Erro na tranferencia"
	}
}

func (c *ContaCorrente)Depositar (valorDeposito float64) (string, float64){
	valor := valorDeposito > 0 && valorDeposito <= 10000
	if valor{
		fmt.Println(valor)
		c.saldo += valorDeposito
		return "saldo altual",c.saldo
	}else{
		fmt.Println("saldo nao possivel")
		return "saldo atual", c.saldo
	}
}


func Sacar (c ContaCorrente, valorSaque float64) float64{

	valor := valorSaque <= c.saldo && valorSaque > 0
	if valor{
		fmt.Println(valor)
		c.saldo -= valorSaque
		
		return c.saldo
	}else{
		fmt.Println("saldo insuficiente e Precisa ser maior que zero")
		return c.saldo
	}

}