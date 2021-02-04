package vm

import "fmt"

// Disassemble permet de desassembler des instructions
// hexadécimal vers un code lisible pour les humains
func Disassemble(prog []int) []string {
	var desProg []string

	for _, instr := range prog {
		var strInstr string
		instrNum, imm1, o1, r1, imm2, o2, r2, a, n := decode(instr)
		switch instrNum {
		case 0:
			strInstr = "stop"
			break
		case 1:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("add r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("add r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 2:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("sub r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("sub r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 3:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("mult r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("mult r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 4:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("div r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("div r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 5:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("and r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("and r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 6:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("or  r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("or  r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 7:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("xor r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("xor r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 8:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("shl r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("shl r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 9:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("shr r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("shr r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 10:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("slt r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("slt r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 11:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("sle r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("sle r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 12:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("seq r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("seq r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 13:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("load r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("load r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 14:
			// registre sinon immediate
			if imm2 == 0 {
				strInstr = fmt.Sprintf("store r%d, r%d, r%d", r1, o2, r2)
			} else {
				strInstr = fmt.Sprintf("store r%d, %d, r%d", r1, o2, r2)
			}
			break
		case 15:
			// registre sinon immediate
			if imm1 == 0 {
				strInstr = fmt.Sprintf("jmp r%d, r%d", o1, r2)
			} else {
				strInstr = fmt.Sprintf("jmp %d, r%d", o1, r2)
			}
			break
		case 16:
			strInstr = fmt.Sprintf("braz r%d, %d", r1, a)
			break
		case 17:
			strInstr = fmt.Sprintf("branz r%d, %d", r1, a)
			break
		case 18:
			strInstr = fmt.Sprintf("scall %d", n)
			break
		}
		desProg = append(desProg, strInstr)
	}
	return desProg
}
