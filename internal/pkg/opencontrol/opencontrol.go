package opencontrol

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var oc OpenControl

// OpenControl creates a struct representing the OpenControl settings.
type OpenControl struct {
	SchemaVersion  string                 `yaml:"schema_version"`
	Name           string                 `yaml:"name"`
	Metadata       map[string]interface{} `yaml:"metadata"`
	Components     []string               `yaml:"components"`
	Standards      []string               `yaml:"standards"`
	Certifications []string               `yaml:"certifications"`
}

// LoadOpenControl loads the values of the opencontrol.yaml file.
func (oc *OpenControl) LoadOpenControl() {
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
	fmt.Printf("%+v", oc)
	fmt.Println(oc.Name)
}
