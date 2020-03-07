package cmd

import (
	"github.com/spf13/cobra"
)

var rootFlag string

var rootCmd = &cobra.Command{
	Use:   "backpack",
	Short: "🎒 CLI to interact with backpack",
	Long:  "🎒 CLI to interact with backpack",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Do nothing
	},
}

// Execute is the main entrypoint for cobra
func Execute() {
	rootCmd.PersistentFlags().StringVar(&rootFlag, "root", "", "path to the root of the project")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
