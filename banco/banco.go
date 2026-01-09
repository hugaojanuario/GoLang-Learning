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

	contaHugo.saldo = sacar(contaHugo, 499)

	fmt.Println("Saldo Atual:", contaHugo.saldo)

}

func sacar (c ContaBancaria, valorSaque float64) float64{

	if valorSaque > c.saldo {
		fmt.Println("Saldo insuficiente")
		return c.saldo
	}

	c.saldo -= valorSaque
	return c.saldo

	

	
}