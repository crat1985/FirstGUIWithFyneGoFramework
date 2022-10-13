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
				dialog.NewError(err, w).Show()
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
			filePath = f.URI().Path()
			fileName = f.URI().Name()
			f.Close()
			filePathNotEmpty()
		}, w).Show()
	} else {
		filePathNotEmpty()
	}
	// if filePath == "" {
	// 	dialog.NewError(fmt.Errorf("erreur 102 : la variable de chemin d'aacès est vide"), w).Show()
	// 	return
	// }

}

func filePathNotEmpty() {
	ferr := os.WriteFile(filePath, []byte(multiLineEntry.Text), 0755)
	if ferr != nil {
		dialog.NewError(ferr, w).Show()
		return
	}
	// dialog.NewInformation("Sauvegardé", "Fichier sauvegardé avec succès !", w).Show()
	a.SendNotification(fyne.NewNotification("Sauvegardé", "Fichier sauvegardé avec succès !"))
	w.SetTitle(titleOfWindow)
	multiLineEntry.Enable()
	titleOfWindow = fileName
	w.SetTitle(titleOfWindow)
	saveFileMenu.Disabled = false
	// fmt.Printf("[LOG] Fichier %s ouvert avec succès...\n", filePath)
}
