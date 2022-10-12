package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var w fyne.Window
var filePath string = ""
var fileContent string
var multiLineEntry *widget.Entry
var nonSavedSymbol = "\u2B24"
var titleOfWindow = "Hello World"

func main() {
	a := app.New()
	w = a.NewWindow(titleOfWindow)
	w.Resize(fyne.NewSize(720, 480))

	newFileMenuItem := fyne.NewMenuItem("Ouvrir un fichier", func() {
		dialog.NewFileOpen(openFileFunction, w).Show()
	})
	menu := fyne.NewMenu("Fichier", newFileMenuItem)

	w.SetMainMenu(fyne.NewMainMenu(menu))

	multiLineEntry = widget.NewMultiLineEntry()
	multiLineEntry.SetPlaceHolder("Le contenu de ton fichier")
	multiLineEntry.Disable()
	ctrl_S := &desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(ctrl_S, saveFileFunction)
	w.SetContent(multiLineEntry)

	w.ShowAndRun()
}
