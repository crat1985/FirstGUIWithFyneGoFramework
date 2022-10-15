package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func saveFileFunction( /*s fyne.Shortcut*/ ) {
	if filePath == "" {
		dialog.NewFileSave(func(f fyne.URIWriteCloser, err error) {
			// defer f.Close()
			if err != nil {
				printError(err)
				return
			}
			if f == nil {
				return
			}
			_, ferr := os.Open(f.URI().Path())
			if ferr != nil {
				_, ferr = os.Create(f.URI().Path())
				if ferr != nil {
					dialog.NewError(ferr, w).Show()
					return
				}
			}
			// setFileInfos(f.URI().Path(),fileContent,f.URI().Name(),f.URI().Name(),true)
			filePath = f.URI().Path()
			fileName = f.URI().Name()
			f.Close()
			filePathNotEmpty()
		}, w).Show()
	} else {
		filePathNotEmpty()
	}
}

func filePathNotEmpty() {
	ferr := os.WriteFile(filePath, []byte(multiLineEntry.Text), 0755)
	if ferr != nil {
		printError(ferr)
		return
	}
	// dialog.NewInformation("Sauvegardé", "Fichier sauvegardé avec succès !", w).Show()
	a.SendNotification(fyne.NewNotification("Sauvegardé", "Fichier sauvegardé avec succès !"))
	setFileInfos(filePath, multiLineEntry.Text, fileName, fileName, true, false)
}
