package main

import (
	//"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
	//"github.com/dianchi/PowerNukkitX-Panel/src"
	//"github.com/fyne-io/terminal"
)

// go build -ldflags -H=windowsgui
//var channel = make(chan string)
var Out string

func main() {
	a := app.New()
	mainwindows := a.NewWindow("Hello")
	//_, writer := src.Runner()
	//writer.WriteString("tree")

	downloadtext := widget.NewEntry()
	downloader := container.NewVBox(downloadtext)

	textentry := widget.NewMultiLineEntry()
	textentry.Disable()

	//src.Runner(textentry, "tree")
	content := container.NewVBox(textentry)

	tabs := container.NewAppTabs(
		container.NewTabItem("downloader", downloader),
		container.NewTabItem("content", content),
	)

	mainwindows.SetContent(tabs)

	textentry.SetMinRowsVisible(10)

	mainwindows.Resize(fyne.NewSize(800, 800))
	mainwindows.SetFixedSize(true)
	//src.Runner(w)
	mainwindows.ShowAndRun()

}
