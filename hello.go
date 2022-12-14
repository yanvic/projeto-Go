package main

import (
	"fmt"
	"net/http"
)
import "os"

func main() {

	exibirDados()
	for {
		exibirMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("exibindo logs")
		case 3:
			fmt.Println("saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("opção inválida")
			os.Exit(-1)
		}
	}
}

func exibirDados() {
	nome := "yan"
	idade := 20
	cargo := "estagiário"
	empresa := "fix pay"
	fmt.Println("olà. Meu nome é", nome)
	fmt.Println("tenho", idade)
	fmt.Println("sou", cargo)
	fmt.Println("na", empresa)
}
func exibirMenu() {
	fmt.Println("1- iniciar")
	fmt.Println("2- histórico")
	fmt.Println("3- sair")
}
func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("o comando escolhido foi:", comando)
	return comando
}
func iniciarMonitoramento() {
	fmt.Println("monitorando")
	site := "https://fixpay.com.br"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("o site", site, "foi carregado com sucesso.")
	} else {
		fmt.Println("o site", site, "deu erro.", resp.StatusCode)
	}
}
