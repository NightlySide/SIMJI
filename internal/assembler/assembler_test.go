package assembler

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestLoadFile(t *testing.T) {
	filename := "../../testdata/allinstructions.asm"
	lines, err := ProgramFileToStringArray(filename)

	if err != nil {
		t.Error(err.Error())
	}

	file, _ := ioutil.ReadFile(filename)
	if len(lines) != len(strings.Split(string(file), "\n")) {
		t.Error("Number of lines don't match")
	}
}

func TestLoadingInstructions(t *testing.T) {
	filename := "../../testdata/allinstructions.asm"
	lines, _ := ProgramFileToStringArray(filename)
	instructions := StringLinesToInstructions(lines)

	assertOpCode := func (instruction []int, opCode int) {
		if (instruction[0] != opCode) {
			t.Errorf("OPCode mismatch - Expected: %d, Got: %d\n", opCode, instruction[0])
		}
	}

	// checking instructions op codes
	assertOpCode(instructions[0], 1)
	assertOpCode(instructions[2], 2)
	assertOpCode(instructions[4], 3)
	assertOpCode(instructions[6], 4)
	assertOpCode(instructions[8], 5)
	assertOpCode(instructions[10], 6)
	assertOpCode(instructions[12], 7)
}

func TestHexInstructions(t *testing.T) {
	filename := "../../testdata/allinstructions.asm"
	lines, _ := ProgramFileToStringArray(filename)
	instructions := StringLinesToInstructions(lines)
	hexInstr := ComputeHexInstructions(instructions)

	assertHexCode := func (instruction int, hex int) {
		if (instruction != hex) {
			t.Errorf("OPCode mismatch - Expected: 0x%x, Got: 0x%x\n", hex, instruction)
		}
	}

	// checking for a good translation
	// add
	assertHexCode(hexInstr[0], 0x08400002)
	assertHexCode(hexInstr[1], 0x086000a1)
	// sub
	assertHexCode(hexInstr[2], 0x10400002)
	assertHexCode(hexInstr[3], 0x10600062)
	// mult
	assertHexCode(hexInstr[4], 0x18400022)
	assertHexCode(hexInstr[5], 0x186000a2)
	// div
	assertHexCode(hexInstr[6], 0x20400022)
	assertHexCode(hexInstr[7], 0x206000a2)
	// and
	assertHexCode(hexInstr[8], 0x28600021)
	assertHexCode(hexInstr[9], 0x28400042)
	// or
	assertHexCode(hexInstr[10], 0x30600022)
	assertHexCode(hexInstr[11], 0x30400042)
	// xor
	assertHexCode(hexInstr[12], 0x38600022)
	assertHexCode(hexInstr[13], 0x38400042)
	// shifting
	assertHexCode(hexInstr[14], 0x40600042)
	assertHexCode(hexInstr[15], 0x48a00042)
	// comparison
	assertHexCode(hexInstr[16], 0x506000a2)
	assertHexCode(hexInstr[17], 0x586000c2)
	assertHexCode(hexInstr[18], 0x60400042)
	// memory
	assertHexCode(hexInstr[19], 0x68400002)
	assertHexCode(hexInstr[20], 0x70800001)
	// jmp
	assertHexCode(hexInstr[21], 0x7c000320)
	assertHexCode(hexInstr[22], 0x80800000)
	assertHexCode(hexInstr[23], 0x88000000)
	// system
	assertHexCode(hexInstr[24], 0x90000001)
	// stop
	assertHexCode(hexInstr[25], 0x00000000)
}

func TestHighestRegister(t *testing.T) {
	filename := "../../testdata/allinstructions.asm"
	lines, _ := ProgramFileToStringArray(filename)

	hiReg := GetHighestRegister(lines)
	if hiReg != 2 {
		t.Error("Not detecting correct number of required registers")
	}
}

func TestLoadingError(t *testing.T) {
	filename := "../../testdata/notexistingfile"
	_, err := ProgramFileToStringArray(filename)

	if err == nil {
		t.Error("Bad handling of loading file errors")
	}
}

func TestBinaryComplements(t *testing.T) {
	number := -3
	compDeux := BinaryComplement(number, 3*4+1)
	if compDeux != 0x1003 {
		t.Errorf("Bad binary complement conversion - Expected: 0x1003, got: 0x%04x\n", compDeux)
	}
}