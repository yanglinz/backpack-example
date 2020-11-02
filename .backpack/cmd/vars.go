package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/google"
)

var varsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Output current list of variables",
	Long:  "Output current list of variables",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		backpack := application.ParseContext(cmd)
		secret := google.GetSecrets(backpack, env)

		fmt.Println(string(secret))
	},
}

var varsPutCmd = &cobra.Command{
	Use:   "put",
	Short: "Put variables from local file to secrets manager",
	Long:  "Put variables from local file to secrets manager",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		file, _ := cmd.Flags().GetString("file")
		backpack := application.ParseContext(cmd)

		envFile := filepath.Join(backpack.Root, file)
		envData, err := ioutil.ReadFile(envFile)
		if err != nil {
			panic(err)
		}

		google.UpdateSecrets(backpack, google.UpdateSecretRequest{
			Env:   env,
			Value: string(envData),
		})
	},
}

var varsCmd = &cobra.Command{
	Use:   "vars",
	Short: "Configure environmental variables and secrets",
	Long:  "Configure environmental variables and secrets",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	varsGetCmd.Flags().String("env", application.EnvDevelopment, "environment")
	varsCmd.AddCommand(varsGetCmd)
	varsPutCmd.Flags().String("env", application.EnvDevelopment, "environment")
	varsPutCmd.Flags().String("file", ".", "file")
	varsCmd.AddCommand(varsPutCmd)

	rootCmd.AddCommand(varsCmd)
}
