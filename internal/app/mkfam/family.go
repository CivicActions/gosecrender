package mkfam

import (
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/civicactions/gossptk/internal/pkg/opencontrol"
)

var (
	f         Family
	component string
)

// Family creates a struct that will contain all of the information to create the
// [FAMILY].md file. Control Families are a grouping of similar controls. The
// Family struct Control fields is a map of all of the fields for a given Family,
// keyed by the Control ID. For instance the Access Control family would contain
// Controls["AC-2"] which would hold the values for all of the components that
// touch the AC-2 Control.
type Family struct {
	Name          string
	Certification []string
	Controls      map[float64]Control
}

// Control creates a struct representing an OpenControl control.
type Control struct {
	CtrlKey     string
	CtrlName    string
	Description string
	Status      string
	Narratives  map[string][]interface{}
}

// parseFamily maps the fields in an OpenControl component file to a Family
// struct.
// The argument p is slice of paths pointing to the component YAML fiels that
// pertain to the Family. The argument a is a string of the abbreviation for
// the family. In the case of Access Control, a would be "AC".
func (f *Family) parseFamily(p []string, a string) {
	f.Controls = make(map[float64]Control)
	for _, fp := range p {
		component = filepath.Base(filepath.Dir(fp))
		c := opencontrol.Control{}
		c.Load(fp)
		if len(f.Name) <= 0 {
			f.Name = a + ": " + c.Family
		}
		f.parseControl(c.Satisfies)
	}
}

// parseControl creates a Satisfies struct for every Control passed in the
// slice of Ctrl parameter. The Satisfy structs make up the
// Parameter o is a slice of opencontorl.Ctrl structs defining OpenControl
// Controls.
func (f *Family) parseControl(o []opencontrol.Ctrl) {
	for _, v := range o {
		sk := getSortKey(v.ControlKey)
		c, ok := f.Controls[sk]
		if !ok {
			c = newControl(v)
		}
		c.parseNarratives(v.Narratives)
		f.Controls[sk] = c
	}
}

func newControl(v opencontrol.Ctrl) Control {
	c := new(Control)
	c.CtrlKey = v.ControlKey
	c.CtrlName = v.ControlName
	c.setDescription()
	return *c
}

// parseNarratives creates a map[string]interface{} and appends it to the
// Narrative slice of map[string]string.
func (c *Control) parseNarratives(o []opencontrol.Narratives) {
	if c.Narratives == nil {
		c.Narratives = make(map[string][]interface{})
	}
	for _, on := range o {
		k := "no key"
		if len(on.Key) > 0 {
			k = on.Key
		}
		nt := map[string]interface{}{
			"Component": component,
			"Text":      on.Text,
		}
		c.Narratives[k] = append(c.Narratives[k], nt)
	}
}

// setDescription gets the Control Description, and, if applicable,
// Supplemental Guidance for a give Control.
func (c *Control) setDescription() {
	standard := st.Standard[c.CtrlKey]
	text := standard.Description
	if len(standard.SupplementalGuidance) > 0 {
		text = text + "\r\n\r\n**Supplemental Guidance**\r\n\r\n" + standard.SupplementalGuidance
	}
	c.Description = text
}

// getSortKey creates a key so that Controls are sorted properly. If we use
// the Control ID controls will be sorted as strings, so AC-1 will be followed
// by AC-10 rather than AC-2. We are using the ID value as a float rather than
// an int to accomodate control enhancements, for example AC-17 (1) will use
// the key 17.1.
func getSortKey(k string) float64 {
	re := regexp.MustCompile(`(\d+)`)
	result := re.FindAllString(k, -1)
	var nk string
	if len(result) > 1 {
		nk = strings.Join(result, ".")
	} else {
		nk = result[0]
	}
	f, err := strconv.ParseFloat(nk, 32)
	if err != nil {
		log.Println("Key error")
	}
	return f
}
