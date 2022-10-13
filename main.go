package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var w fyne.Window
var a fyne.App
var filePath string = ""
var fileContent string = ""
var fileName string = ""
var multiLineEntry *widget.Entry
var nonSavedSymbol = "\u2B24"
var titleOfWindow = "Editeur de texte"
var newFileMenuItem *fyne.MenuItem
var openFileMenuItem *fyne.MenuItem
var saveFileMenu *fyne.MenuItem

func main() {
	a = app.NewWithID("Editeur de texte")
	w = a.NewWindow(titleOfWindow)
	w.Resize(fyne.NewSize(720, 480))

	newFileMenuItem = fyne.NewMenuItem("Nouveau fichier", newFileFunction)
	openFileMenuItem = fyne.NewMenuItem("Ouvrir un fichier", func() {
		dialog.NewFileOpen(openFileFunction, w).Show()
	})
	saveFileMenu = fyne.NewMenuItem("Sauvegarder", saveFileFunction)
	saveFileMenu.Disabled = true
	menu := fyne.NewMenu("Fichier", newFileMenuItem, openFileMenuItem, saveFileMenu)

	w.SetMainMenu(fyne.NewMainMenu(menu))

	multiLineEntry = widget.NewMultiLineEntry()
	multiLineEntry.SetPlaceHolder("Le contenu de ton fichier")
	multiLineEntry.Disable()
	// ctrl_S := &desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl}
	// w.Canvas().AddShortcut(ctrl_S, saveFileFunction)
	w.SetContent(multiLineEntry)

	w.ShowAndRun()
}
