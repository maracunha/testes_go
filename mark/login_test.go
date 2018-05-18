// exemplo do YouTube QA Ninja. 
// acessado em 17/05/2018, em: https://www.youtube.com/watch?v=59Nvz6B2Ifc&t=6944s

package ryan

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"

	"github.com/DATA-DOG/godog"
	"github.com/tebeka/selenium"
	"github.com/testes_go/mark/support"
)

var Driver selenium.WebDriver

func queEuAcesseiAPginaPrincipal() error {
	Driver.Get("http://marktasks.herokuapp.com")
	return nil
}

func faoLoginComE(email, senha string) (err error) {

	campoEmail, err := Driver.FindElement(selenium.ByID, "login_email")

	if err != nil {
		return
	}

	campoEmail.SendKeys(email)

	campoSenha, err := Driver.FindElement(selenium.ByID, "login_password")
	if err != nil {
		return
	}

	campoSenha.SendKeys(senha)

	botaoLogin, err := Driver.FindElement(selenium.ByID, "btLogin")
	if err != nil {
		return
	}

	botaoLogin.Click()

	return nil
}

func souAutenticadoComSucesso() (err error) {

	emailMenu, err := Driver.FindElement(selenium.ByClassName, "profile-address")

	if err != nil {
		return
	}

	saida, _ := emailMenu.Text()

	if saida != "eu@papito.io" {
		return fmt.Errorf("Erro ao validar usuário autenticado")
	}

	return nil
}

func deveVerASeguinteMensagem(mensagem string) (err error) {

	divAlerta, err := Driver.FindElement(selenium.ByClassName, "alert-login")

	if err != nil {
		return
	}

	saida, _ := divAlerta.Text()

	if saida != mensagem {
		return fmt.Errorf("Esperado: %v - Obtido: %v", mensagem, saida)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que eu acessei a página principal$`, queEuAcesseiAPginaPrincipal)
	s.Step(`^faço login com "([^"]*)" e "([^"]*)"$`, faoLoginComE)
	s.Step(`^sou autenticado com sucesso$`, souAutenticadoComSucesso)
	s.Step(`^deve ver a seguinte mensagem "([^"]*)"$`, deveVerASeguinteMensagem)

	s.BeforeScenario(func(interface{}) {
		Driver = support.WDinit()
	})

	s.AfterScenario(func(i interface{}, e error) {
		sc := i.(*gherkin.Scenario)
		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "-"))

		shot, err := Driver.Screenshot()

		if err != nil {
			fmt.Println("err:  ", err)
		}

		//fmt.Println("Shot:  ", shot)

		support.SaveImage(shot, fileName)

		Driver.Quit()

	})
}
