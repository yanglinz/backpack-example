package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/internal"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "🔧 Output debug info",
	Long:  "🔧 Output debug info",
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		data, err := json.MarshalIndent(backpack, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", data)
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
