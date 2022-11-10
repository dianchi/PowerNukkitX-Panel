package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MainMenu(w fyne.Window) *fyne.Container {
	cmdwindow := widget.NewMultiLineEntry()
	cmdwindow.Disable()
	cmdwindow.SetMinRowsVisible(10)
	//src.Runner(textentry, "tree")
	mainmenu := container.NewVBox(cmdwindow)
	return mainmenu

}
