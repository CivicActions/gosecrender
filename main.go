// Package secrender
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

var wg sync.WaitGroup
var cf Config
var tv TemplateVars

// Create a log file and set logging output.
func setupLogging() {
	l, err := os.Create("error.log")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()
	log.SetOutput(l)
}

// Config contains the values from the config.json file.
type Config struct {
	KeyDir      string `json:"keyDir"`
	TemplateDir string `json:"templateDir"`
	OutputDir   string `json:"outputDir"`
}

// Unmarshal the config.json.
func (cf *Config) loadConfig() {
	jsonFile, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(jsonFile, &cf)
	if err != nil {
		log.Println(err)
	}
}

// TemplateVars maps all of the template variables from the keys/ directory
// unmarshalled into maps.
type TemplateVars struct {
	Keys map[string]interface{}
}

func (t *TemplateVars) loadTemplateVars() {
	keyFiles := getFiles(cf.KeyDir)
	for _, v := range keyFiles {
		t.parseJSON(v)
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
func (t *TemplateVars) parseJSON(j string) {
	jsonFile, err := ioutil.ReadFile(j)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(jsonFile, &t.Keys)
	if err != nil {
		log.Println(err)
	}
}

// Do variable replacement on template files and create new files
// in the defined OutputDir.
func secrender() {
	switch filepath.Ext(cf.TemplateDir) {
	case "":
		crawlTemplates()
	case ".tpl":
		if filepath.Ext(cf.TemplateDir) == ".tpl" {
			renderFile(cf.TemplateDir)
		}
	}
}

// Walk into the template directory looking for tpl files.
func crawlTemplates() {
	err := filepath.Walk(cf.TemplateDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".tpl" {
				wg.Add(1)
				go renderFile(path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

// Render the template and output the result to a file.
func renderFile(p string) {
	r := strings.NewReplacer(cf.TemplateDir, cf.OutputDir, ".tpl", "")
	opath := r.Replace(p)
	fmt.Println("Writing file:", opath)

	createFilepath(filepath.Dir(opath))

	tpl := template.Must(template.ParseFiles(p))
	f, err := os.Create(opath)
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

// Create the directory structure for a given template output.
func createFilepath(o string) {
	_, err := os.Stat(o)
	if os.IsNotExist(err) {
		fmt.Println("Creating directory...", o)
		os.MkdirAll(o, 0777)
	}
}

// Set up logging, load config and load the template variables.
func init() {
	setupLogging()
	cf.loadConfig()
	tv.loadTemplateVars()
}

func main() {
	secrender()
	wg.Wait()
}
