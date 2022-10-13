package main

func newFileFunction() {
	//TODO Confirm create new file
	filePath = ""
	fileContent = ""
	fileName = ""
	multiLineEntry.Enable()
	multiLineEntry.SetText("")
	saveFileMenu.Disabled = false
	titleOfWindow = "Nouveau fichier"
	w.SetTitle(nonSavedSymbol + " " + titleOfWindow)
}
