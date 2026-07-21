package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GuiListTables is the root of the GUI metadata document produced by
// gui/gui-metadata.tmpl and consumed (via --gui) when generating GUI code.
type GuiListTables struct {
	Tables []GuiListTable `json:"tables"`
}

// GuiListTable is the per-table GUI configuration. Everything here is defaulted
// from the database schema and then hand-editable by the user.
type GuiListTable struct {
	TableName     string         `json:"tableName"`
	Label         string         `json:"label"`
	LabelPlural   string         `json:"labelPlural"`
	Route         string         `json:"route"`
	ApiPath       string         `json:"apiPath"`
	IdFields      []string       `json:"idFields"`
	DisplayColumn string         `json:"displayColumn"`
	ShowInNav     bool           `json:"showInNav"`
	Permissions   GuiPermissions `json:"permissions"`
	DefaultSort   GuiSort        `json:"defaultSort"`
	PageSize      int            `json:"pageSize"`
	Fields        []GuiField     `json:"fields"`
	Views         GuiViews       `json:"views"`
}

type GuiPermissions struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type GuiSort struct {
	Column string `json:"column"`
	Dir    string `json:"dir"`
}

// GuiField is a single form/list field. Control is a semantic widget name
// (text, number, date, datetime, checkbox, textarea, lookup) that the GUI
// templates map to concrete components.
type GuiField struct {
	ColumnName  string     `json:"columnName"`
	Field       string     `json:"field"`
	Label       string     `json:"label"`
	Control     string     `json:"control"`
	Required    bool       `json:"required"`
	ReadOnly    bool       `json:"readOnly"`
	Placeholder *string    `json:"placeholder"`
	Help        *string    `json:"help"`
	Order       int        `json:"order"`
	MaxLength   int        `json:"maxLength,omitempty"`
	Format      *GuiFormat `json:"format,omitempty"`
	Lookup      *GuiLookup `json:"lookup,omitempty"`
}

type GuiFormat struct {
	Type     string `json:"type"`
	Pattern  string `json:"pattern,omitempty"`
	Currency string `json:"currency,omitempty"`
}

// GuiLookup describes how a foreign-key field is rendered and resolved.
// Strategy is one of: dropdown, popup, autocomplete.
type GuiLookup struct {
	TargetTable   string    `json:"targetTable"`
	Strategy      string    `json:"strategy"`
	ValueColumn   string    `json:"valueColumn"`
	DisplayColumn string    `json:"displayColumn"`
	SearchColumns []string  `json:"searchColumns"`
	Multiple      bool      `json:"multiple"`
	AllowCreate   bool      `json:"allowCreate"`
	Popup         *GuiPopup `json:"popup,omitempty"`
}

type GuiPopup struct {
	Columns  []string `json:"columns"`
	PageSize int      `json:"pageSize"`
}

type GuiViews struct {
	List   GuiListView   `json:"list"`
	Form   GuiFormView   `json:"form"`
	Detail GuiDetailView `json:"detail"`
}

type GuiListView struct {
	Columns    []string `json:"columns"`
	LinkColumn string   `json:"linkColumn"`
	Searchable []string `json:"searchable"`
	Sortable   []string `json:"sortable"`
}

type GuiFormView struct {
	Columns  []string     `json:"columns"`
	Sections []GuiSection `json:"sections"`
}

type GuiSection struct {
	Title   string   `json:"title"`
	Columns []string `json:"columns"`
}

type GuiDetailView struct {
	Columns     []string        `json:"columns"`
	ChildTables []GuiChildTable `json:"childTables"`
}

type GuiChildTable struct {
	Table    string `json:"table"`
	FkColumn string `json:"fkColumn"`
	Label    string `json:"label"`
	As       string `json:"as"`
}

// GetField returns the field config matching a column name or camelCase field
// name; returns an empty GuiField if not found.
func (t *GuiListTable) GetField(name string) GuiField {
	for i := range t.Fields {
		if t.Fields[i].ColumnName == name || t.Fields[i].Field == name {
			return t.Fields[i]
		}
	}
	return GuiField{}
}

// HasControl reports whether any field uses the given control.
func (t *GuiListTable) HasControl(control string) bool {
	for i := range t.Fields {
		if strings.Contains(t.Fields[i].Control, control) {
			return true
		}
	}
	return false
}

// --- Backwards-compatible helpers used by the funcMap and older templates. ---

func (t *GuiListTable) HasGuiControl(controlName string) bool {
	return t.HasControl(controlName)
}

func (t *GuiListTable) GetEditGuiControlForColumn(columnName string) string {
	return t.GetField(columnName).Control
}

func (t *GuiListTable) GetEditSelectKeyForColumn(columnName string) string {
	if f := t.GetField(columnName); f.Lookup != nil {
		return f.Lookup.ValueColumn
	}
	return ""
}

func (t *GuiListTable) GetEditSelectValueForColumn(columnName string) string {
	if f := t.GetField(columnName); f.Lookup != nil {
		return f.Lookup.DisplayColumn
	}
	return ""
}

func ReadGuiTables(templateFile string) GuiListTables {
	jsonFile, err := os.Open(templateFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tables GuiListTables

	err = json.Unmarshal(byteValue, &tables)
	if err != nil {
		return GuiListTables{}
	}

	return tables
}
