package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/internal"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "💻 Run docker shell",
	Long:  "💻 Run docker shell",
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		serviceName := backpack.Projects[0].Name + "_server"
		shell := internal.GetCommand(
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
