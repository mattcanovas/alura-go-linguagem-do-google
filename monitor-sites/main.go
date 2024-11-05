package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs....")
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando....")
	var sites []string = []string{"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br"}
	// for i := 0; i < len(sites); i++ {

	// }
	for _, site := range sites {
		resp, _ := http.Get(site)
		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso!")
		} else {
			fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		}
	}
}
