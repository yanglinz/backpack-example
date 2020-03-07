package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/development"
	"github.com/yanglinz/backpack/internal"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "🐳 Run development server",
	Long:  "🐳 Run development server",
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		prod, _ := cmd.Flags().GetBool("prod")

		if prod {
			os.Setenv("COMPOSE_FILE", "docker-compose-prod.yml")
			command := "docker-compose up"
			shell := internal.GetCommand(command)
			shell.Dir = backpack.Root

			err := shell.Run()
			if err != nil {
				panic(err)
			}
		} else {
			development.RunTaskfile(backpack)
		}
	},
}

func init() {
	runCmd.Flags().Bool("prod", false, "run the pseudo-production image")
	rootCmd.AddCommand(runCmd)
}
