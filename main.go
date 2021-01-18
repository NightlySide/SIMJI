package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/markbates/pkger"

	"simji/lib/assembler"
	"simji/lib/gui"
	"simji/lib/vm"
)

var showRegs = flag.Bool("show-regs", false, "show regs on each step")
var showMem = flag.Bool("show-mem", false, "show memory on each step")
var showDebug = flag.Bool("debug", false, "show each instruction in the terminal")
var launchGUI = flag.Bool("gui", false, "start the gui application")

func helloWorld() {
	fmt.Println("========= SIMJI : Simulateur de Jeu d'Instructions =========")
	fmt.Println("-- Conçut par Alexandre FROEHLICH")
	fmt.Println("-- Dans le cadre de l'U.V. 4.5-Architectures numériques")
	fmt.Println("-- Contact : nightlyside@gmail.com")
	fmt.Println("-- Site web : https://nightlyside.github.io")
	fmt.Println("============================================================")
	fmt.Print("\n\n")
}

func usage() {
	helloWorld()
	fmt.Fprintf(os.Stderr, "usage: %s [inputfile]\n", os.Args[0])
    flag.PrintDefaults()
    os.Exit(2)
}

func init() {
	flag.Usage = usage
	flag.BoolVar(showRegs, "r", false, "alias for -show-regs")
	flag.BoolVar(showMem, "m", false, "alias for -show-mem")
	flag.BoolVar(showDebug, "d", false, "alias for -debug")
	flag.BoolVar(launchGUI, "g", false, "alias for -gui")
}

func main() {
	// Include static files for packaging
	pkger.Include("/static")
	helloWorld()
	flag.Parse()

	if (*launchGUI) {
		fmt.Println("Launching gui...")
		gui.ShowGUI()
	} else {
		args := flag.Args()
		if len(args) < 1 {
			fmt.Println("Input file is missing.");
			fmt.Printf("type %s -h for help\n", os.Args[0])
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