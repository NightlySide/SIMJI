package cmd

import (
	"simji/pkg/vm"
	"simji/pkg/log"

	"github.com/spf13/cobra"
)

var vmDebug bool
var nbBMRuns int64

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <program.bin>",
	Args: cobra.ExactArgs(1), // make sure only one argument is passed
	Short: "Runs a set of instructions in a virtual machine",
	Long: `========================== Help: Run command =========================
Runs a set of instructions from a compatible binary file.
Compatible binary files are assembled using the SIMJI assemble
command. Using the 'd' flag, one can see the debugging 
process of the virtual machine's execution. Using the 'b'
flag, one can benchmark the performances of the virtual machine
in number of cycles per second.

Example:
  ./simji run testdata/program.bin
  ./simji run --debug testdata/program.bin
  ./simji run --benchmark 20000 testdata/program.bin`,
	Run: func(cmd *cobra.Command, args []string) {
		if vmDebug {
			log.GetLogger().SetLevel(log.DEBUG)
		}

		if nbBMRuns == 0 {
			prog := vm.LoadProgFromFile(args[0])

			vm := vm.NewVM(32, 1000)
			vm.LoadProg(prog)
			vm.Run(false, false, vmDebug)
		} else {
			prog := vm.LoadProgFromFile(args[0])

			bm := vm.StartBenchmark(prog, int(nbBMRuns))
			bm.PrintResults()
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&vmDebug, "debug", "d", false, "Print debug infos during the execution")
	runCmd.Flags().Int64VarP(&nbBMRuns, "benchmark", "b", 0, "Benchmark the VM by launching a program N times")
}
