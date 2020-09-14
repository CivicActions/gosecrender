package mkfam

// Family defines a struct containing all of the information to create the
// [FAMILY].md file.
type Family struct {
	Name          string
	Certification []string
	ComponentName string
	Controls      map[string]map[string]map[string]string
}
