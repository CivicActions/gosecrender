package mkfam

import (
	"path/filepath"

	"github.com/tom-camp/gossptk/internal/pkg/opencontrol"
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
	Controls      map[string]Control
}

// Control creates a struct...
type Control struct {
	CtrlKey     string
	CtrlName    string
	Description string
	Status      string
	Narratives  map[string][]NarrativeText
}

type NarrativeText struct {
	Component string
	Text      string
}

// parseFamily maps the fields in an OpenControl component file to a Family
// struct.
// The argument p is slice of paths pointing to the component YAML fiels that
// pertain to the Family. The argument a is a string of the abbreviation for
// the family. In the case of Access Control, a would be "AC".
func (f *Family) parseFamily(p []string, a string) {
	f.Controls = make(map[string]Control)
	for _, fp := range p {
		component = filepath.Base(filepath.Dir(fp))
		c := opencontrol.Controls{}
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
		c, ok := f.Controls[v.ControlKey]
		if !ok {
			c = newControl(v)
		} else {
		}
		c.parseNarratives(v.Narratives)
		f.Controls[c.CtrlKey] = c
	}
}

func newControl(v opencontrol.Ctrl) Control {
	c := new(Control)
	c.CtrlKey = v.ControlKey
	c.CtrlName = v.ControlName
	c.setDescription()
	return *c
}

// parseNarratives creates a struct of NarrativeText and appends it to the
// Narrative slice of map[string]string.
func (c *Control) parseNarratives(o []opencontrol.Narratives) {
	if c.Narratives == nil {
		c.Narratives = make(map[string][]NarrativeText)
	}
	for _, on := range o {
		k := "no key"
		if len(on.Key) > 0 {
			k = on.Key
		}
		nt := NarrativeText{
			Component: component,
			Text:      on.Text,
		}
		c.Narratives[k] = append(c.Narratives[k], nt)
	}
}

// setDescription gets the Control Description, and, if applicable, Supplemental
// Guidance for a give Control.
func (c *Control) setDescription() {
	standard := st.Standard[c.CtrlKey]
	text := standard.Description
	if len(standard.SupplementalGuidance) > 0 {
		text = text + "\r\n\r\n**Supplemental Guidance**\r\n\r\n" + standard.SupplementalGuidance
	}
	c.Description = text
}
