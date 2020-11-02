package development

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yanglinz/backpack/application"
	"github.com/yanglinz/backpack/internal"
)

// CreateCertificates generate self-signed certificates
func CreateCertificates(appContext application.Context) {
	certsDir := filepath.Join(appContext.Root, "etc/certs")
	os.MkdirAll(certsDir, 0777)

	localDomain := appContext.Name + ".localhost"
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
