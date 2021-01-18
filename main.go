package main

import (
	"flag"
	"os"

	"simji/lib/assembler"
	"simji/lib/gui"
	"simji/lib/log"
	"simji/lib/vm"

	"github.com/markbates/pkger"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if (*showDebug) {
		log.GetLogger().SetLevel(log.DEBUG)
	} else {
		log.GetLogger().SetLevel(log.WARN)
	}

	if (*launchGUI) {
		log.GetLogger().Info("Launching gui...")
		// Include static files for packaging
		staticFiles := pkger.Dir("/static")
		gui.ShowGUI(staticFiles)
	} else {
		// If there is not enough arguments
		if len(args) < 1 {
			missingFileMessage()
			os.Exit(1);
		}
		
		// Else we load the program
		lines := assembler.ASMToStringArray(args[0])
		numInstructions := assembler.AsmInstructions(lines, *showDebug)
		prog := assembler.ComputeHexInstructions(numInstructions, *showDebug)

			log.GetLogger().DebugTitle("Launching VM") 
			log.GetLogger().Debug("-- Creating VM with: 32 registers")
		vm := vm.NewVM(32, 1000)
		
		vm.LoadProg(prog)
		vm.Run(*showRegs, *showMem, *showDebug)
	}
}