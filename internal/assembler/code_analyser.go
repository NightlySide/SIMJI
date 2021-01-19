package assembler

import (
	"simji/internal/log"
	"strconv"
	"strings"
)

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

func loadLabels(lines []string) map[string]int {
	log.GetLogger().DebugTitle("Loading Labels Dictionary")

	var labels = make(map[string]int)
	var pc int = 0

	for _, line := range lines {
		isLabel, label, rest := containsLabel(line)
		// on a trouvé un label
		if isLabel {
			log.GetLogger().Debug("Found label: " + label + "\twith address: " + strconv.Itoa(pc))
			// on l'ajoute au dict des labels
			labels[label] = pc
		}
		_, _, rest = containsComment(rest)
		// on incrémente le compteur uniquement si il y a une instruction
		if rest != "" {pc++}
	}

	return labels
}