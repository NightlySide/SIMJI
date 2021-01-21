package vm

import (
	"fmt"
	"simji/internal/log"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	startTime time.Time
	totalTime time.Duration
}

// NewVM permet de créer une nouvelle machine virtuelle
func NewVM(numReg int, numMemReg int) VM {
	vm := VM{numReg: numReg, numMemReg: numMemReg}
	vm.regs = make([]int, vm.numReg)
	vm.mems = make([]int, vm.numMemReg)

	return vm
}

// GetProg retourne le contenu du programme chargé
func (vm VM) GetProg() []int { return vm.prog }

// GetCycles retourne le nombre du cycle de programme
func (vm VM) GetCycles() int { return vm.cycles }

// GetPC retourne le compteur du programme
func (vm VM) GetPC() int { return vm.pc }

// GetRegs le contenu des registres 
func (vm VM) GetRegs() []int { return vm.regs }

// GetMemory le contenu de la mémoire 
func (vm VM) GetMemory() []int { return vm.mems }

func (vm *VM) fetch() int {
	instruction := vm.prog[vm.pc]
	vm.pc++
	vm.cycles++
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

func (vm *VM) printPerf() {
	log.GetLogger().Title(log.INFO, "Performances") 
	fmt.Printf("Nombre de cycles : %d\n", vm.cycles)
	fmt.Printf("Effectués en : %s\n", vm.totalTime)

	p := message.NewPrinter(language.French)

	opParSeconde := int64(float64(vm.cycles) / vm.totalTime.Seconds())
	p.Printf("Perfomances : %d opérations/seconde\n", opParSeconde)
	fmt.Println("===================")
}

// Run permet de lancer l'exécution de la machine virtuelle
func (vm *VM) Run(showRegs bool, showMem bool, debug bool, showPerfs bool) {
	vm.running = true
	vm.debug = debug
	vm.startTime = time.Now()
	for vm.running {
		vm.Step()
		if showRegs { vm.showRegs() }
		if showMem { vm.showMem() }
	}
	vm.totalTime += time.Since(vm.startTime)

	if (showPerfs) { vm.printPerf() }
}

// RunWithCallback permet de lancer l'exécution de la machine virtuelle
// en appelant une fonction à chaque itération
func (vm *VM) RunWithCallback(callback func()) {
	vm.running = true
	vm.debug = false
	for vm.running {
		vm.Step()
		//vm.showRegs()
		//vm.showMem()
		callback();
	}
}

// Step permet de faire une itération du programme
func (vm *VM) Step() {
	instruction := vm.fetch()
	vm.eval(decode(instruction))
}