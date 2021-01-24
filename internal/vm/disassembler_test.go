package vm

import (
	"simji/internal/assembler"
	"strings"
	"testing"
)

func TestDisassembly(t *testing.T) {
	content, _ := assembler.ProgramFileToStringArray("../../testdata/allinstructions.asm")
	instructions := assembler.ComputeHexInstructions(assembler.StringLinesToInstructions(content))

	reversed := Disassemble(instructions)

	assertInstruction := func(reversed string, original string) {
		reversed = strings.TrimSpace(reversed)
		original = strings.TrimSpace(original)
		if reversed != original {
			t.Errorf("Instruction mismatch - Expected: %s, Got: %s\n", original, reversed)
		}
	}

	// add
	assertInstruction(reversed[0], "add r1, r0, r2")
	assertInstruction(reversed[1], "add r1, 5, r1")
	// sub
	assertInstruction(reversed[2], "sub r1, r0, r2")
	assertInstruction(reversed[3], "sub r1, 3, r2")
	// mult
	assertInstruction(reversed[4], "mult r1, r1, r2")
	assertInstruction(reversed[5], "mult r1, 5, r2")
	// div
	assertInstruction(reversed[6], "div r1, r1, r2")
	assertInstruction(reversed[7], "div r1, 5, r2")
	// and
	assertInstruction(reversed[8], "and r1, 1, r1")
	assertInstruction(reversed[9], "and r1, r2, r2")
	// or
	assertInstruction(reversed[10], "or  r1, 1, r2")
	assertInstruction(reversed[11], "or  r1, r2, r2")
	// xor
	assertInstruction(reversed[12], "xor r1, 1, r2")
	assertInstruction(reversed[13], "xor r1, r2, r2")
	// shifting
	assertInstruction(reversed[14], "shl r1, 2, r2")
	assertInstruction(reversed[15], "shl r1, r1, r2")
	assertInstruction(reversed[16], "shr r2, 2, r2")
	assertInstruction(reversed[17], "shr r1, r2, r2")
	// comparison
	assertInstruction(reversed[18], "slt r1, 5, r2")
	assertInstruction(reversed[19], "slt r1, r2, r2")
	assertInstruction(reversed[20], "sle r1, 6, r2")
	assertInstruction(reversed[21], "sle r1, r2, r2")
	assertInstruction(reversed[22], "seq r1, r2, r2")
	assertInstruction(reversed[23], "seq r1, 2, r2")
	// memory
	assertInstruction(reversed[24], "load r1, r0, r2")
	assertInstruction(reversed[25], "load r1, 0, r2")
	assertInstruction(reversed[26], "store r2, r0, r1")
	assertInstruction(reversed[27], "store r2, 0, r1")
	// jmp
	assertInstruction(reversed[28], "jmp 33, r0")
	assertInstruction(reversed[29], "jmp r2, r0")
	assertInstruction(reversed[30], "braz r2, 0")
	assertInstruction(reversed[31], "branz r0, 0")
	// system
	assertInstruction(reversed[32], "scall 1")
	// stop
	assertInstruction(reversed[33], "stop")
}
