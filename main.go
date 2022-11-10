package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/dianchi/PowerNukkitX-Panel/ui"
)

//"fmt"
//"fyne.io/fyne/v2/theme"
//"fyne.io/fyne/v2/container"
//"fyne.io/fyne/v2/widget"
//"github.com/dianchi/PowerNukkitX-Panel/src"

//"github.com/fyne-io/terminal"

// go build -ldflags -H=windowsgui
//var channel = make(chan string)
var Out string

func main() {

	a := app.New()
	mainwindows := a.NewWindow("Hello")
	mainmenu := ui.MainMenu(mainwindows)
	downloadmenu := ui.DownloadMenu(mainwindows)
	tabs := container.NewAppTabs(
		container.NewTabItem("downloader", downloadmenu),
		container.NewTabItem("content", mainmenu),
	)

	mainwindows.SetContent(tabs)
	mainwindows.Resize(fyne.NewSize(800, 800))
	mainwindows.SetFixedSize(true)
	mainwindows.ShowAndRun()

}
