package vm

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"
	"time"
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

func TestLoadProgramFromFile(t *testing.T) {
	prog := LoadProgFromFile("../../testdata/test.bin")
	if len(prog) != 15 {
		t.Error("Error parsing binary file")
	}
}

func TestVMParams(t *testing.T) {
	vm := NewVM(32, 1000)
	if vm.GetCycles() != 0 {
		t.Error("Error while init VM")
	}

	if len(vm.GetProg()) != 0 {
		t.Error("Error program not empty on init")
	}

	if vm.GetPC() != 0 {
		t.Error("Program counter badly initialized")
	}

	if len(vm.GetRegs()) != 32 {
		t.Error("Bad number of registers on init")
	}

	if len(vm.GetMemory()) == 0 {
		t.Error("Bad number of memory blocks on init")
	}
}

func TestVMShowRegAndMem(t *testing.T) {
	/*vm := NewVM(2, 2)
	output := CaptureOutput(vm.showRegs)
	if !strings.Contains(output, "regs = 0000 0000") {
		t.Error("Bad printing of regs")
		t.Error(output)
	}

	output = CaptureOutput(vm.showMem)
	if strings.TrimSpace(output) != "memory = 0000 0000" {
		t.Error("Bad printing of memory blocks")
		t.Error(output)
	}*/
}

func TestVMCyclesPerSecCalc(t *testing.T) {
	vm := NewVM(2, 2)
	vm.cycles = 1
	vm.totalTime = time.Second * 1

	if vm.GetCyclesPerSec() != 1 {
		t.Error("Error in calculating number of cycles per second")
	}
}

func TestBinaryComplement(t *testing.T) {
	number := ReverseBinaryComplement(0b1010, 4)
	if number != -2 {
		t.Error("Error while converting negative number")
	}
}
