package gui

import (
	"net/http"
)

func loadServer() {
	// Start listener
	http.Handle("/", http.FileServer(http.Dir("gui/www")))
	go http.ListenAndServe(":5000", nil)
}