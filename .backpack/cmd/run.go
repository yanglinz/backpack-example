package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/development"
	"github.com/yanglinz/backpack/io/execution"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "🐳 Run development server",
	Long:  "🐳 Run development server",
	Run: func(cmd *cobra.Command, args []string) {
		appContext := application.ParseContext(cmd)
		prod, _ := cmd.Flags().GetBool("prod")

		if prod {
			os.Setenv("COMPOSE_FILE", "docker-compose-prod.yml")
			command := "docker-compose up"
			shell := execution.GetCommand(command)
			shell.Dir = appContext.Root

			err := shell.Run()
			if err != nil {
				panic(err)
			}
		} else {
			development.RunTaskfile(appContext)
		}
	},
}

func init() {
	runCmd.Flags().Bool("prod", false, "run the pseudo-production image")
	rootCmd.AddCommand(runCmd)
}
