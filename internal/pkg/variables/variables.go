package variables

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tom-camp/gossptk/internal/pkg/config"
	"gopkg.in/yaml.v2"
)

var (
	cf config.Config
	tv TemplateVars
)

// TemplateVars creates a struct to hold all of the variables needed for the
// template variable replacement used by the Secrender package.
type TemplateVars struct {
	Keys map[string]interface{}
}

// LoadTemplateVars loads all of the variable key:value pairs.
func (tv *TemplateVars) LoadTemplateVars() {
	keyFiles := getFiles(cf.KeyDir)
	for _, v := range keyFiles {
		tv.parseYAML(v)
	}
}

// Load all of the YAML files into a slice of strings.
func getFiles(p string) []string {
	var files []string
	err := filepath.Walk(p,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			ftype := filepath.Ext(path)
			if ftype == ".yaml" {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return files
}

// Unmarshall YAML.
func (tv *TemplateVars) parseYAML(j string) {
	yamlFile, err := ioutil.ReadFile(j)
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal(yamlFile, &tv.Keys)
	if err != nil {
		log.Println(err)
	}
}

// Load the configuration files.
func init() {
	cf.LoadConfig()
}
