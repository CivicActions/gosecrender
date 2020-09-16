package mkfam

import (
	"path/filepath"

	"github.com/tom-camp/gossptk/internal/pkg/opencontrol"
)

var component string

// Family creates a struct that will contain all of the information to create the
// [FAMILY].md file. Control Families are a grouping of similar controls. The
// Family struct Control fields is a map of all of the fields for a given Family,
// keyed by the Control ID. For instance the Access Control family would contain
// Controls["AC-2"] which would hold the values for all of the components that
// touch the AC-2 Control.
type Family struct {
	Name          string
	Certification []string
	Controls      map[string]interface{}
}

// Satisfies creates a struct that will contain all of the  components that
// pertain to a given control. For instance, we for AC-2 components for the
// web host, the contractor and the application might pertain to the control.
// All of component narratives would be Narratives field slice of Narrative.
type Satisfies struct {
	CtrlKey     string
	CtrlName    string
	Description string
	Status      string
	Narratives  []Narrative
}

// Narrative creates a struct that contains a slice of Narrative text. The Key
// field is a map keyed by the Open Control Narrative Key so that we can group
// all of the component narratives.
type Narrative struct {
	Key map[string][]NarrativeText
}

// NarrativeText creates a struct that contains a singal key:value pair with the
// key being the component and the value being the narrative.
type NarrativeText struct {
	Text map[string]string
}

// parseFamily maps the fields in an OpenControl component file to a Family
// struct.
// The argument p is slice of paths pointing to the component YAML fiels that
// pertain to the Family. The argument a is a string of the abbreviation for
// the family. In the case of Access Control, a would be "AC".
func (f *Family) parseFamily(p []string, a string) {
	var nr []opencontrol.Ctrl
	for _, fp := range p {
		component = filepath.Base(filepath.Dir(fp))
		c := opencontrol.Controls{}
		c.Load(fp)
		if len(f.Name) <= 0 {
			f.Name = a + ": " + c.Family
		}
		nr = append(nr, c.Satisfies...)
	}
	f.Controls = make(map[string]interface{})
	f.parseControl(nr)
}

// parseControl creates a Satisfies struct for every Control passed in the
// slice of Ctrl parameter. The Satisfy structs make up the
// Parameter o is a slice of opencontorl.Ctrl structs defining OpenControl
// Controls.
func (f *Family) parseControl(o []opencontrol.Ctrl) {
	for _, v := range o {
		s := new(Satisfies)
		s.CtrlKey = v.ControlKey
		s.CtrlName = v.ControlName
		s.setDescription()
		s.parseNarratives(v.Narratives)
		f.Controls[s.CtrlKey] = s
	}
}

// parseNarratives creates a struct of NarrativeText and appends it to the
// Narrative slice of NarrativeText.
func (s *Satisfies) parseNarratives(o []opencontrol.Narratives) {
	var n Narrative
	n.Key = make(map[string][]NarrativeText)
	k := "no key"
	for _, on := range o {
		if len(on.Key) >= 0 {
			k = on.Key
		}
		nt := NarrativeText{
			Text: map[string]string{component: on.Text},
		}
		n.Key[k] = append(n.Key[k], nt)
	}
	s.Narratives = append(s.Narratives, n)
}

// setDescription gets the Control Description, and, if applicable, Supplemental
// Guidance for a give Control.
func (s *Satisfies) setDescription() {
	standard := st.Standard[s.CtrlKey]
	text := standard.Description
	if len(standard.SupplementalGuidance) > 0 {
		text = text + "\r\n\r\n**Supplemental Guidance**\r\n\r\n" + standard.SupplementalGuidance
	}
	s.Description = text
}
