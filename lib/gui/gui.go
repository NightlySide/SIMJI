package gui

import (
	"log"

	"github.com/zserge/lorca"
)

// ShowGUI permet de lancer le GUI
func ShowGUI() {
	//loadServer()

	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("http://127.0.0.1:5000/", "", 480, 320)
	if err != nil {
		print("1")
		log.Fatal(err)
	}
	
	bm := newBindingManager(ui)
	bm.setupBindings()

	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}