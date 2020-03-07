package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/google"
	"github.com/yanglinz/backpack/heroku"
	"github.com/yanglinz/backpack/internal"
	"github.com/yanglinz/backpack/symbols"
)

var varsContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Output current context info",
	Long:  "Output current context info",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		google.ListSecrets(backpack)
	},
}

var varsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Sync variables from cloud to local file",
	Long:  "Sync variables from cloud to local file",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		backpack := internal.ParseContext(cmd)

		secretKey := "BERGLAS_APP_DEV_JSON"
		if env == symbols.EnvProduction {
			secretKey = "BERGLAS_APP_JSON"
		}
		secret := google.GetSecret(backpack, secretKey)

		var envJSON map[string]string
		json.Unmarshal([]byte(secret), &envJSON)
		formattedJSON, err := json.MarshalIndent(envJSON, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(formattedJSON))
	},
}

var varsPutCmd = &cobra.Command{
	Use:   "put",
	Short: "Sync variables from local file to cloud",
	Long:  "Sync variables from local file to cloud",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		backpack := internal.ParseContext(cmd)

		secretKey := "BERGLAS_APP_DEV_JSON"
		envFile := filepath.Join(backpack.Root, "etc/development.json")
		if env == symbols.EnvProduction {
			secretKey = "BERGLAS_APP_JSON"
			envFile = filepath.Join(backpack.Root, "etc/production.json")
		}

		envData, err := ioutil.ReadFile(envFile)
		if err != nil {
			panic(err)
		}

		var envJSON map[string]string
		json.Unmarshal(envData, &envJSON)
		formattedJSON, err := json.Marshal(envJSON)
		if err != nil {
			panic(err)
		}

		if backpack.Runtime == symbols.RuntimeHeroku {
			for k, v := range envJSON {
				heroku.PutSecret(heroku.PutSecretRequest{
					App:   backpack.Heroku.AppName,
					Name:  k,
					Value: v,
				})
			}
		}

		if backpack.Runtime == symbols.RuntimeCloudrun {
			google.UpdateSecret(backpack, google.UpdateSecretRequest{
				Name:  secretKey,
				Value: string(formattedJSON),
			})
		}
	},
}

var varsInternalNewCmd = &cobra.Command{
	Use:   "_new",
	Short: "Create a new variable by name",
	Long:  "Create a new variable by name",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		google.CreateSecret(backpack, google.CreateSecretRequest{
			Name:  args[0],
			Value: args[1],
		})
	},
}

var varsInternalUpdateCmd = &cobra.Command{
	Use:   "_update",
	Short: "Update a variable by name",
	Long:  "Update a variable by name",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		google.UpdateSecret(backpack, google.UpdateSecretRequest{
			Name:  args[0],
			Value: args[1],
		})
	},
}

var varsInternalDeleteCmd = &cobra.Command{
	Use:   "_delete",
	Short: "Delete a variable by name",
	Long:  "Update a variable by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		google.DeleteSecret(backpack, args[0])
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
	varsCmd.AddCommand(varsContextCmd)

	varsGetCmd.Flags().String("env", symbols.EnvDevelopment, "environment")
	varsCmd.AddCommand(varsGetCmd)
	varsPutCmd.Flags().String("env", symbols.EnvDevelopment, "environment")
	varsCmd.AddCommand(varsPutCmd)

	varsCmd.AddCommand(varsInternalNewCmd)
	varsCmd.AddCommand(varsInternalUpdateCmd)
	varsCmd.AddCommand(varsInternalDeleteCmd)

	rootCmd.AddCommand(varsCmd)
}
