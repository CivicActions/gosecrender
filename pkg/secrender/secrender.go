// Package secrender implements variable replacement for templates.
//
// The templates use Golangs text/template package and dot syntax. For
// example {{.title}}.
//
// Given a template, a path to output the rendered template and a map
// of variables to use to replace the template "Actions", Secrender will
// will create a new file with the template actions substituted with the
// variables.
package secrender

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	fd          FileData
	secTemplate template.Template
)

// FileData creates a struct to hold the data for the file to be render.
// TemplatePath is a string to the template to be rendered.
// OutputPath is the path where Secrender should create the rendered file.
// Variables is a map of the variables to use when doing the variable replacement
// in the template.
type FileData struct {
	TemplatePath string
	OutputPath   string
	Variables    map[string]interface{}
}

// Secrender does variable replacement on template file(s) and creates
// new file(s) in the defined OutputDir.
// Parameter t is a string representing a path to a template.
// Parameter o is a string representing a path where Secrender should output the
// rendered template.
// Parameter tv is a map of variables to use for the variable replacement in the
// template.
func Secrender(t string, o string, tv map[string]interface{}) {
	fd.TemplatePath = t
	fd.OutputPath = o
	fd.Variables = tv

	if isTemplate(t) {
		renderFile()
	}
}

// renderFile does the heavy lifting, rendering the template and writing it to
// the OutputPath.
// funcMap allows us to add functions to use within the templates. Currently,
// the only function available is strings.ToUpper(). To use it the function, use
// pipe notation. For example, {{.Title | ToUpper}}.
func renderFile() {
	fmt.Println("Creating Out path", fd.OutputPath)
	_, name := filepath.Split(fd.TemplatePath)
	createOutputPath()
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}
	tpl := template.Must(template.New(name).Funcs(funcMap).ParseFiles(fd.TemplatePath))
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

// Check that the template path that is given is indeed a template.
func isTemplate(t string) bool {
	if filepath.Ext(t) == ".tpl" {
		return true
	}
	return false
}

// Creates the path to the output file, creating any directories that are needed.
func createOutputPath() {
	_, err := os.Stat(fd.OutputPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(filepath.Dir(fd.OutputPath), 0777)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
