package engine

import (
	"encoding/json"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/monstarillo/monstarillo/models"
	"io/ioutil"
	"os"
)

type MonstarilloContext struct {
	Tables             []models.Table
	Tags               []models.Tag
	CurrentTable       models.Table
	CurrentGuiTable    models.GuiListTable
	UnitTestValuesFile string
	GuiListTables      models.GuiListTables
}

func (m *MonstarilloContext) GetColumn(tableName, columnName string) models.Column {

	for _, t := range m.Tables {
		if t.TableName == tableName {
			for _, c := range t.Columns {
				if c.ColumnName == columnName {
					return c
				}
			}
		}
	}

	var col models.Column
	return col
}

func (m *MonstarilloContext) GetTable(tableName string) models.Table {

	for _, t := range m.Tables {
		if t.TableName == tableName {
			return t
		}
	}

	var tbl models.Table
	return tbl
}

func (m *MonstarilloContext) GetFkTableName(tableName, columnName string) string {

	for _, t := range m.Tables {
		if t.TableName == tableName {
			for _, fk := range t.ForeignKeys {
				if fk.FkColumnName == columnName {
					return fk.PkTableName
				}
			}
		}
	}

	return ""
}

func (m *MonstarilloContext) GetFkTableNamePlural(tableName, columnName string) string {

	for _, t := range m.Tables {
		if t.TableName == tableName {
			for _, fk := range t.ForeignKeys {
				if fk.FkColumnName == columnName {
					pluralize := pluralize.NewClient()
					if pluralize.IsPlural(fk.PkTableName) {
						return strcase.ToLowerCamel(t.TableName)
					}
					return strcase.ToLowerCamel(pluralize.Plural(fk.PkTableName))
				}
			}
		}
	}

	return ""
}

func ReadMonstrilloContext(contextFile string) MonstarilloContext {
	jsonFile, err := os.Open(contextFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var ctx MonstarilloContext

	err = json.Unmarshal(byteValue, &ctx)
	if err != nil {
		return ctx
	}

	return ctx
}
