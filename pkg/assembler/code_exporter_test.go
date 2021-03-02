package assembler

import (
	"os"
	"github.com/Nightlyside/simji/pkg/log"
	"testing"
)

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
	output := log.CaptureOutput(func() { PrintProgram([]int{0x00000000}) })
	if output != "0x00000000\t0x00000000\n" {
		t.Errorf("Wrong output - Expected: 0x00000000\t0x00000000, Got: %s\n", output)
	}
}
