package development

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/io/execution"
	"github.com/yanglinz/backpack/io/filesystem"
)

// SetupTaskfileBin generates taskfile binary
func SetupTaskfileBin(appContext application.Context) {
	binDir := filepath.Join(appContext.Root, "bin")
	binPath := filepath.Join(binDir, "task")
	if filesystem.Exists(binPath) {
		return
	}

	parts := []string{
		filepath.Join(appContext.Root, ".backpack/bin/install-taskfile.sh"),
		"-b", binDir,
	}
	command := strings.Join(parts, " ")
	shell := execution.GetCommand(command)
	shell.Dir = appContext.Root
	shell.Run()
}

// SetupTaskfile generates the taskfile config
func SetupTaskfile(appContext application.Context) {
	target := ".backpack/development/Taskfile.yml"
	symlink := filepath.Join(appContext.Root, "Taskfile.yml")
	os.Remove(symlink)
	err := os.Symlink(target, symlink)
	if err != nil {
		panic(err)
	}
}

// RunTaskfile runs the development server
func RunTaskfile(appContext application.Context) {
	shell := execution.GetCommand("bin/task -p server ui")
	shell.Dir = appContext.Root
	shell.Run()
}
