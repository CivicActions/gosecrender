package opencontrol

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Narratives creates a struct matching OpenControl narrative.
type Narratives struct {
	Key  string `yaml:"key"`
	Text string `yaml:"text"`
}

// Ctrl creates a struct matching OpenControl control.
type Ctrl struct {
	ControlKey          string       `yaml:"control_key"`
	ControlName         string       `yaml:"control_name"`
	StandardKey         string       `yaml:"standard_key"`
	CoveredBy           []string     `yaml:"covered_by"`
	SecurityControlType string       `yaml:"security_control_type"`
	Narratives          []Narratives `yaml:"narrative"`
	Status              string       `yaml:"implementation_status"`
	Component           string
}

// Controls creates a struct matching OpenControl component.
type Controls struct {
	Family      string `yaml:"family"`
	DocComplete string `yaml:"documentation_complete"`
	Satisfies   []Ctrl `yaml:"satisfies"`
}

// Load a Control given a path to the control YAML file.
func (c *Controls) Load(p string) {
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
