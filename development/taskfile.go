package development

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yanglinz/backpack/internal"
)

// SetupTaskfileBin generates taskfile binary
func SetupTaskfileBin(backpack internal.Context) {
	binDir := filepath.Join(backpack.Root, "bin")
	binPath := filepath.Join(binDir, "task")
	if internal.Exists(binPath) {
		return
	}

	parts := []string{
		filepath.Join(backpack.Root, ".backpack/bin/install-taskfile"),
		"-b", binDir,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	shell.Dir = backpack.Root
	shell.Run()
}

// SetupTaskfile generates the taskfile config
func SetupTaskfile(backpack internal.Context) {
	target := ".backpack/development/Taskfile.yml"
	symlink := filepath.Join(backpack.Root, "Taskfile.yml")
	os.Remove(symlink)
	err := os.Symlink(target, symlink)
	if err != nil {
		panic(err)
	}
}

// RunTaskfile runs the development server
func RunTaskfile(backpack internal.Context) {
	shell := internal.GetCommand("bin/task -p server ui")
	shell.Dir = backpack.Root
	shell.Run()
}
