package github

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/yanglinz/backpack/internal"
)

func copyWorkflow(backpack internal.Context, source string, target string) error {
	content, _ := ioutil.ReadFile(source)

	workflow := string(content)
	workflow = strings.ReplaceAll(workflow, "${{APP_NAME}}", backpack.Name)
	workflow = strings.ReplaceAll(workflow, "${{RUNTIME_PLATFORM}}", backpack.Runtime)
	workflow = strings.ReplaceAll(workflow, "${{GCP_PROJECT_ID}}", backpack.Google.ProjectID)

	err := ioutil.WriteFile(target, []byte(workflow), 0644)
	return err
}

// CreateWorkflows generate github actions configs
func CreateWorkflows(backpack internal.Context) {
	workflowDir := filepath.Join(backpack.Root, ".github/workflows")
	os.MkdirAll(workflowDir, 0777)

	source := filepath.Join(backpack.Root, ".backpack/github/actions/main.yml")
	target := filepath.Join(backpack.Root, ".github/workflows/main.yml")
	err := copyWorkflow(backpack, source, target)
	if err != nil {
		panic(err)
	}

	source = filepath.Join(backpack.Root, ".backpack/github/actions/deployment.yml")
	target = filepath.Join(backpack.Root, ".github/workflows/deployment.yml")
	err = copyWorkflow(backpack, source, target)
	if err != nil {
		panic(err)
	}

	source = filepath.Join(backpack.Root, ".backpack/github/actions/infrastructure.yml")
	target = filepath.Join(backpack.Root, ".github/workflows/infrastructure.yml")
	err = copyWorkflow(backpack, source, target)
	if err != nil {
		panic(err)
	}
}
