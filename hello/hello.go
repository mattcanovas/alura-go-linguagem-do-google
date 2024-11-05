package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Matheus"
	// nome := "Matheus" atribuição implicita e inferindo tipo dela automaticamente apenas com o :
	var versao float32 = 1.1
	var idade = 21
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))
}
