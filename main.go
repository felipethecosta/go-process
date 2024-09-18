package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
)


func main() {
	fmt.Println("Hello, World!")
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("/home/administrador/dev/exe-go/golang.png"))
	systray.SetTitle("Meu programa Go")
	systray.SetTooltip("Exemplo de app Go com icone na bandeja do sistema")
	mQuit := systray.AddMenuItem("Sair", "Sair do programa")


	go func() {
    <-mQuit.ClickedCh
    fmt.Print("Saindo...")
    systray.Quit()
    os.Exit(0)
	}()
}

func onExit() {
	// Limpar
}

func getIcon(path string) []byte {
	// carregar icone
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Não foi possivel carregar o icone:",err)
		return nil
	}
	return b
}
