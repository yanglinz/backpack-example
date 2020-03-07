package heroku

import (
	"strings"

	"github.com/yanglinz/backpack/internal"
)

// PutSecretRequest params
type PutSecretRequest struct {
	App   string
	Name  string
	Value string
}

// PutSecret creates/updates a secret with its value
func PutSecret(req PutSecretRequest) {
	parts := []string{
		"heroku config:set",
		req.Name + "=" + req.Value,
		"-a", req.App,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}
