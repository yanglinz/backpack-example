package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/io/execution"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "🚏 Stop running processes",
	Long:  "🚏 Stop running processes",
	Run: func(cmd *cobra.Command, args []string) {
		appContext := application.ParseContext(cmd)

		command := "docker-compose down"
		shell := execution.GetCommand(command)
		shell.Dir = appContext.Root
		err := shell.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
