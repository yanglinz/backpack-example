package google

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/berglas/pkg/berglas"
	"github.com/yanglinz/backpack/internal"
)

func bucketExists(bucketName string) bool {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	bucket := client.Bucket(bucketName)
	bucketAttrs, err := bucket.Attrs(ctx)

	if bucketAttrs != nil {
		return true
	}
	return false
}

func bootstrapServiceAccount(backpack internal.Context) {
	// Create service account to fetch secrets
	serviceAccountName := "berglas-" + backpack.Name
	parts := []string{
		"gcloud iam service-accounts create",
		serviceAccountName,
		"--project", backpack.Google.ProjectID,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	err := shell.Run()
	if err != nil {
		panic(err)
	}

	// Grant service account access to secrets
	serviceAccountEmail := fmt.Sprintf(
		"serviceAccount:%s@%s.iam.gserviceaccount.com",
		serviceAccountName,
		backpack.Google.ProjectID,
	)
	bucketName := "berglas-" + backpack.Name
	grantKey := bucketName + "/BERGLAS_APP_JSON"
	parts = []string{
		"berglas grant", grantKey,
		"--member", serviceAccountEmail,
	}
	command = strings.Join(parts, " ")
	shell = internal.GetCommand(command)
	err = shell.Run()
	if err != nil {
		panic(err)
	}
}

// BootstrapSecrets for berglas
func BootstrapSecrets(backpack internal.Context) {
	ctx := context.Background()
	bucketName := "berglas-" + backpack.Name
	exists := bucketExists(bucketName)
	if exists {
		return
	}

	// Run the berglas bootstrap command
	err := berglas.Bootstrap(ctx, &berglas.StorageBootstrapRequest{
		ProjectID: backpack.Google.ProjectID,
		Bucket:    bucketName,
	})
	if err != nil {
		panic(err)
	}

	// Bootstrap the initial secrets
	CreateSecret(backpack, CreateSecretRequest{
		Name:  "BERGLAS_APP_JSON",
		Value: "{}",
	})
	CreateSecret(backpack, CreateSecretRequest{
		Name:  "BERGLAS_APP_DEV_JSON",
		Value: "{}",
	})

	bootstrapServiceAccount(backpack)
}

// ListSecrets outputs a list of secrets
func ListSecrets(backpack internal.Context) {
	bucketName := "berglas-" + backpack.Name
	shell := internal.GetCommand("berglas list " + bucketName)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}

// CreateSecretRequest params
type CreateSecretRequest struct {
	Name  string
	Value string
}

// CreateSecret creates or updates a secret
func CreateSecret(backpack internal.Context, req CreateSecretRequest) {
	bucketName := "berglas-" + backpack.Name
	bucketPath := bucketName + "/" + req.Name
	encryptionKey := "projects/" + backpack.Google.ProjectID + "/locations/global/keyRings/berglas/cryptoKeys/berglas-key"
	parts := []string{
		"berglas create", bucketPath, req.Value,
		"--key", encryptionKey,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}

// UpdateSecretRequest params
type UpdateSecretRequest struct {
	Name  string
	Value string
}

// UpdateSecret creates or updates a secret
func UpdateSecret(backpack internal.Context, req UpdateSecretRequest) {
	bucketName := "berglas-" + backpack.Name
	bucketPath := bucketName + "/" + req.Name
	encryptionKey := "projects/" + backpack.Google.ProjectID + "/locations/global/keyRings/berglas/cryptoKeys/berglas-key"
	parts := []string{
		"berglas update", bucketPath, req.Value,
		"--key", encryptionKey,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}

// GetSecret list a single secret
func GetSecret(backpack internal.Context, name string) string {
	bucketName := "berglas-" + backpack.Name
	bucketPath := bucketName + "/" + name
	parts := []string{"berglas access", bucketPath}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	shell.Stdout = nil
	out, err := shell.Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

// DeleteSecret removes a secret
func DeleteSecret(backpack internal.Context, name string) {
	bucketName := "berglas-" + backpack.Name
	bucketPath := bucketName + "/" + name
	parts := []string{"berglas delete", bucketPath}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}
