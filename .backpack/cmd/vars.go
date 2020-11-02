package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/secrets"
)

var varsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Output current list of variables",
	Long:  "Output current list of variables",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		appContext := application.ParseContext(cmd)
		secret := secrets.GetSecrets(appContext, env)

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
		appContext := application.ParseContext(cmd)

		envFile := filepath.Join(appContext.Root, file)
		envData, err := ioutil.ReadFile(envFile)
		if err != nil {
			panic(err)
		}

		secrets.UpdateSecrets(appContext, secrets.UpdateSecretRequest{
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
