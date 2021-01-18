package gui

import (
	"fmt"
	"net/http"
)

// LoadServer lance le server de distrib de fichiers
func LoadServer() {
	fmt.Println("Starting web server on port :5000")
	dir := http.FileServer(http.Dir("./static"))
	go http.ListenAndServe(":5000", dir)
}