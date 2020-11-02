package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "🔧 Output debug info",
	Long:  "🔧 Output debug info",
	Run: func(cmd *cobra.Command, args []string) {
		appContext := application.ParseContext(cmd)
		data, err := json.MarshalIndent(appContext, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", data)
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
