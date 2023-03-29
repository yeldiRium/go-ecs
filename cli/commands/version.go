package commands

import (
	"fmt"

	"github.com/yeldir/go-ecs/version"

	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Long:  "Prints the version.",
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("Revision " + version.GitVersion)
		fmt.Println()
	},
}
