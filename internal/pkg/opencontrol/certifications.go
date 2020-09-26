package opencontrol

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Certifications creates a struct representing the project Certificates.
type Certifications struct {
	Standards map[string]interface{}
}

// Load loads a struct with data representing the Certificates.
func (c *Certifications) Load() {
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
