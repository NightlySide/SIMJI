package assembler

import (
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"strings"
)

// ProgramFileToStringArray permet de charger le contenu d'un fichier assembleur (.asm)
func ProgramFileToStringArray(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error().Msgf("ProgramFileToStringArray: %s", err.Error())
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

	log.Debug().Msg("Translating ASM to hex instr")

	var pc int

	for _, line := range lines {
		_, _, rest := containsLabel(line)
		_, _, rest = containsComment(rest)

		// si la ligne n'est pas vide et qu'il a une instruction
		if rest != "" {
			opName, args := splitInstruction(rest)

			var numInstr []int
			// on ajoute le numéro d'instruction depuis la liste
			numInstr = append(numInstr, OpCodes[opName])

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
					log.Error().Int("nb_args", len(args)).Msg("Wrong number of arguments!")
					break
				}
			}

			log.Debug().
				Int("pc", pc).
				Ints("instr", numInstr).
				Msg(strings.Join(args, " "))

			numInstructions = append(numInstructions, numInstr)
			pc++
		}
	}

	return numInstructions
}

// ComputeHexInstructions traduit les instructions machines en code hexadécimal
func ComputeHexInstructions(numInstructions [][]int) []int {
	log.Debug().Msg("Translate to HEX instructions")

	var decInstructions []int

	for pc, instr := range numInstructions {
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
		log.Debug().Int("pc", pc).Msgf("Hex: 0x%08x\n", decInstr)
	}

	return decInstructions
}
