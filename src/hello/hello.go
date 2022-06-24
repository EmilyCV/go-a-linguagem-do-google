package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			imprimeLogs()
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
	sites := leSitesDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "está com problemas.", resp.Status)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("./src/hello/sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
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
	arquivo, err := os.OpenFile("./src/hello/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - Online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("./src/hello/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
