package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/dianchi/PowerNukkitX-Panel/ui"
	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
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

func init() {
	fontPath, err := findfont.Find("msyh.ttf")
	if err != nil {
		panic(err)
	}

	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}
func main() {
	a := app.New()
	mainwindows := a.NewWindow("PowerNukkitX面板")
	mainmenu := ui.MainMenu(mainwindows)
	downloadmenu := ui.DownloadMenu(mainwindows)
	tabs := container.NewAppTabs(
		container.NewTabItem("主界面", mainmenu),
		container.NewTabItem("下载", downloadmenu),
	)

	mainwindows.SetContent(tabs)
	mainwindows.Resize(fyne.NewSize(800, 800))
	mainwindows.SetFixedSize(true)
	mainwindows.ShowAndRun()

}
