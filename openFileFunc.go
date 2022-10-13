package main

import (
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func openFileFunction(f fyne.URIReadCloser, err error) {
	if err != nil {
		dialog.NewError(err, w).Show()
		return
	}
	if f == nil {
		dialog.NewError(errors.New("erreur lors de l'ouverture du fichier"), w).Show()
		return
	}
	fpath := f.URI().Path()
	fmt.Printf("[LOG] En train d'ouvrir %s...\n", fpath)

	fileContentTemp, rerr := os.ReadFile(fpath)
	defer f.Close()
	if rerr != nil {
		dialog.NewError(rerr, w).Show()
	}
	filePath = f.URI().Path()
	fileName = f.URI().Name()
	fileContent = string(fileContentTemp)
	multiLineEntry.SetText(fileContent)
	multiLineEntry.Enable()
	titleOfWindow = f.URI().Name()
	w.SetTitle(nonSavedSymbol + " " + titleOfWindow)
	saveFileMenu.Disabled = false
	fmt.Printf("[LOG] Fichier %s ouvert avec succès...\n", fpath)
}
