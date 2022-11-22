package models

import (
	"testing"
)

func TestNewTemplate(t *testing.T) {
	value := NewTemplate("Template", "Foo.txt", "code", "path")
	if value.TemplateFile != "Template" {
		t.Errorf("Value was incorrect, got %s, wanted Template", value.TemplateFile)
	}

	if value.GeneratedFileName != "Foo.txt" {
		t.Errorf("Value was incorrect, got %s, wanted Foo.txt", value.GeneratedFileName)
	}
}

func TestReadTemplates(t *testing.T) {
	value := ReadTemplates("../test-data/test.json")

	if len(value.Templates) != 1 {
		t.Errorf("Value was incorrect, got %d, wanted 1", len(value.Templates))
	}
}
