package secrender

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var (
	fd FileData
)

// FileData contains the values from the config.yaml file.
type FileData struct {
	TemplatePath string
	OutputPath   string
	Variables    map[string]interface{}
}

// Secrender does variable replacement on template file(s) and creates
// new file(s) in the defined OutputDir.
func Secrender(t string, o string, tv map[string]interface{}) {
	fd.TemplatePath = t
	fd.OutputPath = o
	fd.Variables = tv

	if isTemplate(t) {
		renderFile()
	} else {
		fmt.Println("Unable to render file:", fd.TemplatePath)
	}
}

// Render the template and output the result to a file.
func renderFile() {
	createOutputPath()
	tpl := template.Must(template.ParseFiles(fd.TemplatePath))
	f, err := os.Create(fd.OutputPath)
	if err != nil {
		log.Println("Error writing file: ", err)
	}
	err = tpl.Execute(f, fd.Variables)
	if err != nil {
		log.Println("Error writing file: ", err)
	}
	f.Close()
}

func isTemplate(t string) bool {
	if filepath.Ext(t) == ".tpl" {
		return true
	}
	return false
}

func createOutputPath() {
	_, err := os.Stat(fd.OutputPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(filepath.Dir(fd.OutputPath), 0777)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
