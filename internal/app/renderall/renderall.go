package renderall

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/civicactions/gossptk/internal/pkg/config"
	"github.com/civicactions/gossptk/internal/pkg/variables"
	"github.com/civicactions/gossptk/pkg/secrender"
)

var (
	wg  sync.WaitGroup
	cf  config.Config
	tv  variables.TemplateVars
	ctr int
)

// RenderAll renders all templates in a given directory.
func RenderAll() {
	err := filepath.Walk(cf.TemplateDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".tpl" {
				wg.Add(1)
				go doSecrender(path, &wg)
				ctr++
				wg.Wait()
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("RenderAll created %d files.\r\n", ctr)
}

func doSecrender(p string, wg *sync.WaitGroup) {
	defer wg.Done()
	r := strings.NewReplacer(cf.TemplateDir, cf.OutputDir, ".tpl", "")
	out := r.Replace(p)
	secrender.Secrender(p, out, tv.Keys)
}

func init() {
	cf.LoadConfig()
	tv.LoadTemplateVars()
}
