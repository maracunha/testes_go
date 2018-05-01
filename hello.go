package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"bufio"
)

const monitoramento = 3 // crio constantes para usar em go
const delay = 5

func main() {

	exibeIntrodução()
	leSitesDoArquivo()

	for {
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

}

func exibeIntrodução() {

	fmt.Println("")
	nome := "Gullas"
	idade := 24
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
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
	fmt.Println("")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites { // RANGE retorno a posição i(indice) e quem está nessa posição, site0
			fmt.Println("Estou passano na possição", i, ":", site)
			testaSite(site)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado corretamente!")
	} else {
		fmt.Println("Site", site, "está com problema. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	//	arquivo, err := os.Open("sites.txt")

	arquivo, err := ioutil.ReadFile("sites.txt") 

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo)) // tem que colocar o string para converter o array de bytes para string (legivel)

	return sites
}

// err - tratando os erro
// os.Open (para abri um arquivo externo) - abre um ponteiro puro.
// Biblioteca "io/ioutil"
// arquivo, err := ioutil (retorna um array de bytes) que podemos converter em string e abrir o arquivo todo.
// "bufio" biblioteca / 