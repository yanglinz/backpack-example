package terraform

import (
	"errors"
	"path/filepath"

	"github.com/yanglinz/backpack/internal"
)

// CreateConfig generates the terraform config
func CreateConfig(backpack internal.Context) {
	// Copy secret variable file
	secretsPath := filepath.Join(backpack.Root, "terraform/secrets.tfvars")
	if !internal.Exists(secretsPath) {
		sourcePath := filepath.Join(backpack.Root, ".backpack/terraform/root/secrets.tfvars")
		internal.CopyFile(sourcePath, secretsPath)
	}
}

// ValidateBackend checks whether backend.tf has been setup
func ValidateBackend(backpack internal.Context) error {
	backend := filepath.Join(backpack.Root, "terraform/backend.tf")
	if internal.Exists(backend) {
		return nil
	}

	return errors.New("Missing terraform/backend.tf")
}
