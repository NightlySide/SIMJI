package log

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	log1 := GetLogger()
	log2 := GetLogger()

	if log1 != log2 {
		t.Error("Singleton not working, multiple instances")
	}
}

func captureOutput(f func()) string {
	// creating a new pipeline
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// don't forget to give back the pipe to the os
	stdout := os.Stdout
	stderr := os.Stderr
	defer func () {
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

func TestLogLevel(t *testing.T) {
	log := GetLogger()

	log.SetLevel(DEBUG)

	if log.level != DEBUG {
		t.Error("Cannot set logging level")
	}

	log.SetLevel(ERROR)
	output := captureOutput(func() {log.Info("should'nt be printed")})
	if output != "" {
		t.Error("sqfqfqf")
	}
}

func TestPrinting(t *testing.T) {
	log := GetLogger()
	GetLogger().SetLevel(DEBUG)
	phrase := "Hello world!"
	output := captureOutput(func() {log.Debug(phrase)})

	if output != phrase {
		t.Errorf("Log debug error - Expected: %s, got: %s\n", phrase, output)
	}
}