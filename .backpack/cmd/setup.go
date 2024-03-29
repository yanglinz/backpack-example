package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/delivery"
	"github.com/yanglinz/backpack/development"
	"github.com/yanglinz/backpack/docker"
	"github.com/yanglinz/backpack/secrets"
	"github.com/yanglinz/backpack/terraform"
)

func setupSecrets(appContext application.Context) {
	envDir := filepath.Join(appContext.Root, "etc")
	os.Mkdir(envDir, 0777)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "🏁 Setup project",
	Long:  "🏁 Setup project",
	Run: func(cmd *cobra.Command, args []string) {
		setupFiles, _ := cmd.Flags().GetBool("files")
		setupResources, _ := cmd.Flags().GetBool("resources")
		appContext := application.ParseContext(cmd)

		setupSecrets(appContext)

		if setupFiles {
			development.SetupTaskfileBin(appContext)
			development.SetupTaskfile(appContext)
			delivery.CreateWorkflows(appContext)
			docker.CreateComposeConfig(appContext)
			terraform.CreateConfig(appContext)
		}
		if setupResources {
			secrets.BootstrapSecrets(appContext)
		}
	},
}

func init() {
	setupCmd.Flags().Bool("files", true, "setup project files")
	setupCmd.Flags().Bool("resources", false, "setup remote resources")
	rootCmd.AddCommand(setupCmd)
}
