package opencontrol

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Standards creates a struct representing the project Standards.
type Standards struct {
	Standard map[string]struct {
		Family               string `yaml:"family"`
		Name                 string `yaml:"name"`
		Description          string `yaml:"description"`
		SupplementalGuidance string `yaml:"supplemental guidance"`
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
