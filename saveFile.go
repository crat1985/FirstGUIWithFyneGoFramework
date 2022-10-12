package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func saveFileFunction(s fyne.Shortcut) {
	if filePath == "" {
		dialog.NewFileSave(func(f fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			_, ferr := os.Open(f.URI().Path())
			if ferr != nil {
				dialog.NewError(ferr, w).Show()
			}
			filePath = f.URI().Path()
			f.Close()
		}, w).Show()
	}
	if filePath == "" {
		dialog.NewError(fmt.Errorf("erreur 102 : la variable de chemin d'aacès est vide"), w).Show()
		return
	}
	ferr := os.WriteFile(filePath, []byte(multiLineEntry.Text), 0755)
	if ferr != nil {
		dialog.NewError(ferr, w).Show()
		return
	}
	dialog.NewInformation("Sauvegardé", "Fichier sauvegardé avec succès !", w).Show()
	w.SetTitle(titleOfWindow)
}
