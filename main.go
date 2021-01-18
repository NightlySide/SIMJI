package main

import (
	"flag"
	"fmt"
	"os"

	"simji/lib/assembler"
	"simji/lib/gui"
	"simji/lib/vm"

	"github.com/markbates/pkger"
)

func main() {
	flag.Parse()

	if (*launchGUI) {
		fmt.Println("Launching gui...")
		// Include static files for packaging
		staticFiles := pkger.Dir("/static")
		gui.ShowGUI(staticFiles)
	} else {
		args := flag.Args()
		if len(args) < 1 {
			missingFileMessage()
			os.Exit(1);
		}
		
		lines := assembler.ASMToStringArray(args[0])
		// numReg := assembler.GetHighestRegister(lines) + 1 // +1 for the r0 that stays at 0
		numInstructions := assembler.AsmInstructions(lines, *showDebug)
		prog := assembler.ComputeHexInstructions(numInstructions, *showDebug)

		if *showDebug { 
			fmt.Println("\n\n===Launching VM===") 
			fmt.Println("-- Creating VM with: ", 32, " registers")
		}
		vm := vm.NewVM(32, 1000)
		
		vm.LoadProg(prog)
		vm.Run(*showRegs, *showMem, *showDebug)
	}
}