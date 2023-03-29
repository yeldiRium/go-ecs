package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootVersionFlag bool

func init() {
	RootCommand.Flags().BoolVarP(&rootVersionFlag, "version", "v", false, "prints the version")
}

var RootCommand = &cobra.Command{
	Use:   "go-ecs",
	Short: "A go Entity Component System.",
	Long:  "Contains hopefully various little games for fun and education.",
	Run: func(command *cobra.Command, args []string) {
		if rootVersionFlag {
			versionCommand.Run(command, args)
			return
		}

		fmt.Println(command.UsageString())
	},
}
