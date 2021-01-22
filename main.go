package main

import (
	"flag"
	"fmt"
	"os"

	"simji/internal/assembler"
	"simji/internal/gui"
	"simji/internal/log"
	"simji/internal/vm"

	"github.com/markbates/pkger"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if *showDebug {
		log.GetLogger().SetLevel(log.DEBUG)
	} else {
		log.GetLogger().SetLevel(log.INFO)
	}

	if *fullHelp {
		fullUsage()
	} else if *launchGUI {
		// If we try to launch the GUI
		staticFiles := pkger.Dir("/internal/static")
		log.GetLogger().Info("Launching gui...")
		// Include static files for packaging
		gui.ShowGUI(staticFiles)
	} else if *assemble {
		// If we just want to assemble the file
		// If there is not enough arguments
		if len(args) < 1 {
			missingFileMessage()
			os.Exit(1)
		}

		lines, _ := assembler.ProgramFileToStringArray(args[0])
		numInstructions := assembler.StringLinesToInstructions(lines)
		prog := assembler.ComputeHexInstructions(numInstructions)

		// no output file specified -> print in console
		if *outputFile == "" {
			log.GetLogger().Info("No output file specified. Printing binary to console.")
			assembler.PrintProgram(prog)
		} else {
			// save to file
			log.GetLogger().Info(fmt.Sprintf("Exporting binary data to file: %s", *outputFile))
			assembler.ExportBinaryToFile(prog, *outputFile)
		}

	} else if *disassemble {
		// If we just want to disassemble the file
		// If there is not enough arguments
		if len(args) < 1 {
			missingFileMessage()
			os.Exit(1)
		}

		prog := vm.LoadProgFromFile(args[0])
		desProg := vm.Disassemble(prog)

		// no output file specified -> print in console
		if *outputFile == "" {
			log.GetLogger().Info("No output file specified. Printing binary to console.")
			for _, line := range desProg {
				fmt.Println(line)
			}
		} else {
			// save to file
			log.GetLogger().Info(fmt.Sprintf("Exporting disassembled data to file: %s", *outputFile))
			assembler.ExportProgramToFile(desProg, *outputFile)
		}
	} else if *runBinary {
		// run the program from a bin file
		// If there is not enough arguments
		if len(args) < 1 {
			missingFileMessage()
			os.Exit(1)
		}

		// else we load the program
		prog := vm.LoadProgFromFile(args[0])

		vm := vm.NewVM(32, 1000)
		vm.LoadProg(prog)
		vm.Run(*showRegs, *showMem, *showDebug)

	} else {
		// If there is not enough arguments
		if len(args) < 1 {
			missingFileMessage()
			os.Exit(1)
		}

		if *nbBMRuns == 0 {
			// Else we load the program
			lines, _ := assembler.ProgramFileToStringArray(args[0])
			numInstructions := assembler.StringLinesToInstructions(lines)
			prog := assembler.ComputeHexInstructions(numInstructions)

			log.GetLogger().Title(log.DEBUG, "Launching VM")
			log.GetLogger().Debug("-- Creating VM with: 32 registers\n")
			vm := vm.NewVM(32, 1000)

			vm.LoadProg(prog)
			vm.Run(*showRegs, *showMem, *showDebug)
		} else {
			// Else we load the program
			lines, _ := assembler.ProgramFileToStringArray(args[0])
			numInstructions := assembler.StringLinesToInstructions(lines)
			prog := assembler.ComputeHexInstructions(numInstructions)

			bm := vm.StartBenchmark(prog, *nbBMRuns)
			bm.PrintResults()
		}
	}
}
