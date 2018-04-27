// git ls-files (para saber qual arquivos o git está controlando)

package main

// para build:  go build hello.go
// para executar: ./hello
// ou para ir direto: go run hello.go

import "fmt"
import "reflect"

func main() {
	nome := "Douglas"
	idade := 24

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))

	versao := 1.1
	fmt.Println("Olá sr.!", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniiaf Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int

	// Sacnf precisa do modificar (%d), varial
	//	fmt.Scanf ("%d", &comando)
	// Scan é mais inteligente e não precisa do modificador
	fmt.Scan(&comando)
	fmt.Println("O endereço da variável comando é", &comando)
	fmt.Println("o comando escolhido foi", comando)

	//	if comando == 1 {
	//		fmt.Println("Monitando...")
	//	} else if comando == 2 {
	//		fmt.Println("Exibindo Logs...")
	//	} else if comando == 0 {
	//		fmt.Println("Saindo do Programa")
	//	} else {
	//		fmt.Println("Não conheço este comando")
	//	}

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Log...")
	case 3:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Não conhço este comando")
	}

}
