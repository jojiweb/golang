package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Exibindo logs...")

	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)

	default:
		fmt.Println("Esse comando não existe.")
		os.Exit(-1)

	}
}

func exibeIntroducao() {
	var nome = "Alexandre"
	fmt.Println("Olá Sr.", nome, ", seja bem-vindo.")
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
