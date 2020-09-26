package opencontrol

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var oc OpenControl

// OpenControl creates a struct representing the OpenControl OpenControl settings.
type OpenControl struct {
	Name           string   `yaml:"name"`
	SchemaVersion  string   `yaml:"schema_version"`
	Metadata       metadata `yaml:"metadata"`
	Components     []string `yaml:"components"`
	Certifications []string `yaml:"certifications"`
	Standards      []string `yaml:"standards"`
	Dependencies   struct {
		certifications []dependencies `yaml:"certifications"`
		standards      []dependencies `yaml:"standards"`
		systems        []dependencies `yaml:"systems"`
	} `yaml:"dependencies"`
}

type metadata struct {
	Description string   `yaml:"description"`
	Maintainers []string `yaml:"maintainers"`
}

type dependencies struct {
	url        string `yaml:"url"`
	contextDir string `yaml:"contextdir"`
	revision   string `yaml:"revision"`
}

// Load loads the values of the opencontrol.yaml file.
func (oc *OpenControl) Load() {
	_, err := os.Stat("opencontrol.yaml")
	if os.IsNotExist(err) {
		log.Panic("No config.yaml file found.")
	}
	yamlFile, err := ioutil.ReadFile("opencontrol.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &oc)
	if err != nil {
		log.Println(err)
	}
}

func init() {
	oc.Load()
}
