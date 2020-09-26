package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var cf Config

// Config contains configuration variables.
type Config struct {
	TemplateDir string `yaml:"templateDir"`
	OutputDir   string `yaml:"outputDir"`
	KeyDir      string `yaml:"keyDir"`
	Format      string `yaml:"format"`
}

// LoadConfig loads variables from the config.yaml file.
func (cf *Config) LoadConfig() {
	_, err := os.Stat("configs/config.yaml")
	if os.IsNotExist(err) {
		log.Panic("No config.yaml file found.")
	}
	yamlFile, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &cf)
	if err != nil {
		log.Println(err)
	}
}
