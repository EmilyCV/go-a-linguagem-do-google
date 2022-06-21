package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const MONITORAMENTOS = 2
const DELAY = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			// Saiu com sucesso
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando!")
			// Saiu com algum problema
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Emily"
	versao := 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Esse programa está na versão", versao)
	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exitir os Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"http://www.alura.com/",
		"http://www.caelum.com/",
	}
	for i := 0; i < MONITORAMENTOS; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
			fmt.Println("")
		}
		time.Sleep(DELAY * time.Second)
	}

}
func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "está com problemas.", resp.Status)
	}
}
