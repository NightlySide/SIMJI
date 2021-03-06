package main

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog/log"

	"github.com/fatih/color"
)

func helloWorld() {
	defaultColor := color.New(color.FgWhite)
	titleColor := color.New(color.FgHiCyan).Add(color.Underline).Add(color.Bold)
	detailColor := color.New(color.FgBlue)

	defaultColor.Println("                      _____ _______  ___    ______")
	defaultColor.Println("                     / ___//_  _/  |/  /   / /  _/")
	defaultColor.Println("                     \\__  \\ / // /|_/ /_  / // /")
	defaultColor.Println("                     ___/ // // /  / / /_/ // / ")
	defaultColor.Println("                    /____/___/_/  /_/\\____/___/ ")
	fmt.Println()

	defaultColor.Print("============== ")
	titleColor.Printf("SIMJI : Simulateur de Jeu d'Instructions ")
	defaultColor.Print("============== \n")
	defaultColor.Print("-- Designed and Developed by ")
	detailColor.Print("Alexandre FROEHLICH \n")
	defaultColor.Println("-- For the U.V. 4.5-Architectures numériques class")
	defaultColor.Print("-- Contact : ")
	detailColor.Println("nightlyside@gmail.com")
	defaultColor.Print("-- Website : ")
	detailColor.Println("https://nightlyside.github.io")
	defaultColor.Println("======================================================================")
	fmt.Print("\n")
}

func usage() {
	fmt.Println("usage: simji [--help | -h] [--gui | -g] [--debug | -d] [--binary | -b]")
	fmt.Println("             [--show-regs | -r] [--show-mem | -m] [--benchmark | -bm]")
	fmt.Println("             [--assemble | -a] [--disassemble | -da] [--output | -o]")
	fmt.Println("             \"filename\"")
	fmt.Println()
}

func fullUsage() {
	usage()

	fmt.Println("Some common commands in different situations:")
	fmt.Println()

	fmt.Println("start and use the program")
	fmt.Println("  simji myprogram.asm\tLaunch the program \"myprogram.asm\" in CLI mode")
	fmt.Println("  simji --binary myprogram.bin\tLaunch the binary \"myprogram.bin\" in CLI mode")
	fmt.Println("  simji --gui\t\tLaunch the program with a Graphical UI")
	fmt.Println()

	fmt.Println("assemble only the program")
	fmt.Println("  simji --assemble myprogram.asm\tAssemble the program and prints the content in the console")
	fmt.Println("  simji --assemble --output=program.bin myprogram.asm\tAssemble the program into a binary file")
	fmt.Println()

	fmt.Println("disassemble a binary file")
	fmt.Println("  simji --disassemble myprogram.bin\tDisassemble the program and prints the content in the console")
	fmt.Println("  simji --disassemble --output=program.asm myprogram.bin\tDisassemble the program into a text file")
	fmt.Println()

	fmt.Println("debug the loaded program in translation/execution")
	fmt.Println("  simji --show-regs\tShow the registry values when running the program")
	fmt.Println("  simji --show-memory\tShow the memory blocks when running the program")
	fmt.Println("  simji --debug\t\tRun the program in debug mode for loading the program")
	fmt.Println()

	fmt.Println("testing the vm, the assembler and the program")
	fmt.Println("  simji --test\t\tRun the test units for the assembler and vm")
	fmt.Println("  simji --benchmark\tComputes the number of cycles/second of the vm")
	fmt.Println()
}

func missingFileMessage() {
	log.Error().Msg("Missing input file\n")
	fmt.Println("Type: simji --help to display the program usage")
}

var (
	runBinary   = flag.Bool("binary", false, "runs a binary file instead of a program file")
	showRegs    = flag.Bool("show-regs", false, "show regs on each step")
	showMem     = flag.Bool("show-mem", false, "show memory on each step")
	showDebug   = flag.Bool("debug", false, "show each instruction in the terminal")
	launchGUI   = flag.Bool("gui", false, "start the gui application")
	nbBMRuns    = flag.Int("benchmark", 0, "evalue les performances du simulateur")
	assemble    = flag.Bool("assemble", false, "assemble the targeted program")
	outputFile  = flag.String("output", "", "filename of the output file")
	disassemble = flag.Bool("disassemble", false, "disassembly of a binary file")
	fullHelp    = flag.Bool("help", false, "print the full usage of the program")
)

func init() {
	flag.Usage = usage
	flag.BoolVar(runBinary, "b", false, "alias for --binary")
	flag.BoolVar(showRegs, "r", false, "alias for --show-regs")
	flag.BoolVar(showMem, "m", false, "alias for --show-mem")
	flag.BoolVar(showDebug, "d", false, "alias for --debug")
	flag.BoolVar(launchGUI, "g", false, "alias for --gui")
	flag.IntVar(nbBMRuns, "bm", 0, "alias for --benchmark")
	flag.BoolVar(assemble, "a", false, "alias for --assemble")
	flag.StringVar(outputFile, "o", "", "alias for --output")
	flag.BoolVar(disassemble, "da", false, "alias for --disassemble")

	helloWorld()
}
