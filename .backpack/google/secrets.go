package google

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/berglas/pkg/berglas"
	"github.com/yanglinz/backpack/internal"
)

const namespacePrefix = "backpack-berglas-"
const secretName = "BACKPACK_VARS_JSON"
const secretNameDev = "BACKPACK_VARS_DEV_JSON"

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
	serviceAccountName := namespacePrefix + backpack.Name
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
	bucketName := namespacePrefix + backpack.Name
	grantKey := bucketName + "/" + secretName
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

func bootstrapBucket(backpack internal.Context) {
	ctx := context.Background()
	bucketName := namespacePrefix + backpack.Name
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
		Name:  secretName,
		Value: "{}",
	})
	CreateSecret(backpack, CreateSecretRequest{
		Name:  secretNameDev,
		Value: "{}",
	})
}

// BootstrapSecrets for berglas
func BootstrapSecrets(backpack internal.Context) {
	bootstrapBucket(backpack)
	bootstrapServiceAccount(backpack)
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
	bucketName := namespacePrefix + backpack.Name
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
	bucketName := namespacePrefix + backpack.Name
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
