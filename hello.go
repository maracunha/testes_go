package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	exibeNomes()  // Slice
	exibeCarros() // Array

	//exibeIntrodução()

	for {
		//	exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Log...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntrodução() {

	nome := "Gullas"
	idade := 24
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("o comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado corretamente!")
	} else {
		fmt.Println("Site", site, "está com problema. Status Code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Doulgas", "Mario", "Bernado"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))                                  // para saber o tipo de "nome"
	fmt.Println("O meu slice tem", len(nomes), "itens")                 // LEN para saber o tamnaho do array
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens") // CAP para saber a capacidade do array

	fmt.Println()
	nomes = append(nomes, "Aparecida")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}

func exibeCarros() {
	var carros [4]string
	carros[0] = "chevete"
	carros[1] = "civic"
	carros[2] = "celta"

	fmt.Println()
	fmt.Println(carros)
	fmt.Println(reflect.TypeOf(carros))
	fmt.Println("O meu array tem", len(carros), "itens")
	fmt.Println("O meu array tem capacidade para", cap(carros), "itens")
}
