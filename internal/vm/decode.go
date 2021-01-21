package vm

import (
	"fmt"
	"simji/internal/log"
)

func decode(instruction int) (int, int, int, int, int, int, int, int, int) {
	log.GetLogger().Debug(fmt.Sprintf("Decoding : %08x  -  ", instruction))
	instrNum  := (instruction & 0xF8000000) >> 27
	imm1	  := (instruction & 0x04000000) >> 26
	o1		  := (instruction & 0x03FFFFE0) >> 5
	r1        := (instruction & 0x07C00000) >> 22
	imm2	  := (instruction & 0x00200000) >> 21
	o2        := (instruction & 0x001FFFE0) >> 5
	r2        := (instruction & 0x0000001F)
	a 		  := (instruction & 0x003FFFFF)
	n 		  := (instruction & 0x07FFFFFF)

	o1 = BinaryComplement(o1, 21)
	o2 = BinaryComplement(o2, 16)

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