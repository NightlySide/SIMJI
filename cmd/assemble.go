package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime/pprof"
	"simji/pkg/assembler"

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
				log.Error().Msgf("cmd/assemble - %s", err.Error())
			}
			_ = pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}

		initLogger()
		if assembleDebug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		lines, _ := assembler.ProgramFileToStringArray(args[0])
		numInstructions := assembler.StringLinesToInstructions(lines)
		prog := assembler.ComputeHexInstructions(numInstructions)

		// no output file specified -> print in console
		if assembleOutputPath == "" {
			log.Info().Msg("No output file specified. Printing binary to console.")
			assembler.PrintProgram(prog)
		} else {
			// save to file
			log.Info().Msgf("Exporting binary data to file: %s", assembleOutputPath)
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
