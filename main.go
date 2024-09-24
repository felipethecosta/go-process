package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
)

//go:embed golang.png
var iconData []byte

func main() {

	if iconData == nil {
		log.Fatal("Erro: Não foi possível carregar o ícone")
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	if len(iconData) == 0 {
		log.Fatal("Erro: iconData está vazio")
	}
	systray.SetIcon(iconData)
	systray.SetTitle("Meu programa Go")
	systray.SetTooltip("Exemplo de app Go com icone na bandeja do sistema")
	mQuit := systray.AddMenuItem("Sair", "Sair do programa")

	// https url
	// mOpenWeb := systray.AddMenuItem("Abrir Site", "Abre o site oficial")
	// go func() {
	// 	for {
	// 		<-mOpenWeb.ClickedCh
	// 		// Em cima do Println podemos deixar o comando para abrir o navegador com a url desejada.
	// 		fmt.Println("Abrindo site...")
	// 	}
	// 	}()

	// refresh do sistema
	mRefresh := systray.AddMenuItem("Reiniciar Sistema", "Reinicia o sistema")
	go func() {
		for {
			<-mRefresh.ClickedCh
			fmt.Println("Reiniciando sistema...")
			cmd := exec.Command(os.Args[0])
			cmd.Start()
			os.Exit(0)
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
