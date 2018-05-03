package soma

import "github.com/DATA-DOG/godog"

func faoASomaDe(arg1, arg2 int) error {
	return godog.ErrPending
}

func oResultadoDeveSer(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^faço a soma de (\d+) \+ (\d+)$`, faoASomaDe)
	s.Step(`^o resultado deve ser (\d+)$`, oResultadoDeveSer)
}
