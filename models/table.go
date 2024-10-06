package models

import (
	"encoding/json"
	"fmt"
	pluralize "github.com/gertd/go-pluralize"
	"io/ioutil"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

type Table struct {
	TableName, DatabaseType string
	Columns                 []Column
	ForeignKeys             []ForeignKey
	ReferencedForeignKeys   []ForeignKey
	GuiListTable            GuiListTable
}

func NewTable(tableName, databaseType string) Table {
	var t Table
	t.DatabaseType = databaseType
	t.TableName = tableName
	return t
}

func (t *Table) AddColumn(column Column) {
	t.Columns = append(t.Columns, column)
}

func (t *Table) GetPrimaryColumns() []Column {
	var columns []Column

	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			columns = append(columns, t.Columns[c])
		}
		c++
	}

	return columns
}

func (t *Table) GetFirstPrimaryColumn() Column {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c]
		}
		c++
	}
	return Column{}
}
func (t *Table) HasCompositePrimaryKey() bool {
	return len(t.GetPrimaryColumns()) > 1
}

func (t *Table) GetNullableColumns() []Column {
	var columns []Column

	c := 0
	for range t.Columns {
		if t.Columns[c].IsNullable {
			columns = append(columns, t.Columns[c])
		}
		c++
	}
	return columns
}

func (t *Table) HasJavascriptStringColumn() bool {
	response := false

	c := 0
	for range t.Columns {
		if t.Columns[c].GetJavascriptDataType() == "String" {
			response = true
		}
		c++
	}
	return response
}

func (t *Table) HasJavascriptNumberColumn() bool {
	response := false

	c := 0
	for range t.Columns {
		if t.Columns[c].GetJavascriptDataType() == "Number" {
			response = true
		}
		c++
	}
	return response
}

func (t *Table) HasJavaTypeColumn(javaType string) bool {
	response := false

	c := 0
	for range t.Columns {
		if t.Columns[c].GetJavaDataType() == javaType {
			response = true
		}
		c++
	}
	return response
}

func (t *Table) HasAutoIncrementColumn() bool {
	response := false

	c := 0
	for range t.Columns {
		if t.Columns[c].IsAutoIncrement {
			response = true
		}
		c++
	}
	return response
}

func (t *Table) HasAnyDateColumn() bool {
	response := false

	c := 0
	for range t.Columns {
		if t.Columns[c].DataType == "date" {
			response = true
		}
		if t.Columns[c].DataType == "datetime" {
			response = true
		}
		if t.Columns[c].DataType == "year" {
			response = true
		}
		if t.Columns[c].DataType == "timestamp" {
			response = true
		}
		c++
	}
	return response
}

func (t *Table) HasDateColumn() bool {
	response := false

	c := 0
	for range t.Columns {
		if t.Columns[c].DataType == "date" {
			response = true
		}

		c++
	}
	return response
}

func (t *Table) HasYearColumn() bool {
	response := false

	c := 0
	for range t.Columns {

		if t.Columns[c].DataType == "year" {
			response = true
		}

		c++
	}
	return response
}

func (t *Table) HasDateTimeColumn() bool {
	response := false

	c := 0
	for range t.Columns {

		if t.Columns[c].DataType == "datetime" {
			response = true
		}

		c++
	}
	return response
}

func (t *Table) HasTimestampColumn() bool {
	response := false

	c := 0
	for range t.Columns {
		if t.Columns[c].DataType == "timestamp" {
			response = true
		}
		c++
	}
	return response
}

func (t *Table) GetFkTableNameForColumn(columnName string) string {
	response := ""
	c := 0
	for range t.ForeignKeys {
		if t.ForeignKeys[c].FkColumnName == columnName {
			response = t.ForeignKeys[c].PkTableName
			break
		}
		c++
	}

	return response
}

func (t *Table) GetCamelCaseTableName() string {
	return strcase.ToLowerCamel(t.TableName)
}

func (t *Table) GetCamelCaseTableNamePlural() string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(t.TableName) {
		return strcase.ToLowerCamel(t.TableName)
	}
	return strcase.ToLowerCamel(pluralize.Plural(t.TableName))
}

func (t *Table) GetPascalCaseTableName() string {
	return strcase.ToCamel(t.TableName)
}

func (t *Table) GetTableName() string {
	return t.TableName
}

func (t *Table) GetPascalCaseTableNamePlural() string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(t.TableName) {
		return strcase.ToCamel(t.TableName)
	}
	return strcase.ToCamel(pluralize.Plural(t.TableName))
}

func (t *Table) GetCamelCaseTableNameEF() string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(t.TableName) {
		return strcase.ToLowerCamel(pluralize.Singular(t.TableName))
	}
	return strcase.ToLowerCamel(t.TableName)
}

func (t *Table) GetJavaFirstPrimaryUnitTestValue() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetJavaFirstUnitTestValue()
		}
		c++
	}
	return ""
}

func (t *Table) GetJavaSecondPrimaryUnitTestValue() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetJavaSecondUnitTestValue()
		}
		c++
	}
	return ""
}

func (t *Table) GetCSharpFirstPrimaryUnitTestValue() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetCSharpFirstUnitTestValue()
		}
		c++
	}
	return ""
}

func (t *Table) GetCSharpSecondPrimaryUnitTestValue() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetCSharpSecondUnitTestValue()
		}
		c++
	}
	return ""
}

func (t *Table) GetPascalCaseTableNameEF() string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(t.TableName) {
		return strcase.ToCamel(pluralize.Singular(t.TableName))
	}
	return strcase.ToCamel(t.TableName)
}

func (t *Table) GetColumnListWithCSharpTypes() string {
	c := 0
	listing := ""
	for range t.Columns {
		listing = listing + t.Columns[c].GetCSharpDataType() + " " + t.Columns[c].GetCamelCaseColumnName() + ", "
		c++
	}

	lastComma := strings.LastIndex(listing, ",")
	if lastComma > -1 {
		listing = listing[0:lastComma] + listing[lastComma+1:]
	}
	return listing
}

func ReadTables(templateFile string) []Table {
	jsonFile, err := os.Open(templateFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var templates []Table

	err = json.Unmarshal(byteValue, &templates)
	if err != nil {
		return templates
	}

	return templates
}
