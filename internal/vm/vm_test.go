package vm

import (
	"simji/internal/log"
	"strings"
	"testing"
	"time"
)

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

	if len(vm.GetMemory()) != 1000 {
		t.Error("Bad number of memory blocks on init")
	}
}

func TestVMShowRegAndMem(t *testing.T) {
	vm := NewVM(2, 2)
	output := log.CaptureOutput(vm.showRegs)
	if strings.TrimSpace(output) != "regs = 0000 0000" {
		t.Error("Bad printing of regs")
		t.Error(output)
	}

	output = log.CaptureOutput(vm.showMem)
	if strings.TrimSpace(output) != "memory = 0000 0000" {
		t.Error("Bad printing of memory blocks")
		t.Error(output)
	}
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
