package cmd

import (
	"fmt"
	"os"
	"runtime/pprof"
	"simji/pkg/assembler"
	"simji/pkg/log"

	"github.com/spf13/cobra"
)

var assembleOutputPath string
var assembleDebug bool
var assembleProfiling string

// assembleCmd represents the assemble command
var assembleCmd = &cobra.Command{
	Use:   "assemble <filename.asm>",
	Args: cobra.ExactArgs(1), // make sure only one argument is passed
	Short: "Assemble an MIPS-assembly programm",
	Long: `======================= Help: Assemble command =======================
Assemble an MIPS-assembly program into a set of hexadecimal 
instructions that the computer can understand. Defaults print the 
result to the console. Using the '-o' flag, one can export the 
assembled data to a binary file. Using the '-d' flag, one can print
the debug logs of the program assembly.

Example:
  ./simji assemble testdata/syracuse.asm
  ./simji assemble -d testdata/syracuse.asm
  ./simji assemble --output=program.bin testdata/syracuse.asm`,
	Run: func(cmd *cobra.Command, args []string) {
		if assembleProfiling != "" {
			f, err := os.Create(assembleProfiling)
			if err != nil {
				log.GetLogger().Error(err.Error())
			}
			err = pprof.StartCPUProfile(f)
			if err != nil {
				log.GetLogger().Error(err.Error())
			}
			defer pprof.StopCPUProfile()
		}

		if assembleDebug {
			log.GetLogger().SetLevel(log.DEBUG)
		}

		lines, _ := assembler.ProgramFileToStringArray(args[0])
		numInstructions := assembler.StringLinesToInstructions(lines)
		prog := assembler.ComputeHexInstructions(numInstructions)

		// no output file specified -> print in console
		if assembleOutputPath == "" {
			log.GetLogger().Info("No output file specified. Printing binary to console.")
			assembler.PrintProgram(prog)
		} else {
			// save to file
			log.GetLogger().Info(fmt.Sprintf("Exporting binary data to file: %s", assembleOutputPath))
			_ = assembler.ExportBinaryToFile(prog, assembleOutputPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(assembleCmd)

	// Local flags definition
	assembleCmd.Flags().StringVarP(&assembleOutputPath, "output", "o", "", "Exports hex instructions to binary file")
	assembleCmd.Flags().BoolVarP(&assembleDebug, "debug", "d", false, "Print debugging logs of the assembly process")
	assembleCmd.Flags().StringVarP(&assembleProfiling, "cpuprofile", "p", "", "Exports profiling data into the specified file")
}
