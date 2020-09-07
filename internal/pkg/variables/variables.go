package variables

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tom-camp/gossptk/internal/pkg/config"
)

var (
	cf config.Config
	tv TemplateVars
)

// TemplateVars contains all of the file define in the keys directory.
type TemplateVars struct {
	Keys map[string]interface{}
}

// LoadTemplateVars loads all of the variable key:value pairs.
func (tv *TemplateVars) LoadTemplateVars() {
	keyFiles := getFiles(cf.KeyDir)
	for _, v := range keyFiles {
		tv.parseJSON(v)
	}
}

// Load all of the JSON files into a slice of strings.
func getFiles(p string) []string {
	var files []string
	err := filepath.Walk(p,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			ftype := filepath.Ext(path)
			if ftype == ".json" {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return files
}

// Unmarshall JSON.
func (tv *TemplateVars) parseJSON(j string) {
	jsonFile, err := ioutil.ReadFile(j)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(jsonFile, &tv.Keys)
	if err != nil {
		log.Println(err)
	}
}

func init() {
	cf.LoadConfig()
}
