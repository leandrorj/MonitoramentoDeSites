package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			fmt.Println("Exibindo logs...")
		case 2:
			iniciarMonitoramento()
		case 0:
			fmt.Println("Você saiu do programa. :)")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("0 - Sair do Programa")
	fmt.Println("1 - iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
}

func exibeIntroducao() {
	nome := "Douglas"
	idade := 24
	var versao float32 = 1.1
	fmt.Println("Olá, Sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.google.com.br/"
	sites[2] = "https://www.avanade.com.br/"
	sites[3] = "https://www.teste.com.br/"
	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"

	resp, _ := http.Get(site)
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Erro ao tentar carregar a página:", site, "StatusCode:", resp.StatusCode)
	}
}
