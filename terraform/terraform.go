package terraform

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/stoewer/go-strcase"
	"github.com/yanglinz/backpack/internal"
)

func createSecretConfig(backpack internal.Context) {
	secretsPath := filepath.Join(backpack.Root, "terraform/secrets.tfvars")
	if !internal.Exists(secretsPath) {
		sourcePath := filepath.Join(backpack.Root, ".backpack/terraform/root/secrets.tfvars")
		internal.CopyFile(sourcePath, secretsPath)
	}
}

func createMainConfig(backpack internal.Context) {
	config := make(map[string]interface{})

	// Define outermost module
	modules := []interface{}{}

	// Define main module
	appContext := make(map[string]interface{})
	appContext["app_name"] = backpack.Name

	module := make(map[string]interface{})
	module["source"] = "../.backpack/terraform/core-web-digitalocean"
	module["app_context"] = appContext

	moduleContainer := make(map[string]interface{})
	moduleName := strcase.SnakeCase(backpack.Name + "_web")
	moduleContainer[moduleName] = []interface{}{module}
	modules = append(modules, moduleContainer)

	// Define outputs
	outputs := []interface{}{}

	addressOutput := make(map[string]interface{})
	addressOutput["value"] = "${module." + moduleName + ".ip_address}"

	outputContainer := make(map[string]interface{})
	outputContainer["ip_address"] = []interface{}{addressOutput}
	outputs = append(outputs, outputContainer)

	config["module"] = modules
	config["output"] = outputs

	// Write module to file
	content, _ := json.MarshalIndent(config, "", "  ")
	configPath := filepath.Join(backpack.Root, "terraform/main.tf.json")
	err := ioutil.WriteFile(configPath, []byte(content), 0666)
	if err != nil {
		panic(err)
	}
}

func createMetaConfig(backpack internal.Context) {
	sourcePath := filepath.Join(backpack.Root, ".backpack/terraform/root/meta.tf")
	targetPath := filepath.Join(backpack.Root, "terraform/meta.tf")
	content, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		panic(err)
	}

	output := string(content)
	output = strings.ReplaceAll(output, "{{BACKPACK_DEFAULT_ORG}}", "yanglin")
	output = strings.ReplaceAll(output, "{{BACKPACK_WORKSPACE}}", backpack.Name)
	err = ioutil.WriteFile(targetPath, []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}

// CreateConfig generates the terraform config
func CreateConfig(backpack internal.Context) {
	terraformDir := filepath.Join(backpack.Root, "terraform")
	os.MkdirAll(terraformDir, 0777)

	createSecretConfig(backpack)
	createMainConfig(backpack)
	createMetaConfig(backpack)
}
