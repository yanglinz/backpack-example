package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/development"
	"github.com/yanglinz/backpack/io/execution"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "🛠  Build the docker images",
	Long:  "🛠  Build the docker images",
	Run: func(cmd *cobra.Command, args []string) {
		appContext := application.ParseContext(cmd)
		prod, _ := cmd.Flags().GetBool("prod")

		if prod {
			os.Setenv("COMPOSE_FILE", "docker-compose-prod.yml")
		}
		development.CreateCertificates(appContext)
		command := "docker-compose build"
		shell := execution.GetCommand(command)
		shell.Dir = appContext.Root
		err := shell.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	buildCmd.Flags().Bool("prod", false, "build the pseudo-production image")
	rootCmd.AddCommand(buildCmd)
}
