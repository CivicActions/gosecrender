package opencontrol

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Component creates a struct representing an OpenControl Component.
type Component struct {
	Name                  string          `yaml:"name"`
	Key                   string          `yaml:"key"`
	System                bool            `yaml:"system"`
	SchemaVersion         string          `yaml:"schema_version"`
	DocumentationComplete bool            `yaml:"documentation_complete"`
	ResponsibleRole       string          `yaml:"responsible_role"`
	References            []references    `yaml:"references"`
	Verifications         []verifications `yaml:"verifications"`
	Satisfies             []Ctrl          `yaml:"satisfies"`
}

type references struct {
	Name    string `yaml:"name"`
	Path    string `yaml:"path"`
	RefType string `yaml:"type"`
}

type verifications struct {
	Key     string `yaml:"key"`
	Name    string `yaml:"name"`
	Path    string `yaml:"path"`
	VerType string `yaml:"type"`
}

// Load loads a struct with data representing a Component.
func (cp *Component) Load(p string) {
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		log.Panicf("File %s not found", p)
	}
	yamlFile, err := ioutil.ReadFile(p)
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &cp)
	if err != nil {
		log.Println(err)
	}
}
