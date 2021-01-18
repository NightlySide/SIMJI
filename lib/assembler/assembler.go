package assembler

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// OP_CODES représente tous les codes associés aux instructions
var OP_CODES = map[string]int{
	"stop": 	0,		"add": 		1, 
	"sub": 		2, 		"mult": 	3, 
	"div": 		4,		"and": 		5, 
	"or": 		6,		"xor": 		7, 
	"shl": 		8, 		"shr": 		9, 
	"slt": 		10, 	"sle": 		11, 
	"seq": 		12, 	"load": 	13,
	"store": 	14,		"jmp": 		15, 
	"braz": 	16, 	"branz":	17, 
	"scall": 	18,
}

// parseArgument permet de parser un argument et de dire si c'est un registre 
func parseArgument(argument string, labels map[string]int) (int, bool) {
	// on retire le "r" du registre si il est présent
	if argument[0] == 'r' {
		// on essaie de parser l'argument
		value, err := strconv.Atoi(argument[1:])
		// si il y a une erreur c'est un mauvais registre
		if err != nil { 
			log.Fatal("Error while parsing: ", argument)
		}
		return value, true
	}

	// si il s'agit d'un label on retourne son adresse
	if value, ok := labels[argument]; ok { return value, false }

	// on essaie de parser l'argument
	value, err := strconv.Atoi(argument)

	// si il a une erreur c'est qu'on ne sait pas quelle est cette valeur
	if err != nil { 
		log.Fatal("Error while parsing: ", argument)
	}

	// sinon on a réussi à parser la valeur
	return value, false
}

// LoadASM permet de charger le contenu d'un fichier assembleur (.asm)
func LoadASM(filename string) []string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return lines
}

// GetHighestRegister permet de récupérer le registre le plus haut nécessaire au programme
func GetHighestRegister(lines []string) int {
	var max int
	for _, line := range lines {
		asmInstr := strings.Split(strings.TrimSpace(line), " ")
		for _, instr := range asmInstr {
			if instr[0] == 'r' {
				value, err := strconv.Atoi(instr[1:])
				if err == nil {
					if value > max {
						max = value
					}
				}
			}
		}
	}
	return max
}

func loadLabels(lines []string, debug ...bool) map[string]int {
	var showDebug bool
	if len(debug) >=1 { showDebug = debug[0] }
	if showDebug { fmt.Println("===Loading Labels Dictionary===") }

	var labels = make(map[string]int)
	var pc int = 0

	for _, line := range lines {
		// on a trouvé un label
		if line[len(line) - 1] == ':' {
			if showDebug { fmt.Println("Found label: ", line[:len(line)-1], "\twith address: ", pc) }
			// on l'ajoute au dict des labels
			labels[line[:len(line)-1]] = pc
		} else if line[0] != ';' { pc++ }
		// on incrémente le compteur uniquement si c'est une instruction
		// et que ça n'est pas un commentaire
	}

	return labels
}

// AsmInstructions traduit des instructions asm en instructions machine
func AsmInstructions(lines []string, debug ...bool) [][]int {
	var numInstructions [][]int
	var labels = loadLabels(lines)

	var showDebug bool
	if len(debug) >=1 { showDebug = debug[0] }

	if showDebug { fmt.Println("===Translating ASM to hex instr===") }

	var pc int

	for _, line := range lines {
		// si on n'est pas sur un label ni un commentaire
		if line[len(line) - 1] != ':' && line[0] != ';' {
			if showDebug { fmt.Printf("%03d\t", pc) }

			asmInstr := strings.Split(strings.TrimSpace(line), " ")
			
			var numInstr []int
			// on ajoute le numéro d'instruction depuis la liste
			numInstr = append(numInstr, OP_CODES[asmInstr[0]])
			if showDebug { fmt.Print(asmInstr[0] + "\t") }

			var value int
			var isReg bool
			if len(asmInstr) > 1 {
				switch len(asmInstr) - 1{
					case 1:
						value, isReg = parseArgument(asmInstr[1], labels)
						numInstr = append(numInstr, value)
						break
					case 2:
						if (asmInstr[0] == "jmp") {
							value, isReg = parseArgument(asmInstr[1], labels)
							var res int
							if !isReg { res = 1 }
							numInstr = append(numInstr, res)
							numInstr = append(numInstr, value)

							value, _ = parseArgument(asmInstr[2], labels)
							numInstr = append(numInstr, value)
						} else {
							value, _ = parseArgument(asmInstr[1], labels)
							numInstr = append(numInstr, value)

							value, _ = parseArgument(asmInstr[2], labels)
							numInstr = append(numInstr, value)
						}
						break
					case 3:
						// r1
						value, isReg = parseArgument(asmInstr[1], labels)
						numInstr = append(numInstr, value)

						// o
						// 0 si un registre, 1 si valeur immédiate
						value, isReg = parseArgument(asmInstr[2], labels)
						var imm int
						if !isReg { imm = 1 }
						numInstr = append(numInstr, imm)
						numInstr = append(numInstr, value)

						// r2
						value, isReg = parseArgument(asmInstr[3], labels)
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
				decInstr += instr[2] << (4*2) // o
				decInstr += instr[3] // r
			case 5:
				// add, load, store ...
				decInstr += instr[1] << (4*4) // reg
				decInstr += instr[2] << (4*4 - 1) // imm
				decInstr += instr[3] << (4*2) // o
				decInstr += instr[4] // reg
				break
		}

		decInstructions = append(decInstructions, decInstr)
		if showDebug { fmt.Printf("0x%08x\n", decInstr) }
	}

	return decInstructions
}