package main

import "fmt"

func main() {
	var nome string = "Matheus"
	var versao float32 = 1.1
	var idade int
	// se o valor da variavel int não é declarada ele assumi o valor 0
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}
