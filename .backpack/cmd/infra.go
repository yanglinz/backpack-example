package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/internal"
)

var terraformPlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "Thin wrapper around terraform plan",
	Long:  "Thin wrapper around terraform plan",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		fmt.Println(backpack)
	},
}

var terraformApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Thin wrapper around terraform apply",
	Long:  "Thin wrapper around terraform apply",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		fmt.Println(backpack)
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

	rootCmd.AddCommand(terraformCmd)
}
