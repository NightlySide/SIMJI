package vm

import "fmt"

func (vm VM) decode(instruction int) (int, int, int, int, int, int, int, int, int) {
	if (vm.debug) { fmt.Printf("Decoding : %08x  -  ", instruction) }
	instrNum  := (instruction & 0xFF000000) >> (4*6)
	imm1	  := (instruction & 0x00800000) >> (4*6-1)
	o1		  := (instruction & 0x007FFF00) >> (4*2)
	r1        := (instruction & 0x00FF0000) >> (4*4)
	imm2	  := (instruction & 0x00008000) >> (4*4-1)
	o2        := (instruction & 0x00007F00) >> (4*2)
	r2        := (instruction & 0x000000FF)
	a 		  := (instruction & 0x0000FFFF)
	n 		  := (instruction & 0x00FFFFFF)

	o1 = BinaryComplement(o1, 17)
	o2 = BinaryComplement(o2, 7)

	return instrNum, imm1, o1, r1, imm2, o2, r2, a, n
}

// BinaryComplement permet de retourner un entier signé
// a partir d'un nombre binaire en complément à 2
func BinaryComplement(number int, size int) int {
	bit := number >> (size-1)
	if bit == 1 {
		return -1 * (number - 1 << (size-1))
	}
	return number
}