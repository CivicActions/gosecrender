// Package secrender
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var wg sync.WaitGroup

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

func (t *TemplateVars) loadTemplateVars(td string) {
	keyFiles := getFiles(td)
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

// Create a log file and set logging output.
func setupLogging() {
	l, err := os.Create("error.log")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()
	log.SetOutput(l)
}

// Do variable replacement on template files and create new files
// in the defined OutputDir.
func secrender(cf Config, t TemplateVars) {
	switch filepath.Ext(cf.TemplateDir) {
	case "":
		crawlTemplates(cf.TemplateDir, cf.OutputDir, t)
	case ".tpl":
		if filepath.Ext(cf.TemplateDir) == ".tpl" {
			renderFile(cf.TemplateDir, cf.TemplateDir, cf.OutputDir, t)
		}
	}

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
}

func crawlTemplates(tm string, o string, v TemplateVars) {
	err := filepath.Walk(tm,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".tpl" {
				wg.Add(1)
				go renderFile(tm, path, o, v)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

// Render the template and output the result to a file.
func renderFile(tm string, t string, o string, v TemplateVars) {
	r := strings.NewReplacer(tm, o, ".tpl", "")
	opath := r.Replace(t)
	fmt.Println("Creating file:", opath)

	createFilepath(filepath.Dir(opath))

	tpl := template.Must(template.ParseFiles(t))
	f, err := os.Create(opath)
	if err != nil {
		log.Println("Error creating file: ", err)
	}

	err = tpl.Execute(f, v)
	if err != nil {
		log.Print("Error writing file: ", err)
	}
	f.Close()
	wg.Done()
}

// Create the directory structure for a given template output.
func createFilepath(o string) {
	_, err := os.Stat(o)
	if os.IsNotExist(err) {
		fmt.Println("creating...", o)
		os.MkdirAll(o, 0777)
	}
}

func main() {
	setupLogging()
	var cf Config
	cf.loadConfig()
	var t TemplateVars
	t.loadTemplateVars(cf.KeyDir)

	secrender(cf, t)
	wg.Wait()
}
