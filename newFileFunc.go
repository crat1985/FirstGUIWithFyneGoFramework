package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newFileFunction() {
	// dialog.NewCustomConfirm("Enregistrer les modifications ?", "Voulez-vous enregistrer les modifications avant de quitter (les données non-sauvegardées seront perdues) ?", "Annuler", fyne.NewContainer(), doChanges, w).Show()
	if filePath != "" || multiLineEntry.Text != "" {
		createWindow()
	} else {
		setFileInfos("", "", "", "Nouveau fichier", true, true)
	}
}

var confirmWindow fyne.Window
var isOpenedDialog = false

func createWindow() {
	if isOpenedDialog {
		return
	}
	confirmWindow = a.NewWindow("Enregistrer les modifications ?")
	isOpenedDialog = true
	// defer confirmWindow.Close()
	defer func() { isOpenedDialog = false }()
	confirmWindow.CenterOnScreen()
	confirmWindow.SetContent(container.NewVBox(widget.NewLabel("Voulez-vous vraiment quitter l'application (les données non-sauvegardées seront supprimées) ?"), container.NewAdaptiveGrid(3, widget.NewButton("Annuler", cancel), widget.NewButton("Non", noDontSave), widget.NewButton("Oui", yesSave))))
	confirmWindow.Show()
}

func yesSave() {
	confirmWindow.Close()
	saveFileFunction()
	setFileInfos("", "", "", "Nouveau fichier", true, true)
}

func noDontSave() {
	confirmWindow.Close()
	setFileInfos("", "", "", "Nouveau fichier", true, true)
}

func cancel() {
	confirmWindow.Close()
}
