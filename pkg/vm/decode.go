package vm

import (
	"simji/pkg/log"
)

func decode(instruction int) (int, int, int, int, int, int, int, int, int) {
	logger := log.GetLogger()
	if logger.GetLevel() <= log.DEBUG {
		logger.Debug("Decoding : %08x  -  ", instruction)
	}
	instrNum := (instruction & 0xF8000000) >> 27
	imm1 := (instruction & 0x04000000) >> 26
	o1 := (instruction & 0x03FFFFE0) >> 5
	r1 := (instruction & 0x07C00000) >> 22
	imm2 := (instruction & 0x00200000) >> 21
	o2 := (instruction & 0x001FFFE0) >> 5
	r2 := (instruction & 0x0000001F)
	a := (instruction & 0x003FFFFF)
	n := (instruction & 0x07FFFFFF)

	o1 = ReverseBinaryComplement(o1, 21)
	o2 = ReverseBinaryComplement(o2, 16)

	return instrNum, imm1, o1, r1, imm2, o2, r2, a, n
}
