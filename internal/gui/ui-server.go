package gui

import (
	"net/http"
	"simji/internal/log"
)

// LoadServer lance le server de distrib de fichiers
func LoadServer(staticfiles http.FileSystem) {
	log.GetLogger().Info("Starting web server on port :5000")
	dir := http.FileServer(staticfiles)
	go http.ListenAndServe(":5000", dir)
}