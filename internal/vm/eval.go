package vm

import "fmt"

func (vm *VM) eval(instrNum int, imm1 int, o1 int, r1 int, imm2 int, o2 int, r2 int, a int, n int) {
	switch(instrNum) {
		case 0:
			if (vm.debug) { fmt.Printf("stop\n") }
			vm.running = false
			break
		case 1:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("add r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] + vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("add r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] + o2
			}
			break
		case 2:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("sub r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] - vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("sub r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] - o2
			}
			break
		case 3:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("mult r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] * vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("mult r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] * o2
			}
			break
		case 4:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("div r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] / vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("div r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] / o2
			}
			break
		case 5:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("and r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] & vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("and r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] & o2
			}
			break;
		case 6:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("or  r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] | vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("or  r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] | o2
			}
			break
		case 7:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("xor r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] ^ vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("xor r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] ^ o2
			}
			break
		case 8:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("shl r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] << vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("shl r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] << o2
			}
			break
		case 9:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("shr r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] >> vm.regs[o2]
			} else {
				if (vm.debug) { fmt.Printf("shr r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.regs[r1] >> o2
			}
			break
		case 10:
			var res int
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("slt r%d r%d r%d\n", r1, o2, r2) }
				if vm.regs[r1] < vm.regs[o2] {res = 1}
				vm.regs[r2] = res
			} else {
				if (vm.debug) { fmt.Printf("slt r%d #%d r%d\n", r1, o2, r2) }
				if vm.regs[r1] < o2 {res = 1}
				vm.regs[r2] = res
			}
			break
		case 11:
			var res int
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("sle r%d r%d r%d\n", r1, o2, r2) }
				if vm.regs[r1] <= vm.regs[o2] {res = 1}
				vm.regs[r2] = res
			} else {
				if (vm.debug) { fmt.Printf("sle r%d #%d r%d\n", r1, o2, r2) }
				if vm.regs[r1] <= o2 {res = 1}
				vm.regs[r2] = res
			}
			break
		case 12:
			var res int
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("seq r%d r%d r%d\n", r1, o2, r2) }
				if vm.regs[r1] == vm.regs[o2] {res = 1}
				vm.regs[r2] = res
			} else {
				if (vm.debug) { fmt.Printf("seq r%d #%d r%d\n", r1, o2, r2) }
				if vm.regs[r1] == o2 {res = 1}
				vm.regs[r2] = res
			}
			break
		case 13:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("load r%d r%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.mems[r1 + vm.regs[o2]]
			} else {
				if (vm.debug) { fmt.Printf("load r%d #%d r%d\n", r1, o2, r2) }
				vm.regs[r2] = vm.mems[r1 + o2]
			}
			break
		case 14:
			// registre sinon immediate
			if (imm2 == 0) {
				if (vm.debug) { fmt.Printf("store r%d r%d r%d\n", r1, o2, r2) }
				vm.mems[vm.regs[r1] + vm.regs[o2]] = vm.regs[r2]
			} else {
				if (vm.debug) { fmt.Printf("store r%d #%d r%d\n", r1, o2, r2) }
				vm.mems[vm.regs[r1] + o2] = vm.regs[r2]
			}
			break
		case 15:
			// registre sinon immediate
			if (imm1 == 0) {
				if (vm.debug) { fmt.Printf("jmp r%d r%d\n", o1, r2) }
				vm.regs[r2] = vm.pc + 1
				vm.pc = vm.regs[o1]
			} else {
				if (vm.debug) { fmt.Printf("jmp #%d r%d\n", o1, r2) }
				vm.regs[r2] = vm.pc + 1
				vm.pc = o1
			}
			break
		case 16:
			if (vm.debug) { fmt.Printf("braz r%d #%d\n", r1, a) }
			if vm.regs[r1] == 0 {
				vm.pc = a
			}
			break
		case 17:
			if (vm.debug) { fmt.Printf("branz r%d #%d\n", r1, a) }
			if vm.regs[r1] != 0 {
				vm.pc = a
			}
			break
		case 18:
			if (vm.debug) { fmt.Printf("scall %d\n", n) }
			// print r1 to screen
			switch (n) {
				case 0:
					var i int
					validInput := false
					for !validInput {
						fmt.Print("[SCALL 0] Enter : R1 <= ")
						_, err := fmt.Scanf("%d", &i)
						if err == nil { validInput = true } else {
							fmt.Println("\nPlease enter a NUMBER!")
						}
					}
					vm.regs[1] = i
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
			// implémenter lecture de chiffre etc...
			break
		default:
			if (vm.debug) { fmt.Println("Cannot understand instrNum : ", instrNum) }
			break
	}

	// le registre r0 vaut toujours 0
	vm.regs[0] = 0
}