package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"simji/pkg/assembler"
	"simji/pkg/vm"

	"github.com/spf13/cobra"
)

var disassembleOutputPath string

// disassembleCmd represents the assemble command
var disassembleCmd = &cobra.Command{
	Use:   "disassemble <program.bin>",
	Args: cobra.ExactArgs(1), // make sure only one argument is passed
	Short: "Disassemble a compatible binary programm",
	Long: `===================== Help: Disassemble command ======================
Disassemble a binary program into human-readable instructions.
The labels cannot be printed back as the label is replaced by 
its address during assembly. Using the 'o' flag, one can export 
the disassembled data into an MIPS-assembly program file.

Example:
  ./simji disassemble testdata/program.bin
  ./simji disassemble --output=program.asm testdata/program.bin`,
	Run: func(cmd *cobra.Command, args []string) {
		prog := vm.LoadProgFromFile(args[0])
		desProg := vm.Disassemble(prog)

		initLogger()

		// no output file specified -> print in console
		if disassembleOutputPath == "" {
			log.Info().Msg("No output file specified. Printing binary to console.")
			for _, line := range desProg {
				fmt.Println(line)
			}
		} else {
			// save to file
			log.Info().Msgf("Exporting disassembled data to file: %s", disassembleOutputPath)
			_ = assembler.ExportProgramToFile(desProg, disassembleOutputPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(disassembleCmd)

	// Local flags definition
	disassembleCmd.Flags().StringVarP(&disassembleOutputPath, "output", "o", "", "Exports asm instructions to program file")
}
