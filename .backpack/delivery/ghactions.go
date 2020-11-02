package delivery

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/yanglinz/backpack/application"
)

func copyWorkflow(appContext application.Context, source string, target string) error {
	content, _ := ioutil.ReadFile(source)

	workflow := string(content)
	workflow = strings.ReplaceAll(workflow, "${{APP_NAME}}", appContext.Name)
	workflow = strings.ReplaceAll(workflow, "${{RUNTIME_PLATFORM}}", appContext.Runtime)
	workflow = strings.ReplaceAll(workflow, "${{GCP_PROJECT_ID}}", appContext.Google.ProjectID)

	err := ioutil.WriteFile(target, []byte(workflow), 0644)
	return err
}

// CreateWorkflows generate github actions configs
func CreateWorkflows(appContext application.Context) {
	workflowDir := filepath.Join(appContext.Root, ".github/workflows")
	os.MkdirAll(workflowDir, 0777)

	source := filepath.Join(appContext.Root, ".backpack/delivery/actions/main.yml")
	target := filepath.Join(appContext.Root, ".github/workflows/main.yml")
	err := copyWorkflow(appContext, source, target)
	if err != nil {
		panic(err)
	}

	source = filepath.Join(appContext.Root, ".backpack/delivery/actions/deployment.yml")
	target = filepath.Join(appContext.Root, ".github/workflows/deployment.yml")
	err = copyWorkflow(appContext, source, target)
	if err != nil {
		panic(err)
	}
}
