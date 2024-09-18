package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
)


func main() {
	fmt.Println("Hello, World!")
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("/home/felipe/dev/go-process/golang.png"))
	systray.SetTitle("Meu programa Go")
	systray.SetTooltip("Exemplo de app Go com icone na bandeja do sistema")
	mQuit := systray.AddMenuItem("Sair", "Sair do programa")

    // https url
    mOpenWeb := systray.AddMenuItem("Abrir Site", "Abre o site oficial")
    go func() {
        for {
            <-mOpenWeb.ClickedCh
            // Em cima do Println podemos deixar o comando para abrir o navegador com a url desejada.
            fmt.Println("Abrindo site...")
        }
    }()

    // notificação
    mNotify := systray.AddMenuItem("Notificar", "Mostra uma notificação")
    go func() {
        for {
            <-mNotify.ClickedCh
            // Deixa aqui para caso queira implementar uma notificação com a biblioteca do sistema operacional.
            fmt.Println("Mostrando notificação...")
        }
    }()

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

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Println("Erro ao abrir o navegador:", err)
	}
}
