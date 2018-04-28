// _ é um operador de variavel em branco
	//nome, idade := devolveNomeIdade()    (assim eu uso todas as variáveis)
	//fmt.Println("Meu nome é", nome)
	//fmt.Println("Minha idade é", idade)

 	//_, idade := devolveNomeIdade()    // caso não queira usar alguma, é só usar o _ para ficar em branco
	//fmt.Println("Minha idade é", idade)

// para retornar mais de uma varial, devo colcoar o que será retornado entre ()
//func devolveNomeIdade() (string, int) {
//	nome := "Douglas"
//	idade := 24
//	return nome, idade
// for {  (sem passar nenhum parametro, ele fico em loop infinito.. em go não existe while)
//} 



package main

import "fmt"
import "os"

import "net/http"

func main() {


	exibeIntrodução()
	exibeMenu()

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

func exibeIntrodução() {

	nome := "Gullas"
	idade := 24
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniiaf Monitoramento")
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
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
