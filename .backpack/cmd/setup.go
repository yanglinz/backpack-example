package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/development"
	"github.com/yanglinz/backpack/docker"
	"github.com/yanglinz/backpack/github"
	"github.com/yanglinz/backpack/google"
	"github.com/yanglinz/backpack/internal"
	"github.com/yanglinz/backpack/terraform"
)

func setupSecrets(backpack internal.Context) {
	envDir := filepath.Join(backpack.Root, "etc")
	os.Mkdir(envDir, 0777)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "🏁 Setup project",
	Long:  "🏁 Setup project",
	Run: func(cmd *cobra.Command, args []string) {
		setupFiles, _ := cmd.Flags().GetBool("files")
		setupResources, _ := cmd.Flags().GetBool("resources")
		backpack := internal.ParseContext(cmd)

		setupSecrets(backpack)

		if setupFiles {
			development.SetupTaskfileBin(backpack)
			development.SetupTaskfile(backpack)
			github.CreateWorkflows(backpack)
			docker.CreateComposeConfig(backpack)
			terraform.CreateConfig(backpack)

			err := terraform.ValidateBackend(backpack)
			if err != nil {
				panic(err)
			}
		}
		if setupResources {
			google.BootstrapSecrets(backpack)
		}
	},
}

func init() {
	setupCmd.Flags().Bool("files", true, "setup project files")
	setupCmd.Flags().Bool("resources", false, "setup remote resources")
	rootCmd.AddCommand(setupCmd)
}
