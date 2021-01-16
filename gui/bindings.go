package gui

import (
	"fmt"
	"log"
	"simji/assembler"
	"simji/vm"
	"strings"

	"github.com/zserge/lorca"
)

// BindingsManager permet de gérer les données entre le programme
// et l'interface graphique
type BindingsManager struct {
	vm vm.VM
	ui lorca.UI
}

func newBindingManager(ui lorca.UI) BindingsManager {
	bm := BindingsManager{ui: ui}
	return bm
}

func (bm *BindingsManager) setupBindings() {
	// When the UI is ready
	bm.ui.Bind("start", func() {
		log.Println("UI is ready")
		bm.ui.Eval("consoleHelloWorld()")
	})

	// Create and bind Go object to the UI
	bm.ui.Bind("printHello", bm.helloWorld)
	bm.ui.Bind("sendProgramContent", bm.loadProgramContent)
	bm.ui.Bind("runCode", bm.runProg)
	bm.ui.Bind("getProgLines", func() int {
		return len(bm.vm.GetProg())
	})
}

func (bm BindingsManager) helloWorld() {
	log.Println("Hello world !")
}

func (bm *BindingsManager) loadProgramContent(content string) {
	fmt.Println(content)

	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	numRegs := assembler.GetHighestRegister(lines) + 1
	vm := vm.NewVM(numRegs, 256)

	numInstructions := assembler.AsmInstructions(lines, true)
	prog := assembler.ComputeHexInstructions(numInstructions, true)
	vm.LoadProg(prog)

	for _, line := range prog {
		bm.ui.Eval(fmt.Sprintf("printConsole('%x')", line))
	}

	bm.vm = vm

	// lets unlock everythin
	bm.ui.Eval("updateStats()")
	bm.ui.Eval("unlockButtons()")
	bm.ui.Eval("printConsole('[+] Buffer loaded into the program')")
}

func (bm *BindingsManager) runProg() {
	fmt.Println(len(bm.vm.GetProg()))
	bm.vm.Run(true, true, true)
}