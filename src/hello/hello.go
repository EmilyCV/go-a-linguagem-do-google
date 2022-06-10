package main

import (
	"fmt"
)

func main() {
	nome := "Emily"
	versao := 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Esse programa está na versão", versao)
	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exitir os Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println(comando)
}
