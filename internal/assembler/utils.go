package assembler

import (
	"log"
	"strconv"
	"strings"
)

func sanitizeLine(line string) string {
	args := strings.Split(strings.TrimSpace(line), " ")
	var rest []string
	for _, arg := range args {
		if arg != "" && arg != " " {
			rest = append(rest, arg)
		}
	}

	return strings.Join(rest, " ")
}

func containsLabel(line string) (bool, string, string) {
	var contLabel bool
	var rest string
	var label string

	if line == "" {
		return false, "", ""
	}

	args := strings.Split(line, " ")
	if args[0][len(args[0])-1] == ':' {
		contLabel = true
		rest = strings.TrimSpace(strings.Join(args[1:], " "))
		label = args[0][:len(args[0])-1]
	} else {
		rest = line
	}

	return contLabel, label, rest
}

func containsComment(line string) (bool, string, string) {
	var contComment bool
	var rest string
	var comment string

	if line == "" {
		return false, "", ""
	}

	args := strings.Split(line, " ")
	for i, arg := range args {
		if len(arg) > 0 && arg[0] == ';' {
			contComment = true
			rest = strings.Join(args[:i], " ")
			comment = strings.Join(args[i:], " ")[1:]
			break
		}
	}

	if contComment {
		return true, comment, rest
	}

	return false, "", line
}

// parseArgument permet de parser un argument et de dire si c'est un registre
func parseArgument(argument string, labels map[string]int) (int, bool) {
	// on retire le "r" du registre si il est présent
	if argument[0] == 'r' {
		// on essaie de parser l'argument
		value, err := strconv.Atoi(argument[1:])
		// si il y a une erreur c'est un mauvais registre
		if err != nil {
			log.Fatal("Error while parsing register: ", argument)
		}
		return value, true
	}

	// si il s'agit d'un label on retourne son adresse
	if value, ok := labels[argument]; ok {
		return value, false
	}

	// on essaie de parser l'argument
	value, err := strconv.Atoi(argument)
	// si il a une erreur c'est qu'on ne sait pas quelle est cette valeur
	if err != nil {
		log.Fatal("Error while parsing immediate: ", argument)
	}

	// sinon on a réussi à parser la valeur
	return value, false
}

func splitInstruction(instruction string) (string, []string) {
	var opName string
	var args []string

	data := strings.Split(instruction, " ")
	opName = data[0]

	if len(data) > 1 {
		args = strings.Split(strings.Join(data[1:], ""), ",")
	}

	return opName, args
}

// BinaryComplement permet de calculer le complément à 2 d'un entier
func BinaryComplement(number int, size int) int {
	if number >= 0 {
		return number
	}
	return (1 << (size - 1)) - number
}
