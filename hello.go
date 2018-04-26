package main

// para build:  go build hello.go
// para executar: ./hello
// ou para ir direto: go run hello.go

import "fmt"
import "reflect"

func main() {
	nome := "Douglas"
	idade:= 24

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	
	versao:= 1.1
	fmt.Println("Olá sr.!", nome, "sua idade é", idade)
	fmt.Println	("Este programa está na versão", versao)
	
	fmt.Println("1- Iniiaf Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	var comandos int
	// Sacnf precisa do modificar (%d), varial
	fmt.Scanf ("%d", &comando)
	// Scan é mais inteligente e não precisa do modificador
	fmt.Scan(&comandos)

	fmt.Println("O endereço da variável comando é", &comando)
	fmt.Println("O endereço da variável comando é", &comandos)

	fmt.Println("o comando escolhido foi", comando)
}

//
// git ls-files (para saber qual arquivos o git está controlando)
