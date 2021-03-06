package log

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

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

func TestNewBar(t *testing.T) {
	bar := new(Bar)
	bar.NewOption(0, 100)
	if bar.current != 0 || bar.total != 100 || bar.graph != "█" {
		t.Error("Error while initializing the progress bar")
	}
}

func TestNewBarWithGraph(t *testing.T) {
	bar := new(Bar)
	bar.NewOptionWithGraph(0, 100, "x")

	if bar.graph != "x" {
		t.Error("Error while init with custom graph")
	}
}

func TestBarPercen(t *testing.T) {
	bar := new(Bar)
	bar.NewOption(0, 100)
	if bar.current != 0 || bar.total != 100 || bar.getPercent() != 0 {
		t.Error("Error while getting the bar percentage")
	}

	bar.NewOption(100, 100)
	if bar.getPercent() != 100 {
		t.Error("Error while getting the bar percentage")
	}
}

func TestFinishBar(t *testing.T) {
	bar := new(Bar)
	output := CaptureOutput(bar.Finish)
	if output != "\n" {
		t.Error("Finish function don't retrun empty new line")
	}
}

func TestPlayBar(t *testing.T) {
	bar := new(Bar)
	bar.NewOption(0, 100)
	output := CaptureOutput(func() { bar.Play(0) })
	if strings.TrimSpace(output) != "[                                                  ]  0%        0/100" {
		t.Error("Not getting expected bar output")
	}

	output = CaptureOutput(func() { bar.Play(100) })
	args := strings.Split(strings.TrimSpace(output), " ")
	if args[len(args)-1] != "100/100" {
		t.Error("Not getting expected bar output")
		t.Error(output)
	}
}
