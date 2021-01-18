package vm

import (
	"fmt"
)

// VM est une machine virtuelle
type VM struct {
	pc int
	numReg int
	numMemReg int
	regs []int
	mems []int
	cycles int
	prog []int
	running bool
	debug bool
}

// NewVM permet de créer une nouvelle machine virtuelle
func NewVM(numReg int, numMemReg int) VM {
	vm := VM{numReg: numReg, numMemReg: numMemReg}
	vm.regs = make([]int, vm.numReg)
	vm.mems = make([]int, vm.numMemReg)

	return vm
}

// LoadProg charge une liste d'instruction dans le programme de la VM
func (vm *VM) LoadProg(prog []int) {
	vm.prog = prog
	vm.pc = 0
}

// GetProg retourne le contenu du programme chargé
func (vm VM) GetProg() []int { return vm.prog }

func (vm *VM) fetch() int {
	instruction := vm.prog[vm.pc]
	vm.pc++
	return instruction
}

func (vm VM) showRegs() {
	res := "regs ="
	for k := 0; k < vm.numReg; k++ {
		res += " " + fmt.Sprintf("%04x", vm.regs[k])
	}
	fmt.Println(res)
}

func (vm VM) showMem() {
	res := "memory ="
	for k := 0; k < vm.numMemReg; k++ {
		res += " " + fmt.Sprintf("%04x", vm.mems[k])
	}
	fmt.Println(res)
}

// Run permet de lancer l'exécution de la machine virtuelle
func (vm *VM) Run(showRegs bool, showMem bool, debug bool) {
	vm.running = true
	vm.debug = debug
	for vm.running {
		instruction := vm.fetch()
		vm.eval(vm.decode(instruction))
		if showRegs { vm.showRegs() }
		if showMem { vm.showMem() }
	}
}