package vm

import (
	"fmt"
	"simji/internal/log"
	"time"
)

func (vm *VM) eval(instrNum int, imm1 int, o1 int, r1 int, imm2 int, o2 int, r2 int, a int, n int) {
	switch(instrNum) {
		case 0:
			log.GetLogger().Debug("stop\n") 
			vm.running = false
			break
		case 1:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("add r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] + vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("add r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] + o2
			}
			break
		case 2:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("sub r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] - vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("sub r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] - o2
			}
			break
		case 3:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("mult r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] * vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("mult r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] * o2
			}
			break
		case 4:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("div r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] / vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("div r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] / o2
			}
			break
		case 5:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("and r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] & vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("and r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] & o2
			}
			break;
		case 6:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("or  r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] | vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("or  r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] | o2
			}
			break
		case 7:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("xor r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] ^ vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("xor r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] ^ o2
			}
			break
		case 8:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("shl r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] << vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("shl r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] << o2
			}
			break
		case 9:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("shr r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] >> vm.regs[o2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("shr r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.regs[r1] >> o2
			}
			break
		case 10:
			var res int
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("slt r%d r%d r%d\n", r1, o2, r2))
				if vm.regs[r1] < vm.regs[o2] {res = 1}
				vm.regs[r2] = res
			} else {
				log.GetLogger().Debug(fmt.Sprintf("slt r%d #%d r%d\n", r1, o2, r2))
				if vm.regs[r1] < o2 {res = 1}
				vm.regs[r2] = res
			}
			break
		case 11:
			var res int
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("sle r%d r%d r%d\n", r1, o2, r2))
				if vm.regs[r1] <= vm.regs[o2] {res = 1}
				vm.regs[r2] = res
			} else {
				log.GetLogger().Debug(fmt.Sprintf("sle r%d #%d r%d\n", r1, o2, r2))
				if vm.regs[r1] <= o2 {res = 1}
				vm.regs[r2] = res
			}
			break
		case 12:
			var res int
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("seq r%d r%d r%d\n", r1, o2, r2))
				if vm.regs[r1] == vm.regs[o2] {res = 1}
				vm.regs[r2] = res
			} else {
				log.GetLogger().Debug(fmt.Sprintf("seq r%d #%d r%d\n", r1, o2, r2))
				if vm.regs[r1] == o2 {res = 1}
				vm.regs[r2] = res
			}
			break
		case 13:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("load r%d r%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.mems[r1 + vm.regs[o2]]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("load r%d #%d r%d\n", r1, o2, r2))
				vm.regs[r2] = vm.mems[r1 + o2]
			}
			break
		case 14:
			// registre sinon immediate
			if (imm2 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("store r%d r%d r%d\n", r1, o2, r2))
				vm.mems[vm.regs[r1] + vm.regs[o2]] = vm.regs[r2]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("store r%d #%d r%d\n", r1, o2, r2))
				vm.mems[vm.regs[r1] + o2] = vm.regs[r2]
			}
			break
		case 15:
			// registre sinon immediate
			if (imm1 == 0) {
				log.GetLogger().Debug(fmt.Sprintf("jmp r%d r%d\n", o1, r2))
				vm.regs[r2] = vm.pc + 1
				vm.pc = vm.regs[o1]
			} else {
				log.GetLogger().Debug(fmt.Sprintf("jmp #%d r%d\n", o1, r2))
				vm.regs[r2] = vm.pc + 1
				vm.pc = o1
			}
			break
		case 16:
			log.GetLogger().Debug(fmt.Sprintf("braz r%d #%d\n", r1, a))
			if vm.regs[r1] == 0 {
				vm.pc = a
			}
			break
		case 17:
			log.GetLogger().Debug(fmt.Sprintf("branz r%d #%d\n", r1, a))
			if vm.regs[r1] != 0 {
				vm.pc = a
			}
			break
		case 18:
			log.GetLogger().Debug(fmt.Sprintf("scall %d\n", n))
			vm.handleSysCall(n)
			// implémenter lecture de chiffre etc...
			break
		default:
			log.GetLogger().Warn(fmt.Sprintf("Cannot understand instrNum : %d\n", instrNum))
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
	switch (callNum) {
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