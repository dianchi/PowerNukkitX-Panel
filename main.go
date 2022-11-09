package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	//"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
	//"github.com/dianchi/PowerNukkitX-Panel/src"
	//"github.com/dianchi/PowerNukkitX-Panel/src"
	"github.com/fyne-io/terminal"
)

// go build -ldflags -H=windowsgui
func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	t := terminal.New()
	w.SetContent(t)
	//stdout, stdin := src.Runner()
	go func() {
		_ = t.RunLocalShell()
		a.Quit()
	}()
	//hello := widget.NewLabel("Hello Fyne!")
	//w.SetContent(container.NewVBox(
	//	hello,
	//	widget.NewButton("Hi!", func() {
	//		hello.SetText("Welcome :)")
	//	}),
	//))

	w.Resize(fyne.NewSize(400, 200))
	w.FixedSize()
	//src.Runner(w)
	w.ShowAndRun()

}
