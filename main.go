package main

import (
	"fmt"
	"banco/contas"
	// "banco/titular"
)


func main (){
	contaHeitor := contas.ContaCorrente{}
	contaHeitor.Depositar(100)
	
	fmt.Println("Conta Heitor:",contaHeitor)


	
}