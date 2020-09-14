package opencontrol

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var oc Config

// Config creates a struct representing the OpenControl Config settings.
type Config struct {
	Name           string                 `yaml:"name"`
	SchemaVersion  string                 `yaml:"schema_version"`
	Metadata       map[string]interface{} `yaml:"metadata"`
	Components     []string               `yaml:"components"`
	Standards      []string               `yaml:"standards"`
	Certifications []string               `yaml:"certifications"`
}

// Load loads the values of the opencontrol.yaml file.
func (oc *Config) Load() {
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

// Component creates a struct representing an OpenControl Component.
type Component struct {
	Name                  string
	Key                   string
	DocumentationComplete bool `yaml:"documentation_complete"`
	SchemaVersion         string
	References            []struct {
		Name    string `yaml:"name"`
		Path    string `yaml:"path"`
		RefType string `yaml:"type"`
	} `yaml:"references"`
	Verifications []struct {
		Key     string `yaml:"key"`
		Name    string `yaml:"name"`
		Path    string `yaml:"path"`
		VerType string `yaml:"type"`
	} `yaml:"verifications"`
	Satisfies []struct {
		StandardKey string `yaml:"standard_key"`
		ControlKey  string `yaml:"control_key"`
		Narrative   []struct {
			Key  string `yaml:"key"`
			Text string `yaml:"text"`
		} `yaml:"narrative"`
		ImplementationStatuses []string `yaml:"implementation_statuses"`
		ControlOrigins         []string `yaml:"control_origins"`
		Parameters             []struct {
			Key  string `yaml:"key"`
			Text string `yaml:"text"`
		} `yaml:"parameters"`
		CoveredBy []string `yaml:"covered_by"`
	} `yaml:"satisfies"`
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

// Certs creates a struct representing the project Certificates.
type Certs struct {
	Standards map[string]interface{}
}

// Load loads a struct with data representing the Certificates.
func (c *Certs) Load() {
	for _, p := range oc.Certifications {
		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			log.Panicf("File %s not found", p)
		}
		yamlFile, err := ioutil.ReadFile(p)
		if err != nil {
			log.Println(err)
		}
		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			log.Println(err)
		}
	}
}

// Standards creates a struct representing the project Standards.
type Standards struct {
	Standard map[string]struct {
		Family               string `yaml:"family"`
		Name                 string `yaml:"name"`
		Description          string `yaml:"description"`
		supplementalGuidance string `yaml:"supplemental guidance"`
		relatedControls      string `yaml:"related controls"`
		references           string `yaml:"references"`
	}
}

// Load loads a struct with data representing the Standards.
func (s *Standards) Load() {
	for _, p := range oc.Standards {
		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			log.Panicf("File %s not found", p)
		}
		yamlFile, err := ioutil.ReadFile(p)
		if err != nil {
			log.Println(err)
		}
		err = yaml.Unmarshal(yamlFile, &s.Standard)
		if err != nil {
			log.Println(err)
		}
	}
}

func init() {
	oc.Load()
}
