package assembler

import (
	"fmt"
	"io/ioutil"
)

// ExportBinaryToFile permet de sauvegarder des instructions en hexadécimal
// dans un fichier binaire
func ExportBinaryToFile(hexInstructions []int, filename string) error {
	var res string

	for i, instr := range hexInstructions {
		res += fmt.Sprintf("0x%08x\t0x%08x\n", i, instr)
	}

	res = res[:len(res)-1]

	return ioutil.WriteFile(filename, []byte(res), 0)
}

// ExportProgramToFile permet de sauvegarder des instructions lisibles
// dans un fichier programme
func ExportProgramToFile(strInstructions []string, filename string) error {
	var res string

	for _, instr := range strInstructions {
		res += instr + "\n"
	}

	res = res[:len(res)-1]

	return ioutil.WriteFile(filename, []byte(res), 0)
}

// PrintProgram permet d'afficher dans la console les instructions en hexadécimal
func PrintProgram(hexInstructions []int) {
	var res string

	for i, instr := range hexInstructions {
		res += fmt.Sprintf("0x%08x\t0x%08x\n", i, instr)
	}
	fmt.Print(res)
}
