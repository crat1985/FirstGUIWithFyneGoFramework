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

func printError(err error) {
	dialog.NewError(err, w).Show()
	a.SendNotification(fyne.NewNotification("Error", err.Error()))
}

func setFileInfos(path, content, name, title string, enableEdition, nonsaved bool) {
	filePath = path
	fileContent = content
	fileName = name
	if enableEdition {
		multiLineEntry.Enable()
		multiLineEntry.SetText(content)
	}
	if !enableEdition {
		multiLineEntry.Disable()
	}
	saveFileMenu.Disabled = !enableEdition
	titleOfWindow = title
	if nonsaved {
		w.SetTitle(nonSavedSymbol + " " + titleOfWindow)
	} else {
		w.SetTitle(titleOfWindow)
	}
}

func main() {
	a = app.NewWithID("Editeur de texte")
	w = a.NewWindow(titleOfWindow)
	w.Resize(fyne.NewSize(720, 480))

	newFileMenuItem = fyne.NewMenuItem("Nouveau fichier", newFileFunction)
	openFileMenuItem = fyne.NewMenuItem("Ouvrir un fichier", dialog.NewFileOpen(openFileFunction, w).Show)
	saveFileMenu = fyne.NewMenuItem("Sauvegarder", saveFileFunction)
	saveFileMenu.Disabled = true
	menu := fyne.NewMenu("Fichier", newFileMenuItem, openFileMenuItem, saveFileMenu)

	w.SetMainMenu(fyne.NewMainMenu(menu))
	// w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
	// 	w.SetTitle(nonSavedSymbol + " " + titleOfWindow)
	// })

	multiLineEntry = widget.NewMultiLineEntry()
	multiLineEntry.SetPlaceHolder("Le contenu de ton fichier")
	multiLineEntry.Disable()
	multiLineEntry.OnChanged = func(content string) {
		// if !multiLineEntry.Disabled() {
		// fmt.Println(multiLineEntry.Disabled())
		w.SetTitle(nonSavedSymbol + " " + titleOfWindow)
		// }
	}
	w.SetContent(multiLineEntry)

	w.ShowAndRun()
}
