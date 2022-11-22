package models

import (
	"testing"
)

func TestNewTag(t *testing.T) {
	value := NewTag("Name", "Foo")
	if value.TagName != "Name" {
		t.Errorf("Value was incorrect, got %s, wanted Name", value)
	}

	if value.Value != "Foo" {
		t.Errorf("Value was incorrect, got %s, wanted Foo", value)
	}
}
