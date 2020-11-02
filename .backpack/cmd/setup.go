package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/development"
	"github.com/yanglinz/backpack/docker"
	"github.com/yanglinz/backpack/github"
	"github.com/yanglinz/backpack/google"
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
			github.CreateWorkflows(appContext)
			docker.CreateComposeConfig(appContext)
			terraform.CreateConfig(appContext)
		}
		if setupResources {
			google.BootstrapSecrets(appContext)
		}
	},
}

func init() {
	setupCmd.Flags().Bool("files", true, "setup project files")
	setupCmd.Flags().Bool("resources", false, "setup remote resources")
	rootCmd.AddCommand(setupCmd)
}
