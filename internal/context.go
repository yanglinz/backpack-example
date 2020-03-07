package internal

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// Project defines a domain-reachable app
type Project struct {
	Path string
	Name string
}

// Services represents external dependencies
type Services struct {
	Postgres bool
	Redis    bool
}

type contextYaml struct {
	Name     string
	Runtime  string
	Projects []Project
	Services Services
}

type contextGoogle struct {
	ProjectID     string
	ProjectNumber string
	Region        string
	Zone          string
}

type contextHeroku struct {
	AppName string
}

// Context for the overarching repository
type Context struct {
	Root     string
	Name     string
	Runtime  string
	Projects []Project
	Services Services
	Google   contextGoogle
	Heroku   contextHeroku
}

func parseRootPath(cmd *cobra.Command) (string, error) {
	// Get the root project path based on flag or cwd
	root := ""
	if cmd.Flag("root") != nil {
		root = cmd.Flag("root").Value.String()
	} else {
		cwd, err := os.Getwd()
		root = cwd
		if err != nil {
			return "", err
		}
	}

	// Convert to absolute path
	root, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}

	return root, nil
}

// ParseContext returns Context for a given project
func ParseContext(cmd *cobra.Command) Context {
	rootPath, err := parseRootPath(cmd)
	if err != nil {
		panic(err)
	}

	var parsedContext contextYaml
	source, err := ioutil.ReadFile(filepath.Join(rootPath, "backpack.yml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &parsedContext)
	if err != nil {
		panic(err)
	}

	google := contextGoogle{
		ProjectID:     "default-263000",
		ProjectNumber: "532331252493",
		Region:        "us-central1",
		Zone:          "us-central1-c",
	}

	heroku := contextHeroku{
		AppName: strcase.ToKebab(parsedContext.Name + "-backpack"),
	}

	context := Context{
		Root:     rootPath,
		Name:     parsedContext.Name,
		Runtime:  parsedContext.Runtime,
		Projects: parsedContext.Projects,
		Services: parsedContext.Services,
		Google:   google,
		Heroku:   heroku,
	}
	return context
}
