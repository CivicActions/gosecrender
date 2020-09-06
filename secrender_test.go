package secrender

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func TestLoadTemplateVars(t *testing.T) {
	tv.loadTemplateVars()
	if len(tv.Keys) != 1 {
		t.Errorf("Expected 1 key, got %d", len(tv.Keys))
	}
}

func TestParseJSON(t *testing.T) {
	tv.parseJSON("template_test/keys/test.json")
	_, err := json.Marshal(tv.Keys["test"])
	if err != nil {
		t.Errorf("Unable to marshal JSON: %v", err)
	}
}

func TestGetFiles(t *testing.T) {
	files := getFiles("template_test/keys")
	got := len(files)
	if got != 1 {
		t.Errorf("Expected 1 file, got %d", got)
	}
}

func TestIsTemplate(t *testing.T) {
	isTpl := isTemplate("template_test/templates/test.md.tpl")
	if !isTpl {
		t.Errorf("isTemplate: Expected true, got %t", isTpl)
	}
	isntTpl := isTemplate("template_test/keys/test.json")
	if isntTpl {
		t.Errorf("isTemplate: Expected false, got %t", isTpl)
	}
}

func TestCreateOutput(t *testing.T) {
	fd.OutputPath = "template_test/output/"
	createOutput()
	_, err := os.Stat(fd.OutputPath)
	if os.IsNotExist(err) {
		t.Errorf("Failed to create output dir: %s", fd.OutputPath)
	}
}

func TestLoadParams(t *testing.T) {
	if fd.KeyDir != "template_test/keys/" {
		t.Errorf("Expected template_test/keys/, got %s", fd.KeyDir)
	}
}

func TestRenderFile(t *testing.T) {
	fd.TemplatePath = "template_test/templates/test.md.tpl"
	fd.OutputPath = "template_test/render_test/test.md"
	wg.Add(1)
	renderFile()
	_, err := os.Stat("template_test/render_test/test.md")
	if err != nil {
		t.Errorf("Expected file, got %v", err)
	}
}

func TestSecrender(t *testing.T) {
	fd.OutputPath = "template_test/output/test.md"
	Secrender(fd.TemplatePath, fd.OutputPath, fd.KeyDir)
	_, err := os.Stat(fd.OutputPath)
	if err != nil {
		t.Errorf("Failed to create file %s", fd.OutputPath)
	}
	t.Cleanup(func() {
		fmt.Println("Removing test files")
		os.RemoveAll("template_test")
	})
}

func init() {
	os.MkdirAll("template_test/keys", 0777)
	k, err := os.Create("template_test/keys/test.json")
	if err != nil {
		log.Fatal("Could not create test template.")
	}
	len, err := k.WriteString("{\"test\": {\"name\": \"Secrender Test\"}}")
	if err != nil {
		log.Fatalf("Could not create file test.json. Bytes: %d", len)
	}
	k.Close()

	os.Mkdir("template_test/templates", 0777)
	f, err := os.Create("template_test/templates/test.md.tpl")
	if err != nil {
		log.Fatal("Could not create test template.")
	}
	len, er := f.WriteString("# {{.Keys.test.name}} test template.")
	if er != nil {
		log.Fatalf("Could not create file test.md.tpl. Bytes: %d", len)
	}
	f.Close()

	fd.TemplatePath = "template_test/templates/test.md.tpl"
	fd.OutputPath = "template_test/output/test.md"
	fd.KeyDir = "template_test/keys/"
}
