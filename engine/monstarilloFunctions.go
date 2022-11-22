package engine

import (
	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/monstarillo/monstarillo/models"
	"os"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"getTag":                        getTag,
	"ToUpper":                       strings.ToUpper,
	"ToLower":                       strings.ToLower,
	"getTable":                      getTable,
	"getTableFirstPk":               getTableFirstPk,
	"getTableSecondColumn":          getTableSecondColumn,
	"makePlural":                    makePlural,
	"makePascalCase":                makePascalCase,
	"makeCamelCase":                 makeCamelCase,
	"GetFkTableNamePlural":          GetFkTableNamePlural,
	"GetEditSelectKeyForColumn":     GetEditSelectKeyForColumn,
	"GetEditSelectValueForColumn":   GetEditSelectValueForColumn,
	"getColumnCountByDataType":      getColumnCountByDataType,
	"GetGoParseIntConversionSuffix": GetGoParseIntConversionSuffix,
	"GoIntCast":                     GoIntCast,
}

func getTable(tableName string) models.Table {
	tables := models.ReadTables("tables.json")
	for _, t := range tables {
		if t.TableName == tableName {
			return t
		}
	}

	var tbl models.Table
	return tbl
}

func getTableFirstPk(tableName string) string {
	tables := models.ReadTables("tables.json")
	for _, t := range tables {
		if t.TableName == tableName {
			return t.GetPrimaryColumns()[0].ColumnName
		}
	}

	return ""
}

func getTableSecondColumn(tableName string) string {
	tables := models.ReadTables("tables.json")

	pkColumn := getTableFirstPk(tableName)
	for _, t := range tables {
		if t.TableName == tableName {
			cols := t.Columns

			if len(cols) == 1 {
				return cols[0].ColumnName
			}

			if len(pkColumn) == 0 {
				return cols[1].ColumnName
			} else if cols[0].ColumnName == pkColumn {
				return cols[1].ColumnName
			} else {
				return cols[0].ColumnName
			}
		}
	}

	return ""
}

func getColumnCountByDataType(tableName, dataType string) int {
	count := 0
	table := getTable(tableName)
	for _, c := range table.Columns {
		if c.DataType == dataType {
			count++
		}
	}
	return count
}

func makePlural(s string) string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(s) {
		return s
	}
	return pluralize.Plural(s)
}

func makeCamelCase(s string) string {
	return strcase.ToLowerCamel(s)
}

func makePascalCase(s string) string {
	return strcase.ToCamel(s)
}

func getTag(tags []models.Tag, tag string) string {
	t := 0
	for range tags {
		if tags[t].TagName == tag {
			return tags[t].Value
		}
		t++

	}
	return ""
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func deleteFile(filename string) {
	if fileExists(filename) {
		e := os.Remove(filename)
		check(e)
	}
}

func GetFkTableNamePlural(tableName, columnName string) string {

	t := getTable(tableName)

	for _, fk := range t.ForeignKeys {
		if fk.FkColumnName == columnName {
			pluralize := pluralize.NewClient()
			if pluralize.IsPlural(fk.PkTableName) {
				return strcase.ToLowerCamel(t.TableName)
			}
			return strcase.ToLowerCamel(pluralize.Plural(fk.PkTableName))
		}
	}

	return ""
}

func GetEditSelectKeyForColumn(tableName, columnName string) string {
	context := ReadMonstrilloContext("context.json")

	for _, t := range context.GuiListTables.Tables {
		if t.TableName == tableName {
			return t.GetEditSelectKeyForColumn(columnName)
		}
	}
	return ""
}

func GetEditSelectValueForColumn(tableName, columnName string) string {
	context := ReadMonstrilloContext("context.json")

	for _, t := range context.GuiListTables.Tables {
		if t.TableName == tableName {
			return t.GetEditSelectValueForColumn(columnName)
		}
	}
	return ""
}
func GetMonstarilloContext(file string) MonstarilloContext {
	return ReadMonstrilloContext(file)
}

func GetGoParseIntConversionSuffix(datatype string) string {
	switch datatype {
	case "uint8":
		fallthrough
	case "int8":

		return ", 10, 8)"

	case "uint16":
		fallthrough
	case "int16":
		return ", 10, 16)"
	case "uint32":
		fallthrough
	case "int32":
		return ", 10, 32)"
	case "uint64":
		fallthrough
	case "int":
		fallthrough
	case "int64":
		return ", 10, 64)"

	}
	return "Datatype not supported yet :("
}

func GoIntCast(parameterName, datatype string) string {
	switch datatype {
	case "uint8":
		fallthrough
	case "int8":

		return "int8(" + parameterName + ")"

	case "uint16":
		fallthrough
	case "int16":
		return "int16(" + parameterName + ")"
	case "uint32":
		fallthrough
	case "int32":
		return "int32(" + parameterName + ")"
	case "uint64":
		fallthrough

	case "int64":
		return "int64(" + parameterName + ")"
	case "int":
		return "int(" + parameterName + ")"

	}
	return parameterName
}
