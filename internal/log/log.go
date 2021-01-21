package log

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
)

// Level est le niveau de log
type Level int

// Variables pour le niveau de débug
const (
	DEBUG Level = iota
	INFO = iota
	SUCCESS = iota
	WARN = iota
	ERROR = iota
)

// Log est une structure gérant le logging
type Log struct {
	level Level
}

var once sync.Once
var instance *Log

// GetLogger permet de récupérer l'unique instance de logger
func GetLogger() *Log {
	once.Do(func() {
		instance = &Log{}
	})
	return instance
}

// SetLevel permet de définir le niveau de log
func (l *Log) SetLevel(newLevel Level) {
	l.level = newLevel
}

// Debug permet de logger un message de débug
func (l Log) Debug(message string) {
	l.log(DEBUG, message)
}

// Title permet de logger un titre de débug
func (l Log) Title(level Level, message string) {
	titleColor := color.New(color.FgGreen).Add(color.Bold)
	message = "\n===" + titleColor.Sprintf("%s", message) + "===\n"
	l.log(level, message)
}

// Info permet de logger un message d'information
func (l Log) Info(message string) {
	message = fmt.Sprintf("[%s] %s: %s\n", color.BlueString("+"), color.BlueString("INFO"), message)
	l.log(INFO, message)
}

// Success permet de logger un message de succès
func (l Log) Success(message string) {
	message = fmt.Sprintf("[%s] %s\n", color.GreenString("+"), message)
	l.log(SUCCESS, message)
}

// Warn permet de logger un message d'avertissement
func (l Log) Warn(message string) {
	message = fmt.Sprintf("[%s] %s: %s\n", color.YellowString("/"), color.YellowString("WARN"), message)
	l.log(WARN, message)
}

// Error permet de logger un message d'erreur
func (l Log) Error(message string) {
	message = fmt.Sprintf("[%s] %s: %s\n", color.RedString("-"), color.RedString("ERROR"), message)
	l.log(ERROR, message)
}

func (l Log) log(level Level, message string) {
	if level >= l.level {
		fmt.Print(message)
	}
}