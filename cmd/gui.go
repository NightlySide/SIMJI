package cmd

import (
	"github.com/markbates/pkger"
	"github.com/rs/zerolog/log"
	"github.com/Nightlyside/simji/pkg/gui"

	"github.com/spf13/cobra"
)

// guiCmd represents the gui command
var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Launch the graphical interface",
	Long: `========================== Help: GUI command =========================
Launches the graphical interface to edit, assemble and run
MIPS-assembly files. On other tabs one can see the registers
state and the memory values.

Example:
  ./simji gui`,
	Run: func(cmd *cobra.Command, args []string) {
		initLogger()
		staticFiles := pkger.Dir("/pkg/static")
		log.Info().Msg("Launching gui...")
		// Include static files for packaging
		gui.ShowGUI(staticFiles)
	},
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
