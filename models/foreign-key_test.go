package models

import (
	"testing"
)

func TestForeignKey_GetCamelCaseFKTableName(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetCamelCaseFKTableName()

	if value != "film" {
		t.Errorf("Value was incorrect, got %s, want film", value)
	}
}

func TestForeignKey_GetPascalCaseFKTableName(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetPascalCaseFKTableName()

	if value != "Film" {
		t.Errorf("Value was incorrect, got %s, want Film", value)
	}
}
func TestForeignKey_GetCamelCaseFKTableNamePlural(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetCamelCaseFKTableNamePlural()

	if value != "films" {
		t.Errorf("Value was incorrect, got %s, want film", value)
	}
}

func TestForeignKey_GetPascalCaseFKTableNamePlural(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetPascalCaseFKTableNamePlural()

	if value != "Films" {
		t.Errorf("Value was incorrect, got %s, want Films", value)
	}
}

func TestForeignKey_GetCamelCasePKTableName(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetCamelCasePKTableName()

	if value != "language" {
		t.Errorf("Value was incorrect, got %s, want language", value)
	}
}

func TestForeignKey_GetPascalCasePKTableNameName(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetPascalCasePKTableName()

	if value != "Language" {
		t.Errorf("Value was incorrect, got %s, want Language", value)
	}
}
func TestForeignKey_GetCamelCasePKTableNamePlural(t *testing.T) {
	table, _ := getTable("film")
	PK := table.ForeignKeys[0]

	value := PK.GetCamelCasePKTableNamePlural()

	if value != "languages" {
		t.Errorf("Value was incorrect, got %s, want languages", value)
	}
}

func TestForeignKey_GetPascalCasePKTableNamePlural(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetPascalCasePKTableNamePlural()

	if value != "Languages" {
		t.Errorf("Value was incorrect, got %s, want Languages", value)
	}
}

func TestForeignKey_GetCamelCaseFKColumnName(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetCamelCaseFKColumnName()

	if value != "languageId" {
		t.Errorf("Value was incorrect, got %s, want languageId", value)
	}
}

func TestForeignKey_GetPascalCaseFKColumnName(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetPascalCaseFKColumnName()

	if value != "LanguageId" {
		t.Errorf("Value was incorrect, got %s, want LanguageId", value)
	}
}

func TestForeignKey_GetCamelCasePKColumnName(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetCamelCasePKColumnName()

	if value != "languageId" {
		t.Errorf("Value was incorrect, got %s, want languageId", value)
	}
}
func TestForeignKey_GetPascalCasePKColumnName(t *testing.T) {
	table, _ := getTable("film")
	fk := table.ForeignKeys[0]

	value := fk.GetPascalCasePKColumnName()

	if value != "LanguageId" {
		t.Errorf("Value was incorrect, got %s, want LanguageId", value)
	}
}

func TestNewForeignKey(t *testing.T) {
	fk := NewForeignKey("constraintName", "person", "color_id", "color", "id")

	value := fk.ConstraintName

	if value != "constraintName" {
		t.Errorf("Value was incorrect, got %s, want constraintName", value)
	}
}
