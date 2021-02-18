package gui

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

// LoadServer lance le server de distrib de fichiers
func LoadServer(staticfiles http.FileSystem) {
	log.Info().Msg("Starting web server on port :5000")
	dir := http.FileServer(staticfiles)
	go http.ListenAndServe(":5000", dir)
}
