package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/fatih/color"
)

// Level est le niveau de log
type Level int

// Variables pour le niveau de débug
const (
	DEBUG   Level = iota
	INFO          = iota
	SUCCESS       = iota
	WARN          = iota
	ERROR         = iota
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
func (l Log) Debug(message string, arg ...interface{}) {
	l.log(DEBUG, message, arg...)
}

// Title permet de logger un titre de débug
func (l Log) Title(level Level, message string, arg ...interface{}) {
	titleColor := color.New(color.FgGreen).Add(color.Bold)
	message = "\n===" + titleColor.Sprint(message) + "===\n"
	l.log(level, message, arg...)
}

// Info permet de logger un message d'information
func (l Log) Info(message string, arg ...interface{}) {
	message = fmt.Sprintf("[%s] %s: %s\n", color.BlueString("+"), color.BlueString("INFO"), message)
	l.log(INFO, message, arg...)
}

// Success permet de logger un message de succès
func (l Log) Success(message string, arg ...interface{}) {
	message = "[" + color.GreenString("+") + "] " + message
	l.log(SUCCESS, message, arg...)
}

// Warn permet de logger un message d'avertissement
func (l Log) Warn(message string, arg ...interface{}) {
	message = "[" + color.YellowString("/") + "] " + color.YellowString("WARN") + ": " + message
	l.log(WARN, message, arg...)
}

// Error permet de logger un message d'erreur
func (l Log) Error(message string, arg ...interface{}) {
	message = "[" + color.RedString("-") + "] " + color.RedString("ERROR") + ": " + message
	l.log(ERROR, message, arg...)
}

func (l Log) log(level Level, message string, arg ...interface{}) {
	if level >= l.level {
		fmt.Printf(message, arg...)
	}
}

// CaptureOutput permet de capturer la sortie d'une fonction
// par exemple un print
func CaptureOutput(f func()) string {
	// creating a new pipeline
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// don't forget to give back the pipe to the os
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	// new channel to get the printing logs
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	// executing the function to get the output
	f()
	writer.Close()
	return <-out
}
