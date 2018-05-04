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

const monitoramento = 3 // crio constantes para usar em go
const delay = 5

func main() {

	exibeIntrodução()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Log...")
			imprimeLogs()
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
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problema. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
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

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arquivo)

	// go representa tempo de uma forma diferente (https://golang.org/src/time/format.go)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
