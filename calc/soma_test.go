package soma

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var Resultado int

func faoASomaDe(v1, v2 int) error {
	Resultado = v1 + v2
	return nil
}

func oResultadoDeveSer(soma int) error {

	if Resultado != soma {
		return fmt.Errorf("Erro ao realizar a soma. Experado: %v Obtido: %v", soma, Resultado) //%v para inserir valores. Declarar a varrável em seguida no linha)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^faço a soma de (\d+) \+ (\d+)$`, faoASomaDe)
	s.Step(`^o resultado deve ser (\d+)$`, oResultadoDeveSer)
}
