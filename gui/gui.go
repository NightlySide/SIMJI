package gui

import (
	"io/ioutil"
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func loadHTMLPage(ui lorca.UI, filename string) {
	fileContent, err := ioutil.ReadFile("gui/www/" + filename)
	if err != nil {
		log.Fatal("Cannot open : gui/www/" + filename +" ...")
	}

	loadableContents := "data:text/html," + url.PathEscape(string(fileContent))

	err2 := ui.Load(loadableContents)
	if err2 != nil {
		log.Fatal("Cannot load file content onto page...")
	}
}

// ShowGUI permet de lancer le GUI
func ShowGUI() {
	loadServer()

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