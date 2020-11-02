package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/io/execution"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "💻 Run docker shell",
	Long:  "💻 Run docker shell",
	Run: func(cmd *cobra.Command, args []string) {
		appContext := application.ParseContext(cmd)
		serviceName := appContext.Projects[0].Name + "_server"
		shell := execution.GetCommand(
			"docker-compose run " + serviceName + " .backpack/runtime/entry-shell.sh",
		)
		err := shell.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
