package secrender

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

var (
	wg   sync.WaitGroup
	fd   FileData
	tv   TemplateVars
	tpl  string
	out  string
	keys string
	help bool
)

// FileData contains the values from the config.json file.
type FileData struct {
	KeyDir       string
	TemplatePath string
	OutputPath   string
}

// Set all of the FileData parameters.
func (fd *FileData) loadParams() {
	flag.StringVar(&out, "o", "", "Output Directory")
	flag.StringVar(&tpl, "t", "", "Template filepath")
	flag.StringVar(&keys, "k", "", "Key directory path")
	flag.Parse()

	fd.OutputPath = out
	fd.TemplatePath = tpl
	fd.KeyDir = keys
}

// TemplateVars maps all of the template variables from the keys/ directory
// unmarshalled into maps.
type TemplateVars struct {
	Keys map[string]interface{}
}

func (tv *TemplateVars) loadTemplateVars() {
	if fd.KeyDir != "" {
		keyFiles := getFiles(fd.KeyDir)
		fmt.Printf("Loading variables from JSON files in %s\n\r", fd.KeyDir)
		for _, v := range keyFiles {
			tv.parseJSON(v)
		}
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

// Secrender does variable replacement on template file(s) and creates
// new file(s) in the defined OutputDir.
func Secrender() {
	if isTemplate(fd.TemplatePath) {
		wg.Add(1)
		renderFile()
	} else {
		fmt.Println("Unable to render file:", fd.TemplatePath)
	}
}

// Render the template and output the result to a file.
func renderFile() {
	log.Printf("Writing file: %s\r\n", fd.OutputPath)
	createOutput()

	tpl := template.Must(template.ParseFiles(fd.TemplatePath))
	f, err := os.Create(fd.OutputPath)
	if err != nil {
		log.Println("Error writing file: ", err)
	}

	err = tpl.Execute(f, tv)
	if err != nil {
		log.Println("Error writing file: ", err)
	}
	f.Close()
	wg.Done()
}

func isTemplate(t string) bool {
	if filepath.Ext(t) == ".tpl" {
		return true
	}
	return false
}

func createOutput() {
	_, err := os.Stat(fd.OutputPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(filepath.Dir(fd.OutputPath), 0777)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

// Set config and load the template variables.
func init() {
	flag.BoolVar(&help, "help", false, "Help text")
	if help {
		fmt.Println("Secrender Help")
		flag.Usage()
	}
	fd.loadParams()
	tv.loadTemplateVars()
}
