package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var cf Config

// Config contains configuration variables.
type Config struct {
	TemplateDir string `json:"templateDir"`
	OutputDir   string `json:"outputDir"`
	KeyDir      string `json:"keyDir"`
}

// LoadConfig loads variables from the config.json file.
func (cf *Config) LoadConfig() {
	_, err := os.Stat("config.json")
	if os.IsNotExist(err) {
		log.Panic("No config.json file found.")
	}
	jsonFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonFile, &cf)
	if err != nil {
		log.Println(err)
	}
}
