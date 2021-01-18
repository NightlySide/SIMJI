package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func helloWorld() {
	defaultColor := color.New(color.FgWhite)
	titleColor := color.New(color.FgHiCyan).Add(color.Underline).Add(color.Bold)
	detailColor := color.New(color.FgBlue)

	defaultColor.Print("========= ")
	titleColor.Printf("SIMJI : Simulateur de Jeu d'Instructions ")
	defaultColor.Print("========= \n")
	defaultColor.Print("-- Conçut par ")
	detailColor.Print("Alexandre FROEHLICH \n")
	defaultColor.Println("-- Dans le cadre de l'U.V. 4.5-Architectures numériques")
	defaultColor.Print("-- Contact : ")
	detailColor.Println("nightlyside@gmail.com")
	defaultColor.Print("-- Site web : ")
	detailColor.Println("https://nightlyside.github.io")
	defaultColor.Println("============================================================")
	fmt.Print("\n")
}

func usage() {
	fmt.Println("usage: simji [--help | -h] [--gui | -g] [--debug | -d]")
	fmt.Println("             [--show-regs | -r] [--show-mem | -m]")
	fmt.Println("             \"filename\"")
	fmt.Println()

	fmt.Println("Some common commands in different situations:")
	fmt.Println()

	fmt.Println("start and use the program")
	fmt.Println("  simji myprogram.asm\tLaunch the program \"myprogram.asm\" in CLI mode")
	fmt.Println("  simji --gui\t\tLaunch the program with a Graphical UI")
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

    os.Exit(2)
}

func missingFileMessage() {
	fmt.Printf("[%s] %s: Missing input file\n", color.RedString("-"), color.RedString("ERROR"))
	fmt.Println("Type: simji --help to display the program usage")
}

var showRegs = flag.Bool("show-regs", false, "show regs on each step")
var showMem = flag.Bool("show-mem", false, "show memory on each step")
var showDebug = flag.Bool("debug", false, "show each instruction in the terminal")
var launchGUI = flag.Bool("gui", false, "start the gui application")


func init() {
	flag.Usage = usage
	flag.BoolVar(showRegs, "r", false, "alias for --show-regs")
	flag.BoolVar(showMem, "m", false, "alias for --show-mem")
	flag.BoolVar(showDebug, "d", false, "alias for --debug")
	flag.BoolVar(launchGUI, "g", false, "alias for --gui")

	helloWorld()
}