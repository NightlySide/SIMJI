package vm

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

func (vm *VM) eval(instrNum int, imm1 int, o1 int, r1 int, imm2 int, o2 int, r2 int, a int, n int) {
	switch instrNum {
	case 0:
		log.Debug().Msg("stop")
		vm.running = false
		break
	case 1:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("add r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] + vm.regs[o2]
		} else {
			log.Debug().Msgf("add r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] + o2
		}
		break
	case 2:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("sub r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] - vm.regs[o2]
		} else {
			log.Debug().Msgf("sub r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] - o2
		}
		break
	case 3:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("mult r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] * vm.regs[o2]
		} else {
			log.Debug().Msgf("mult r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] * o2
		}
		// cette opération vaut 2 cycles
		vm.cycles++
		break
	case 4:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("div r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] / vm.regs[o2]
		} else {
			log.Debug().Msgf("div r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] / o2
		}
		break
	case 5:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("and r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] & vm.regs[o2]
		} else {
			log.Debug().Msgf("and r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] & o2
		}
		break
	case 6:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("or  r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] | vm.regs[o2]
		} else {
			log.Debug().Msgf("or  r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] | o2
		}
		break
	case 7:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("xor r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] ^ vm.regs[o2]
		} else {
			log.Debug().Msgf("xor r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] ^ o2
		}
		break
	case 8:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("shl r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] << vm.regs[o2]
		} else {
			log.Debug().Msgf("shl r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] << o2
		}
		break
	case 9:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("shr r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] >> vm.regs[o2]
		} else {
			log.Debug().Msgf("shr r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.regs[r1] >> o2
		}
		break
	case 10:
		var res int
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("slt r%d r%d r%d", r1, o2, r2)
			if vm.regs[r1] < vm.regs[o2] {
				res = 1
			}
			vm.regs[r2] = res
		} else {
			log.Debug().Msgf("slt r%d #%d r%d", r1, o2, r2)
			if vm.regs[r1] < o2 {
				res = 1
			}
			vm.regs[r2] = res
		}
		break
	case 11:
		var res int
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("sle r%d r%d r%d", r1, o2, r2)
			if vm.regs[r1] <= vm.regs[o2] {
				res = 1
			}
			vm.regs[r2] = res
		} else {
			log.Debug().Msgf("sle r%d #%d r%d", r1, o2, r2)
			if vm.regs[r1] <= o2 {
				res = 1
			}
			vm.regs[r2] = res
		}
		break
	case 12:
		var res int
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("seq r%d r%d r%d", r1, o2, r2)
			if vm.regs[r1] == vm.regs[o2] {
				res = 1
			}
			vm.regs[r2] = res
		} else {
			log.Debug().Msgf("seq r%d #%d r%d", r1, o2, r2)
			if vm.regs[r1] == o2 {
				res = 1
			}
			vm.regs[r2] = res
		}
		break
	case 13:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("load r%d r%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.memory.GetValueFromIndex(r1+vm.regs[o2])
		} else {
			log.Debug().Msgf("load r%d #%d r%d", r1, o2, r2)
			vm.regs[r2] = vm.memory.GetValueFromIndex(r1+o2)
		}
		break
	case 14:
		// registre sinon immediate
		if imm2 == 0 {
			log.Debug().Msgf("store r%d r%d r%d", r1, o2, r2)
			vm.memory.SetValueFromIndex(vm.regs[r1]+vm.regs[o2], vm.regs[r2])
		} else {
			log.Debug().Msgf("store r%d #%d r%d", r1, o2, r2)
			vm.memory.SetValueFromIndex(vm.regs[r1]+o2, vm.regs[r2])
		}
		break
	case 15:
		// registre sinon immediate
		if imm1 == 0 {
			log.Debug().Msgf("jmp r%d r%d", o1, r2)
			vm.regs[r2] = vm.pc + 1
			vm.pc = vm.regs[o1]
		} else {
			log.Debug().Msgf("jmp #%d r%d", o1, r2)
			vm.regs[r2] = vm.pc + 1
			vm.pc = o1
		}
		// cette opération vaut 2 cycles
		vm.cycles++
		break
	case 16:
		log.Debug().Msgf("braz r%d #%d", r1, a)
		if vm.regs[r1] == 0 {
			vm.pc = a
		}
		// cette opération vaut 2 cycles
		vm.cycles++
		break
	case 17:
		log.Debug().Msgf("branz r%d #%d", r1, a)
		if vm.regs[r1] != 0 {
			vm.pc = a
		}
		// cette opération vaut 2 cycles
		vm.cycles++
		break
	case 18:
		log.Debug().Msgf("scall %d", n)
		vm.handleSysCall(n)
		// implémenter lecture de chiffre etc...
		break
	default:
		log.Warn().Int("instrNum", instrNum).Msgf("Cannot understand instrNum")
		break
	}

	// le registre r0 vaut toujours 0
	vm.regs[0] = 0
}

func (vm *VM) handleSysCall(callNum int) {
	// les opérations systèmes pouvant être bloquantes
	// on ne comptabilise pas ce temps
	vm.totalTime += time.Since(vm.startTime)
	// print r1 to screen
	switch callNum {
	case 0:
		vm.regs[1] = getNumberInput()
	case 1:
		fmt.Println("[SCALL 1] R1 => ", vm.regs[1])
		break
	case 2:
		if vm.debug {
			vm.showRegs()
			vm.showMem()
		}
		break
	default:
		fmt.Println("System call not recognized...")
		break
	}
	// on reprends le timer une fois la fonction effectuée
	vm.startTime = time.Now()
}
