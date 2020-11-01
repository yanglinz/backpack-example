package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/internal"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "🚏 Stop running processes",
	Long:  "🚏 Stop running processes",
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)

		command := "docker-compose down"
		shell := internal.GetCommand(command)
		shell.Dir = backpack.Root
		err := shell.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
