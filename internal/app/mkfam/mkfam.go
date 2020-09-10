package mkfam

import (
	"fmt"

	"github.com/tom-camp/gossptk/internal/pkg/opencontrol"
)

var oc opencontrol.Config
var cp opencontrol.Component
var ct opencontrol.Certs
var st opencontrol.Standards

// Family defines a struct containing all of the information to create the
// [FAMILY].md file.
type Family struct {
	Name          string
	Certification []string
	ComponentName string
	Controls      []struct {
		description string
		Control     []map[string]interface{}
	}
}

// MakeFamilies aggregates the Control components by control family.
func MakeFamilies() {
	fmt.Println(st.Standard["AC-1"].Description)
	// for _, p := range oc.Components {
	// 	_, n := filepath.Split(p)
	// 	f := Family{}
	// 	f.Name = n
	// }
}

func getNarratives() {

}

func init() {
	ct.LoadCerts()
	st.LoadStandards()
}
