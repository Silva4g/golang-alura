package main

import (
	"fmt"
)

func main() {
	nome := "Kauan"
	idade := 20
	versao := 1.1
	fmt.Println("Olá, Sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	var comando int
	fmt.Scanf("%d", comando)
}
