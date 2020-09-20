package internal

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Exists checks if filename exists
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CopyFile copies a file
func CopyFile(source string, dest string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		return err
	}

	return nil
}

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
