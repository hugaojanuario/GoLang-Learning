package main

import (
	"fmt"

)

type ContaBancaria struct{
	titular string
	numeroConta int
	saldo float64
}

func main (){
	contaHugo := ContaBancaria{"Hugo", 123, 500.00}
	fmt.Println(contaHugo)

	contaHugo.saldo = sacar(contaHugo, 100)
	fmt.Println("Saldo Atual:", contaHugo.saldo)

	contaHugo.saldo = depositar(contaHugo,3400)
	fmt.Println("Saldo Atual:", contaHugo.saldo)


}

func depositar (c ContaBancaria, valorDeposito float64) float64{
	valor := valorDeposito > 0 && valorDeposito <= 10000
	if valor{
		fmt.Println(valor)
		c.saldo += valorDeposito
		return c.saldo
	}else{
		fmt.Println("Saldo nao possivel")
		return c.saldo
	}
}

func sacar (c ContaBancaria, valorSaque float64) float64{

	valor := valorSaque <= c.saldo && valorSaque > 0
	if valor{
		fmt.Println(valor)
		c.saldo -= valorSaque
		
		return c.saldo
	}else{
		fmt.Println("Saldo insuficiente e Precisa ser maior que zero")
		return c.saldo
	}

}