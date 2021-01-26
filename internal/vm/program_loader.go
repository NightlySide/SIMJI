package vm

import (
	"io/ioutil"
	"os"
	"simji/internal/log"
	"strconv"
	"strings"
)

// LoadProgFromFile charge une liste d'instruction dans le programme de la VM
// à partir d'un fichier binaire
func LoadProgFromFile(filename string) []int {
	var prog []int
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.GetLogger().Error(err.Error())
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	for i := range lines {
		data := strings.Split(lines[i], " ")
		data = strings.Split(data[len(data)-1], "\t")
		// idx := strings.TrimSpace(data[0])
		instru := strings.TrimSpace(data[len(data)-1])

		instru = strings.Replace(instru, "0x", "", -1)
		hexInstr, err := strconv.ParseUint(instru, 16, 64)
		if err != nil {
			log.GetLogger().Error(err.Error())
		}

		prog = append(prog, int(hexInstr))
	}

	return prog
}
