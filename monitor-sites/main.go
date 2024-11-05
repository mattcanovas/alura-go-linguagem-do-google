package main

import "fmt"

func main() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
	var comando int
	fmt.Scan("%d", &comando)

	fmt.Println("O comando escolhido foi", comando)
}
