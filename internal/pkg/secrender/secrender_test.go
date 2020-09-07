package secrender

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}()

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
	createOutputPath()
	_, err := os.Stat(fd.OutputPath)
	if os.IsNotExist(err) {
		t.Errorf("Failed to create output dir: %s", fd.OutputPath)
	}
}

func TestRenderFile(t *testing.T) {
	fd.TemplatePath = "template_test/templates/test.md.tpl"
	fd.OutputPath = "template_test/render_test/test.md"
	renderFile()

	_, err := os.Stat("template_test/render_test/test.md")
	if err != nil {
		t.Errorf("Expected file, got %v", err)
	}
}

func TestSecrender(t *testing.T) {
	fd.OutputPath = "template_test/output/test.md"
	testVars := map[string]interface{}{
		"test": map[string]interface{}{
			"name": "Secrender Test",
		},
	}
	Secrender(fd.TemplatePath, fd.OutputPath, testVars)
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
	os.MkdirAll("template_test/templates", 0777)
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
}
