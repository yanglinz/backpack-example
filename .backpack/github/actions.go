package github

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/yanglinz/backpack/internal"
)

// CreateWorkflows generate github actions configs
func CreateWorkflows(backpack internal.Context) {
	workflowDir := filepath.Join(backpack.Root, ".github/workflows")
	os.MkdirAll(workflowDir, 0777)

	sourcePath := filepath.Join(backpack.Root, ".backpack/github/action-workflow.yml")
	targetPath := filepath.Join(backpack.Root, ".github/workflows/main.yml")
	content, err := ioutil.ReadFile(sourcePath)

	workflow := string(content)
	workflow = strings.ReplaceAll(workflow, "${{APP_NAME}}", backpack.Name)
	workflow = strings.ReplaceAll(workflow, "${{RUNTIME_PLATFORM}}", backpack.Runtime)
	workflow = strings.ReplaceAll(workflow, "${{GCP_PROJECT_ID}}", backpack.Google.ProjectID)

	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(targetPath, []byte(workflow), 0644)
	if err != nil {
		panic(err)
	}
}
