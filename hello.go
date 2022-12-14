package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)
import "os"

const monitoramento = 5
const delay = 5

func main() {

	exibirDados()
	registraLog("site-falso", false)
	for {
		exibirMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimirLogs()
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
	fmt.Println("")
	return comando
}
func iniciarMonitoramento() {
	fmt.Println("monitorando")

	//sites := []string{"https://fixpay.com.br", "https://credit2b.com.br", "https://youtube.com.br"}
	sites := lerArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("o site:", site, "esta na posicao", i)
			testarSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testarSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("o site", site, "foi carregado com sucesso.")
		registraLog(site, true)
	} else {
		fmt.Println("o site", site, "deu erro.", resp.StatusCode)
		registraLog(site, false)
	}
}
func lerArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.text")

	if err != nil {
		fmt.Println("ocorreu erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}
func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.text", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("ocorreu erro:", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15/04/05 ") + site + "- online" + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}
func imprimirLogs() {
	arquivo, err := ioutil.ReadFile("log.text")

	if err != nil {
		fmt.Println("ocorreu erro:", err)
	}
	fmt.Println(string(arquivo))
}
