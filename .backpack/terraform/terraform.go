package terraform

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/rodaine/hclencoder"
	"github.com/yanglinz/backpack/internal"
	"github.com/yanglinz/backpack/symbols"
)

type gcpProvider struct {
	ID      string `hcl:",key"`
	Project string `hcl:"project" hcle:"omitempty"`
	Region  string `hcl:"zone" hcle:"omitempty"`
	Zone    string `hcl:"region" hcle:"omitempty"`
}

type variable struct {
	ID      string `hcl:",key"`
	Type    string `hcl:"type" hcle:"omitempty"`
	Default string `hcl:"default" hcle:"omitempty"`
}

type webModule struct {
	ID                   string `hcl:",key"`
	Source               string `hcl:"source" hcle:"omitempty"`
	ContextName          string `hcl:"context_name" hcle:"omitempty"`
	ProjectName          string `hcl:"project_name" hcle:"omitempty"`
	DjangoSettingsModule string `hcl:"django_settings_module" hcle:"omitempty"`
	ImageTag             string `hcl:"image_tag" hcle:"omitempty"`
	GCPProject           string `hcl:"gcp_project" hcle:"omitempty"`
}

type autoconfig struct {
	Providers []gcpProvider `hcl:"provider"`
	Variables []variable    `hcl:"variable"`
	Modules   []webModule   `hcl:"module"`
}

func getCloudrunConfig(backpack internal.Context) autoconfig {
	providers := []gcpProvider{
		gcpProvider{
			ID:      "google",
			Project: backpack.Google.ProjectID,
			Region:  backpack.Google.Region,
			Zone:    backpack.Google.Zone,
		},
	}

	variables := []variable{
		variable{
			ID:      "context_name",
			Type:    "string",
			Default: backpack.Name,
		},
		variable{
			ID:      "image_tag",
			Type:    "string",
			Default: "latest",
		},
	}

	project := backpack.Projects[0]
	modules := []webModule{
		webModule{
			ID:                   "app_web",
			Source:               "../.backpack/terraform/web-django-module",
			ContextName:          "${var.context_name}",
			ProjectName:          "core",
			DjangoSettingsModule: strings.ReplaceAll(project.Path, "/", ".") + ".settings",
			ImageTag:             "${var.image_tag}",
			GCPProject:           backpack.Google.ProjectID,
		},
	}

	config := autoconfig{
		Providers: providers,
		Variables: variables,
		Modules:   modules,
	}
	return config
}

func getHerokuConfig(backpack internal.Context) autoconfig {
	config := autoconfig{
		Variables: nil,
	}
	return config
}

// CreateConfig generates the terraform config
func CreateConfig(backpack internal.Context) {
	input := getCloudrunConfig(backpack)
	if backpack.Runtime == symbols.RuntimeHeroku {
		input = getHerokuConfig(backpack)
	}

	hcl, err := hclencoder.Encode(input)
	if err != nil {
		panic(err)
	}

	configPath := filepath.Join(backpack.Root, "terraform/backpack.tf")
	err = ioutil.WriteFile(configPath, hcl, 0644)
	if err != nil {
		panic(err)
	}
}
