package vm

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

func getNumberInput() int {
	stdin := bufio.NewReader(os.Stdin)

	var value int

	for {
		fmt.Print("[SCALL 0] Enter : R1 <= ")
		_, err := fmt.Fscan(stdin, &value)
		if err == nil {
			break
		}

		stdin.ReadString('\n')
		log.Error().Msg("Input not valid. Please enter a NUMBER.")
	}

	return value
}

// ReverseBinaryComplement permet de retourner un entier signé
// a partir d'un nombre binaire en complément à 2
func ReverseBinaryComplement(number int, size int) int {
	bit := number >> (size - 1)
	if bit == 1 {
		return -1 * (number - 1<<(size-1))
	}
	return number
}
