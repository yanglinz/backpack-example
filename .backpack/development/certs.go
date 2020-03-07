package development

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yanglinz/backpack/internal"
)

// CreateCertificates generate self-signed certificates
func CreateCertificates(backpack internal.Context) {
	certsDir := filepath.Join(backpack.Root, "etc/certs")
	os.MkdirAll(certsDir, 0777)

	localDomain := backpack.Name + ".localhost"
	commandParts := []string{
		"mkcert",
		"-cert-file app-localhost.pem",
		"-key-file app-localhost-key.pem",
		localDomain,
	}
	command := strings.Join(commandParts, " ")
	shell := internal.GetCommand(command)
	shell.Dir = certsDir
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}
