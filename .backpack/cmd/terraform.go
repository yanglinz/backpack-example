package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/io/execution"
)

var terraformPlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "Thin wrapper around terraform plan",
	Long:  "Thin wrapper around terraform plan",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		appContext := application.ParseContext(cmd)
		terraformDir := filepath.Join(appContext.Root, "terraform")

		// Run terraform init
		shell := execution.GetCommand("terraform init")
		shell.Dir = terraformDir
		err := shell.Run()
		if err != nil {
			panic(err)
		}

		// Run terraform plan
		shell = execution.GetCommand("terraform plan -var-file=secrets.tfvars")
		shell.Dir = terraformDir
		err = shell.Run()
		if err != nil {
			panic(err)
		}
	},
}

var terraformApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Thin wrapper around terraform apply",
	Long:  "Thin wrapper around terraform apply",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		appContext := application.ParseContext(cmd)
		terraformDir := filepath.Join(appContext.Root, "terraform")

		// Run terraform apply
		shell := execution.GetCommand("terraform apply -var-file=secrets.tfvars")
		shell.Dir = terraformDir
		err := shell.Run()
		if err != nil {
			panic(err)
		}

		// Get output
		// TODO: Create Ansible inventory from this output
		shell = execution.GetCommand("terraform output")
		shell.Dir = terraformDir
		err = shell.Run()
		if err != nil {
			panic(err)
		}
	},
}

var terraformDestroyCmd = &cobra.Command{
	Use:   "__unsafe_destroy__",
	Short: "Thin wrapper around terraform destroy",
	Long:  "Thin wrapper around terraform destory",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		appContext := application.ParseContext(cmd)
		terraformDir := filepath.Join(appContext.Root, "terraform")

		// Run terraform apply
		shell := execution.GetCommand("terraform destroy -var-file=secrets.tfvars")
		shell.Dir = terraformDir
		err := shell.Run()
		if err != nil {
			panic(err)
		}
	},
}

var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "🚀 Thin wrapper around terraform",
	Long:  "🚀 Thin wrapper around terraform",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	terraformCmd.AddCommand(terraformPlanCmd)
	terraformCmd.AddCommand(terraformApplyCmd)
	terraformCmd.AddCommand(terraformDestroyCmd)

	rootCmd.AddCommand(terraformCmd)
}
