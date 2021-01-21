package gui

import (
	"fmt"
	"simji/internal/assembler"
	"simji/internal/log"
	"simji/internal/vm"
	"strconv"
	"strings"

	"github.com/zserge/lorca"
)

// BindingsManager permet de gérer les données entre le programme
// et l'interface graphique
type BindingsManager struct {
	vm vm.VM
	ui lorca.UI
}

func newBindingManager(ui lorca.UI) *BindingsManager {
	bm := BindingsManager{ui: ui}
	return &bm
}

func (bm *BindingsManager) setupBindings() {
	// When the UI is ready
	bm.ui.Bind("start", func() {
		log.GetLogger().Info("UI is ready")
		bm.ui.Eval("consoleHelloWorld()")
	})

	// Create and bind Go object to the UI
	bm.ui.Bind("sendProgramContent", bm.loadProgramContent)
	bm.ui.Bind("runCode", bm.runProg)
	bm.ui.Bind("runStep", bm.runStep)
}

func (bm *BindingsManager) loadProgramContent(content string) {
	fmt.Println(content)

	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	//numRegs := assembler.GetHighestRegister(lines) + 1
	vm := vm.NewVM(32, 1000)

	numInstructions := assembler.StringLinesToInstructions(lines, true)
	prog := assembler.ComputeHexInstructions(numInstructions, true)
	vm.LoadProg(prog)

	for _, line := range prog {
		bm.ui.Eval(fmt.Sprintf("printConsole('%x')", line))
	}

	bm.vm = vm

	// lets unlock everything
	bm.ui.Eval("printConsole('[+] Buffer loaded into the program')")
	bm.ui.Eval("unlockButtons()")
	bm.update()
}

func (bm *BindingsManager) runProg() {
	fmt.Println(len(bm.vm.GetProg()))
	bm.vm.RunWithCallback(bm.update)
}

func (bm *BindingsManager) runStep() {
	fmt.Println("Doing a step")
	bm.vm.Step()
	bm.update()
}

func (bm *BindingsManager) update() {
	nbLines := strconv.Itoa(len(bm.vm.GetProg()))
	nbCycles := strconv.Itoa(bm.vm.GetCycles())
	pc := strconv.Itoa(bm.vm.GetPC())
	bm.ui.Eval("updateStats(" + nbLines + "," + nbCycles + "," + pc + ")")

	regsStr := "["
	for _, reg := range bm.vm.GetRegs() {
		regsStr += strconv.Itoa(reg) + ","
	}
	regsStr = regsStr[:len(regsStr) - 1] + "]"
	bm.ui.Eval("setRegisters(" + regsStr + ")")

	memsStr := "["
	for _, mem := range bm.vm.GetMemory() {
		memsStr += strconv.Itoa(mem) + ","
	}
	memsStr = memsStr[:len(memsStr) - 1] + "]"
	bm.ui.Eval("setMemoryBlocks(" + memsStr + ")")
}