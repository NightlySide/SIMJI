package assembler

import (
	"bytes"
	"io"
	"log"
	"os"
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

func TestExportBinToFile(t *testing.T) {
	content, _ := ProgramFileToStringArray("../../testdata/negatif.asm")
	instru := ComputeHexInstructions(StringLinesToInstructions(content))

	err := ExportBinaryToFile(instru, "../../testdata/test_program.bin")
	if err != nil {
		t.Error(err.Error())
	}
	os.Remove("../../testdata/test_program.bin")
}

func TestExportPrgmToFile(t *testing.T) {
	content, _ := ProgramFileToStringArray("../../testdata/negatif.asm")

	err := ExportProgramToFile(content, "../../testdata/test_program.asm")
	if err != nil {
		t.Error(err.Error())
	}
	os.Remove("../../testdata/test_program.asm")
}

func TestPrintProgram(t *testing.T) {
	output := CaptureOutput(func() { PrintProgram([]int{0x00000000}) })
	if output != "0x00000000\t0x00000000\n" {
		t.Errorf("Wrong output - Expected: 0x00000000\t0x00000000, Got: %s\n", output)
	}
}
