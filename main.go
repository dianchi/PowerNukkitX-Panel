package main

import (
	//"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dianchi/PowerNukkitX-Panel/src"
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
	w := a.NewWindow("Hello")
	//_, writer := src.Runner()
	//writer.WriteString("tree")

	textentry := widget.NewMultiLineEntry()
	textentry.Disable()
	src.Runner(textentry)
	//textentry.SetText(Out)
	//fmt.Println("Hello \n" + Out)
	content := container.NewVBox(textentry)
	w.SetContent(content)
	//hello := widget.NewLabel("Hello Fyne!")
	//w.SetContent(container.NewVBox(
	//	hello,
	//	widget.NewButton("Hi!", func() {
	//		hello.SetText("Welcome :)")
	//	}),
	//))

	w.Resize(fyne.NewSize(800, 800))
	w.SetFixedSize(true)
	//src.Runner(w)
	w.ShowAndRun()

}
