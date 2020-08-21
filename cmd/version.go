package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"resource-optim/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "resource-optim version",
	Long:  `resource-optim version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	value := version.GetVersion()
	log.Println(value)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
