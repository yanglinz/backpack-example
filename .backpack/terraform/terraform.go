package terraform

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/stoewer/go-strcase"
	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/io/filesystem"
)

func createSecretConfig(appContext application.Context) {
	secretsPath := filepath.Join(appContext.Root, "terraform/secrets.tfvars")
	if !filesystem.Exists(secretsPath) {
		sourcePath := filepath.Join(appContext.Root, ".backpack/terraform/root/secrets.tfvars")
		filesystem.CopyFile(sourcePath, secretsPath)
	}
}

func createMainConfig(appContext application.Context) {
	config := make(map[string]interface{})

	// Define outermost module
	modules := []interface{}{}

	// Define main module
	appModule := make(map[string]interface{})
	appModule["app_name"] = appContext.Name

	module := make(map[string]interface{})
	module["source"] = "../.backpack/terraform/core-web-digitalocean"
	module["app_context"] = appModule

	moduleContainer := make(map[string]interface{})
	moduleName := strcase.SnakeCase(appContext.Name + "_web")
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
	configPath := filepath.Join(appContext.Root, "terraform/main.tf.json")
	err := ioutil.WriteFile(configPath, []byte(content), 0666)
	if err != nil {
		panic(err)
	}
}

func createMetaConfig(appContext application.Context) {
	sourcePath := filepath.Join(appContext.Root, ".backpack/terraform/root/meta.tf")
	targetPath := filepath.Join(appContext.Root, "terraform/meta.tf")
	content, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		panic(err)
	}

	output := string(content)
	output = strings.ReplaceAll(output, "{{BACKPACK_DEFAULT_ORG}}", "yanglin")
	output = strings.ReplaceAll(output, "{{BACKPACK_WORKSPACE}}", appContext.Name)
	err = ioutil.WriteFile(targetPath, []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}

// CreateConfig generates the terraform config
func CreateConfig(appContext application.Context) {
	terraformDir := filepath.Join(appContext.Root, "terraform")
	os.MkdirAll(terraformDir, 0777)

	createSecretConfig(appContext)
	createMainConfig(appContext)
	createMetaConfig(appContext)
}
