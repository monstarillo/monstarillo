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

// Table represents the structure of a database table, including its name, type, columns, and foreign key relationships.
type Table struct {
	TableName, DatabaseType string
	Columns                 []Column
	ForeignKeys             []ForeignKey
	ReferencedForeignKeys   []ForeignKey
	GuiListTable            GuiListTable
}

// NewTable creates and returns a new Table instance with the specified table name and database type.
func NewTable(tableName, databaseType string) Table {
	var t Table
	t.DatabaseType = databaseType
	t.TableName = tableName
	return t
}

// AddColumn adds a new column to the table by appending it to the Columns slice.
func (t *Table) AddColumn(column Column) {
	t.Columns = append(t.Columns, column)
}

// GetPrimaryColumns returns a slice of Column objects that are marked as primary keys from the table's columns.
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

func (t *Table) GetTableNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, t.TableName)
}

// GetPrimaryNonDateColumns retrieves all primary key columns from the table that are not of "Date" or "Timestamp" type.
func (t *Table) GetPrimaryNonDateColumns() []Column {
	var columns []Column

	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			if t.Columns[c].GetJavascriptDataType() != "Timestamp" && t.Columns[c].GetJavaDataType() != "Date" {
				columns = append(columns, t.Columns[c])
			}
		}
		c++
	}

	return columns
}

// GetPrimaryColumnJavaTypesAndVariables returns a string containing Java types and variables for primary key columns.
// The returned string is formatted as "Type1 var1, Type2 var2" for all primary key columns in the table.
func (t *Table) GetPrimaryColumnJavaTypesAndVariables() string {

	var primary = ""
	c := 0
	var first = true
	for range t.Columns {

		if t.Columns[c].IsPrimaryKey {
			if !first {
				primary += ", "
			}
			primary = primary + t.Columns[c].GetJavaDataType() + " " + t.Columns[c].GetCamelCaseColumnName()
			first = false
		}
		c++
	}

	return primary
}

// GetPrimaryColumnVariables returns a comma-separated string of camel case names for all primary key columns in the table.
func (t *Table) GetPrimaryColumnVariables() string {

	var primary = ""
	c := 0
	var first = true
	for range t.Columns {

		if t.Columns[c].IsPrimaryKey {
			if !first {
				primary += ", "
			}
			primary += t.Columns[c].GetCamelCaseColumnName()
			first = false
		}
		c++
	}

	return primary
}

// GetFirstPrimaryColumn returns the first column marked as a primary key within the table. Returns an empty Column if none found.
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

// GetFirstPrimaryColumnJavaDataType returns the Java data type of the first primary key column in the table. If none, returns empty string.
func (t *Table) GetFirstPrimaryColumnJavaDataType() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetJavaDataType()
		}
		c++
	}
	return ""
}

// GetFirstPrimaryColumnJavaFirstUnitTestValue returns the Java first unit test value of the first primary key column in the table.
func (t *Table) GetFirstPrimaryColumnJavaFirstUnitTestValue() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetJavaFirstUnitTestValue()
		}
		c++
	}
	return ""
}

// GetFirstPrimaryColumnSetString retrieves the set string of the first primary key column in the table, if one exists.
func (t *Table) GetFirstPrimaryColumnSetString() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetSetString()
		}
		c++
	}
	return ""
}

// GetFirstPrimaryColumnJavaSecondUnitTestValue retrieves the second unit test value of the first primary key column in the table.
func (t *Table) GetFirstPrimaryColumnJavaSecondUnitTestValue() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetJavaSecondUnitTestValue()
		}
		c++
	}
	return ""
}

// GetFirstPrimaryColumnCamelCaseColumnName retrieves the camel case name of the first primary key column in the table.
func (t *Table) GetFirstPrimaryColumnCamelCaseColumnName() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetCamelCaseColumnName()
		}
		c++
	}
	return ""
}

// GetFirstPrimaryColumnPascalCaseColumnName returns the PascalCase name of the first primary key column in the table.
func (t *Table) GetFirstPrimaryColumnPascalCaseColumnName() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return t.Columns[c].GetPascalCaseColumnName()
		}
		c++
	}
	return ""
}

// HasCompositePrimaryKey checks if the table has a composite primary key by verifying if it has more than one primary column.
func (t *Table) HasCompositePrimaryKey() bool {
	return len(t.GetPrimaryColumns()) > 1
}

// GetNullableColumns returns a slice of columns that are marked as nullable in the table.
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

// HasJavascriptStringColumn checks if the table contains at least one column with a JavaScript data type of "String".
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

// HasJavascriptNumberColumn checks if the table has at least one column with the JavaScript data type "Number".
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

// HasJavaTypeColumn checks if the table contains a column with the specified Java data type.
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

// HasAutoIncrementColumn checks if the table contains at least one column with the auto-increment property.
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

// HasAnyDateColumn checks if the table contains at least one column with a date-related data type such as date, datetime, year, or timestamp.
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

// HasDateColumn checks if the table contains at least one column with a "date" data type and returns true if found.
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

// HasYearColumn checks if the table contains at least one column with a data type of "year".
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

// HasDateTimeColumn checks if the table contains at least one column with the "datetime" data type. Returns true if found.
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

// HasTimestampColumn checks if the table contains at least one column with the "timestamp" data type. Returns true if found.
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

// GetFkTableNameForColumn returns the primary key table name associated with the given foreign key column name.
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

// GetCamelCaseTableName converts the table's name to camel case format and returns it as a string.
func (t *Table) GetCamelCaseTableName() string {
	return strcase.ToLowerCamel(t.TableName)
}

// GetCamelCaseTableNamePlural converts the table name to camel case and ensures it is pluralized.
func (t *Table) GetCamelCaseTableNamePlural() string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(t.TableName) {
		return strcase.ToLowerCamel(t.TableName)
	}
	return strcase.ToLowerCamel(pluralize.Plural(t.TableName))
}

// GetPascalCaseTableName returns the table name converted to PascalCase format.
func (t *Table) GetPascalCaseTableName() string {
	return strcase.ToCamel(t.TableName)
}

// GetTableName returns the name of the table associated with the Table instance.
func (t *Table) GetTableName() string {
	return t.TableName
}

// GetPascalCaseTableNamePlural returns the table name in PascalCase format, ensuring the name is pluralized.
func (t *Table) GetPascalCaseTableNamePlural() string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(t.TableName) {
		return strcase.ToCamel(t.TableName)
	}
	return strcase.ToCamel(pluralize.Plural(t.TableName))
}

// GetCamelCaseTableNameEF returns the table name in camelCase format, singularized if it is in plural form.
func (t *Table) GetCamelCaseTableNameEF() string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(t.TableName) {
		return strcase.ToLowerCamel(pluralize.Singular(t.TableName))
	}
	return strcase.ToLowerCamel(t.TableName)
}

// GetJavaFirstPrimaryUnitTestValue returns the first unit test value of the primary key column in Java format for the table.
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

// GetFirstPrimarySetString returns the modified set string of the first column marked as a primary key in the table.
func (t *Table) GetFirstPrimarySetString() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return strings.Replace(t.Columns[c].GetSetString(), "L", "", 1)
		}
		c++
	}
	return ""
}

// GetJavaFirstPrimaryUnitTestValueAsString returns the first Java unit test string value of the primary key in the table.
func (t *Table) GetJavaFirstPrimaryUnitTestValueAsString() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return strings.Replace(t.Columns[c].GetJavaFirstUnitTestValue(), "L", "", 1)
		}
		c++
	}
	return ""
}

// GetJavaSecondPrimaryUnitTestValue returns the second primary key's unit test value as a string from the table's columns.
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

// GetJavaSecondPrimaryUnitTestValueAsString returns the second primary unit test value for a column as a modified string.
func (t *Table) GetJavaSecondPrimaryUnitTestValueAsString() string {
	c := 0
	for range t.Columns {
		if t.Columns[c].IsPrimaryKey {
			return strings.Replace(t.Columns[c].GetJavaSecondUnitTestValue(), "L", "", 1)
		}
		c++
	}
	return ""
}

// GetCSharpFirstPrimaryUnitTestValue retrieves the first unit test value of the primary key column from the table.
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

// GetCSharpSecondPrimaryUnitTestValue iterates through table columns to return the second primary key's unit test value in C#.
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

// GetPascalCaseTableNameEF converts the table name to PascalCase format and ensures it's singular if it is plural.
func (t *Table) GetPascalCaseTableNameEF() string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(t.TableName) {
		return strcase.ToCamel(pluralize.Singular(t.TableName))
	}
	return strcase.ToCamel(t.TableName)
}

// GetColumnListWithCSharpTypes generates a formatted string listing column names and their C# data types from the table.
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

// ReadTables reads a JSON file specified by templateFile and unmarshals its content into a slice of Table structs.
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
