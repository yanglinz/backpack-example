package execution

import (
	"os"
	"os/exec"
	"strings"
)

// GetCommand returns a execution command
func GetCommand(command string) *exec.Cmd {
	commandList := strings.Split(command, " ")
	first := commandList[0]
	rest := commandList[1:]
	cmd := exec.Command(first, rest...)
	cwd, _ := os.Getwd()
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd
}
