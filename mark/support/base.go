// exemplo do YouTube QA Ninja. 
// acessado em 17/05/2018, em: https://www.youtube.com/watch?v=59Nvz6B2Ifc&t=6944s

package support

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// WDinit retorna uma instancia do WebDriver
func WDinit() selenium.WebDriver {

	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, "")

	if err != nil {
		fmt.Println("Erro ao instanciar o driver", err.Error())
	}

	driver.SetImplicitWaitTimeout(time.Second * 10)

	//verificar a resolução minima do aplicativo
	driver.ResizeWindow("note", 1280, 800)

	return driver
}

// SaveImage pega o print do webdriver em bytes, converte epara png e salva no projeto
func SaveImage(foto []byte, name string) {

	img, _, _ := image.Decode(bytes.NewReader(foto))

	out, err := os.Create("./log/screenshots/" + name + ".png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, img)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
