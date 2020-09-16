package mkfam

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/tom-camp/gossptk/internal/pkg/config"
	"github.com/tom-camp/gossptk/internal/pkg/opencontrol"
	"github.com/tom-camp/gossptk/pkg/secrender"
	"gopkg.in/yaml.v2"
)

var (
	oc opencontrol.Config
	ct opencontrol.Certs
	st opencontrol.Standards
	cl componentList
	cf config.Config
)

// componentList is a list of filenames for files containing components.
type componentList struct {
	Satisfies []string `yaml:"satisfies,flow"`
	Families  map[string][]string
}

// getComponentList unmarshals the component YAML file.
func (cl *componentList) getComponentList(p string) {
	var tmp componentList
	cp := path.Join(p, "component.yaml")
	_, err := os.Stat(cp)
	if os.IsNotExist(err) {
		log.Panicf("File %s not found", p)
	}
	yamlFile, err := ioutil.ReadFile(cp)
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &tmp)
	if err != nil {
		log.Println(err)
	}
	for i, v := range tmp.Satisfies {
		tmp.Satisfies[i] = path.Join(p, v)
	}
	cl.Satisfies = append(cl.Satisfies, tmp.Satisfies...)
}

func (cl *componentList) groupComponentFamilies() {
	for _, v := range cl.Satisfies {
		_, f := path.Split(v)
		fam := f[0:2]
		cl.Families[fam] = append(cl.Families[fam], v)
	}
}

// MakeFamilies aggregates the Control components by control family.
func MakeFamilies() {
	for _, p := range oc.Components {
		cl.getComponentList(p)
	}
	cl.Families = make(map[string][]string)
	cl.groupComponentFamilies()
	createFamily()
}

func createFamily() {
	for i, fl := range cl.Families {
		f := Family{}
		f.parseFamily(fl, i)
		// jsonString, _ := json.Marshal(f)
		// fmt.Printf("%-v\n\r", string(jsonString))
		o := getOutPath(i)
		renderFile(o, f)
	}
}

// getOutPath creates the file to create for the family.
func getOutPath(f string) string {
	return path.Join(cf.OutputDir, "families", f+".md")
}

// Render the template and output the result to a file.
func renderFile(o string, f Family) {
	var vars map[string]interface{}
	tpl := "assets/templates/family.md.go.tpl"
	inrec, _ := yaml.Marshal(f)
	yaml.Unmarshal(inrec, &vars)
	secrender.Secrender(tpl, o, vars)
}

func init() {
	oc.Load()
	cf.LoadConfig()
	ct.Load()
	st.Load()
}
