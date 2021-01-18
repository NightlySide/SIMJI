package assembler

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// ASMToStringArray permet de charger le contenu d'un fichier assembleur (.asm)
func ASMToStringArray(filename string) []string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = sanitizeLine(lines[i])
	}
	return lines
}

// AsmInstructions traduit des instructions asm en instructions machine
func AsmInstructions(lines []string, debug ...bool) [][]int {
	var numInstructions [][]int
	var labels = loadLabels(lines, debug...)

	var showDebug bool
	if len(debug) >=1 { showDebug = debug[0] }

	if showDebug { fmt.Println("===Translating ASM to hex instr===") }

	var pc int

	for _, line := range lines {
		_, _, rest := containsLabel(line)
		_, _, rest = containsComment(rest)

		// si la ligne n'est pas vide et qu'il a une instruction
		if rest != "" {
			if showDebug { fmt.Printf("%03d\t", pc) }

			opName, args := splitInstruction(rest)
			
			var numInstr []int
			// on ajoute le numéro d'instruction depuis la liste
			numInstr = append(numInstr, OpCodes[opName])
			if showDebug { fmt.Print(opName + "\t") }

			var value int
			var isReg bool
			if len(args) > 0 {
				switch len(args){
					case 1:
						value, isReg = parseArgument(args[0], labels)
						numInstr = append(numInstr, value)
						break
					case 2:
						if (opName == "jmp") {
							value, isReg = parseArgument(args[0], labels)
							var res int
							if !isReg { res = 1 }
							numInstr = append(numInstr, res)
							numInstr = append(numInstr, value)

							value, _ = parseArgument(args[1], labels)
							numInstr = append(numInstr, value)
						} else {
							value, _ = parseArgument(args[0], labels)
							numInstr = append(numInstr, value)

							value, _ = parseArgument(args[1], labels)
							numInstr = append(numInstr, value)
						}
						break
					case 3:
						// r1
						value, isReg = parseArgument(args[0], labels)
						numInstr = append(numInstr, value)

						// o
						// 0 si un registre, 1 si valeur immédiate
						value, isReg = parseArgument(args[1], labels)
						var imm int
						if !isReg { imm = 1 }
						numInstr = append(numInstr, imm)
						numInstr = append(numInstr, value)

						// r2
						value, isReg = parseArgument(args[2], labels)
						numInstr = append(numInstr, value)
						break
					default:
						log.Fatal("Wrong number of arguments !")
						break
				}
			}

			if showDebug { fmt.Println(numInstr) }

			numInstructions = append(numInstructions, numInstr)
			pc++
		}
	}

	return numInstructions
}

// ComputeHexInstructions traduit les instructions machines en code hexadécimal
func ComputeHexInstructions(numInstructions [][]int, debug ...bool) []int {
	var showDebug bool
	if len(debug) >=1 { showDebug = debug[0] }

	if showDebug { fmt.Println("===Translate to HEX instructions===") }

	var decInstructions []int

	for pc, instr := range numInstructions {
		if showDebug { fmt.Printf("%03d\t", pc) }

		decInstr := instr[0] << (4*6)

		switch (len(instr)) {
			case 1:
				break
			case 2:
				// scall
				decInstr += instr[1] // num
				break
			case 3:
				// braz
				decInstr += instr[1] << (4*4) // reg
				decInstr += instr[2] // address
			case 4:
				// jmp
				decInstr += instr[1] << (4*6-1) // imm
				decInstr += BinaryComplement(instr[2], 15) << (4*2)  // o
				decInstr += instr[3] // r
			case 5:
				// add, load, store ...
				decInstr += instr[1] << (4*4) // reg
				decInstr += instr[2] << (4*4 - 1) // imm
				decInstr += BinaryComplement(instr[3], 7) << (4*2) // o
				decInstr += instr[4] // reg
				break
		}

		decInstructions = append(decInstructions, decInstr)
		if showDebug { fmt.Printf("0x%08x\n", decInstr) }
	}

	return decInstructions
}