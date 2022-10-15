package main

import (
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
)

func openFileFunction(f fyne.URIReadCloser, err error) {
	if err != nil {
		printError(err)
		return
	}
	if f == nil {
		printError(errors.New("erreur lors de l'ouverture du fichier"))
		return
	}
	fpath := f.URI().Path()
	fmt.Printf("[LOG] En train d'ouvrir %s...\n", fpath)

	fileContentTemp, rerr := os.ReadFile(fpath)
	defer f.Close()
	if rerr != nil {
		printError(rerr)
	}
	setFileInfos(f.URI().Path(), string(fileContentTemp), f.URI().Name(), f.URI().Name(), true, false)
	fmt.Printf("[LOG] Fichier %s ouvert avec succès...\n", fpath)
}
