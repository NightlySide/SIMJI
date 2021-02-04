package assembler

import (
	"fmt"
	"io/ioutil"
	"simji/pkg/log"
	"strings"
)

// ProgramFileToStringArray permet de charger le contenu d'un fichier assembleur (.asm)
func ProgramFileToStringArray(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.GetLogger().Error(err.Error())
		return []string{}, err
	}

	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = sanitizeLine(lines[i])
	}
	return lines, nil
}

// StringLinesToInstructions traduit des instructions asm en instructions machine
func StringLinesToInstructions(lines []string) [][]int {
	var numInstructions [][]int
	labels := loadLabels(lines)

	log.GetLogger().Title(log.DEBUG, "Translating ASM to hex instr")

	var pc int

	for _, line := range lines {
		_, _, rest := containsLabel(line)
		_, _, rest = containsComment(rest)

		// si la ligne n'est pas vide et qu'il a une instruction
		if rest != "" {
			log.GetLogger().Debug("%08x\t", pc)

			opName, args := splitInstruction(rest)

			var numInstr []int
			// on ajoute le numéro d'instruction depuis la liste
			numInstr = append(numInstr, OpCodes[opName])
			log.GetLogger().Debug(opName + "\t")

			var value int
			var isReg bool
			if len(args) > 0 {
				switch len(args) {
				case 1:
					value, isReg = parseArgument(args[0], labels)
					numInstr = append(numInstr, value)
					break
				case 2:
					if opName == "jmp" {
						value, isReg = parseArgument(args[0], labels)
						var res int
						if !isReg {
							res = 1
						}
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
					value, _ = parseArgument(args[0], labels)
					numInstr = append(numInstr, value)

					// o
					// 0 si un registre, 1 si valeur immédiate
					value, isReg = parseArgument(args[1], labels)
					var imm int
					if !isReg {
						imm = 1
					}
					numInstr = append(numInstr, imm)
					numInstr = append(numInstr, value)

					// r2
					value, isReg = parseArgument(args[2], labels)
					numInstr = append(numInstr, value)
					break
				default:
					log.GetLogger().Error("Wrong number of arguments !")
					break
				}
			}

			spacer := "\t"
			if opName == "scall" {
				spacer = "\t\t"
			}
			log.GetLogger().Debug(fmt.Sprint(numInstr) + spacer + strings.Join(args, " ") + "\n")

			numInstructions = append(numInstructions, numInstr)
			pc++
		}
	}

	return numInstructions
}

// ComputeHexInstructions traduit les instructions machines en code hexadécimal
func ComputeHexInstructions(numInstructions [][]int) []int {
	log.GetLogger().Title(log.DEBUG, "Translate to HEX instructions")

	var decInstructions []int

	for pc, instr := range numInstructions {
		log.GetLogger().Debug("%08x\t", pc)

		decInstr := instr[0] << 27

		switch len(instr) {
		case 1:
			break
		case 2:
			// scall
			decInstr += instr[1] // num
			break
		case 3:
			// braz
			decInstr += instr[1] << 22 // reg
			decInstr += instr[2]       // address
		case 4:
			// jmp
			decInstr += instr[1] << 26                      // imm
			decInstr += BinaryComplement(instr[2], 21) << 5 // o
			decInstr += instr[3]                            // r
		case 5:
			// add, load, store ...
			decInstr += instr[1] << 22                      // reg
			decInstr += instr[2] << 21                      // imm
			decInstr += BinaryComplement(instr[3], 16) << 5 // o
			decInstr += instr[4]                            // reg
			break
		}

		decInstructions = append(decInstructions, decInstr)
		log.GetLogger().Debug("0x%08x\n", decInstr)
	}

	return decInstructions
}
