package google

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/berglas/pkg/berglas"
	"github.com/yanglinz/backpack/internal"
	"github.com/yanglinz/backpack/symbols"
)

const namespacePrefix = "backpack-berglas-"
const namespacePrefixShort = "backpack-"
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

func bootstrapBucket(backpack internal.Context) {
	ctx := context.Background()
	bucketName := namespacePrefix + backpack.Name
	exists := bucketExists(bucketName)
	if exists {
		return
	}

	err := berglas.Bootstrap(ctx, &berglas.StorageBootstrapRequest{
		ProjectID: backpack.Google.ProjectID,
		Bucket:    bucketName,
	})
	if err != nil {
		panic(err)
	}

	CreateSecret(backpack, CreateSecretRequest{
		Env:   symbols.EnvDevelopment,
		Value: "{}",
	})
	CreateSecret(backpack, CreateSecretRequest{
		Env:   symbols.EnvProduction,
		Value: "{}",
	})
}

func bootstrapServiceAccount(backpack internal.Context) {
	// Create service account to fetch secrets
	serviceAccountName := namespacePrefixShort + backpack.Name
	parts := []string{
		"gcloud iam service-accounts create",
		serviceAccountName,
		"--project", backpack.Google.ProjectID,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	shell.Run()

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
	shell.Run()

	// Grant the global service account
	serviceAccountEmailGlobal := fmt.Sprintf(
		"serviceAccount:%s@%s.iam.gserviceaccount.com",
		"backpack-global-service",
		backpack.Google.ProjectID,
	)
	parts = []string{
		"berglas grant", grantKey,
		"--member", serviceAccountEmailGlobal,
	}
	command = strings.Join(parts, " ")
	shell = internal.GetCommand(command)
	shell.Run()
}

// BootstrapSecrets for berglas
func BootstrapSecrets(backpack internal.Context) {
	bootstrapBucket(backpack)
	bootstrapServiceAccount(backpack)
}

// CreateSecretRequest params
type CreateSecretRequest struct {
	Env   string
	Value string
}

// CreateSecret creates or updates a secret
func CreateSecret(backpack internal.Context, req CreateSecretRequest) {
	bucketName := namespacePrefix + backpack.Name

	name := secretNameDev
	if req.Env == symbols.EnvProduction {
		name = secretName
	}
	bucketPath := bucketName + "/" + name

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
	Env   string
	Value string
}

// UpdateSecrets updates the composite secrets
func UpdateSecrets(backpack internal.Context, req UpdateSecretRequest) {
	bucketName := namespacePrefix + backpack.Name

	name := secretNameDev
	if req.Env == symbols.EnvProduction {
		name = secretName
	}
	bucketPath := bucketName + "/" + name
	encodedValue := base64.StdEncoding.EncodeToString([]byte(req.Value))
	encryptionKey := "projects/" + backpack.Google.ProjectID + "/locations/global/keyRings/berglas/cryptoKeys/berglas-key"
	parts := []string{
		"berglas update", bucketPath, encodedValue,
		"--key", encryptionKey,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}

// GetSecrets fetches the composite secrets
func GetSecrets(backpack internal.Context, env string) string {
	bucketName := namespacePrefix + backpack.Name

	name := secretNameDev
	if env == symbols.EnvProduction {
		name = secretName
	}
	bucketPath := bucketName + "/" + name
	parts := []string{"berglas access", bucketPath}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	shell.Stdout = nil
	out, err := shell.Output()
	if err != nil {
		panic(err)
	}

	decodedValue, err := base64.StdEncoding.DecodeString(string(out))
	if err != nil {
		panic(err)
	}

	return string(decodedValue)
}
