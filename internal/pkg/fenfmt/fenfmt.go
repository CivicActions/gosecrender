// Package fenfmt uses the Fen Format aka Fen Labalme Format aka FenFmt
// to handle OpenControl data. The FenFmt uses the component.yaml file to
// include separate control family files. For example, the Contractor component.yaml
// might look like:
// -------------------------------------
// name: Project
// schema_version: 3.0.0
// satisfies:
//   - AC-ACCESS_CONTROL.yaml
//   - AT-AWARENESS_AND_TRAINING.yaml
//   - AU-AUDIT_AND_ACCOUNTABILITY.yaml
//   - CA-SECURITY_ASSESSMENT_AND_AUTHORIZATION.yaml
//   - CM-CONFIGURATION_MANAGEMENT.yaml
package fenfmt

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/civicactions/gossptk/internal/pkg/config"
	"github.com/civicactions/gossptk/internal/pkg/opencontrol"
	"gopkg.in/yaml.v2"
)

var (
	c   config.Config
	oc  opencontrol.OpenControl
	cmp opencontrol.Component
	fn  FenComponent
)

// FenComponent creates a struct representing a Fen Formatted component.yaml
// with the Satisfies includes list.
type FenComponent struct {
	Name          string   `yaml:"name"`
	SchemaVersion string   `yaml:"schema_version"`
	Satisfies     []string `yaml:"satisfies"`
}

// DeFen aggregates the Controls from within the individual FenFmt Control
// Family files into a component.yaml file for each component.
func DeFen() {
	if isFenFmt() {
		for _, p := range oc.Components {
			loadYaml(p)
			cmp.Name = filepath.Base(p)
			cmp.SchemaVersion = fn.SchemaVersion
			loadControls(p)
		}
	}
}

// Marshal a component.yaml file
func loadYaml(p string) {
	cp := p + "/component.yaml"
	_, err := os.Stat(cp)
	if os.IsNotExist(err) {
		log.Printf("No component.yaml file in %s.\r\n", p)
	}
	yamlFile, err := ioutil.ReadFile(cp)
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &fn)
	if err != nil {
		log.Println(err)
	}
}

// Create an OpenControl Control and add the controls.
func loadControls(p string) {
	for _, v := range fn.Satisfies {
		nc := new(opencontrol.Control)
		nc.Load(filepath.Join(p, v))
		aggregateSatisfies(nc.Satisfies)
		writeDefened(p)
	}
}

// Add the Controls to the component.yaml Satisifies list.
func aggregateSatisfies(s []opencontrol.Ctrl) {
	for _, v := range s {
		cmp.Satisfies = append(cmp.Satisfies, v)
	}
}

func writeDefened(p string) {
	out := filepath.Join("rendered/ocformatted", p, "component.yaml")
	err := os.MkdirAll(filepath.Dir(out), 0777)
	if err != nil {
		log.Println("Unable to create dir for file", out)
	}

	f, err := os.Create(out)
	if err != nil {
		log.Println("Unable to create file", out)
	}
	y, err := yaml.Marshal(&cmp)
	if err != nil {
		log.Println(err)
	}
	w, err := f.Write(y)
	if err != nil {
		log.Println(err)
	}
	if w <= 0 {
		log.Printf("%d bytes written for: %s", w, out)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}

// Check that configuration is set to Fen Format.
func isFenFmt() bool {
	if strings.ToLower(c.Format) == "fenfmt" {
		return true
	}
	return false
}

func init() {
	c.LoadConfig()
	oc.Load()
}
