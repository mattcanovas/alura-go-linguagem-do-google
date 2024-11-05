package main

import "fmt"

func main() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("O comando escolhido foi", comando)

	// if comando == 1 {
	// 	fmt.Println("Monitorando....")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs....")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando....")
	case 2:
		fmt.Println("Exibindo logs....")
	case 3:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Não conheço este comando")
	}
}
