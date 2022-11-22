package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type GuiList struct {
	FirstColumn GuiListColumn   `json:"firstColumn"`
	Columns     []GuiListColumn `json:"columns"`
}

type GuiEdit struct {
	Columns []GuiListColumn `json:"columns"`
}

type GuiListColumn struct {
	ColumnName      string `json:"columnName"`
	Title           string `json:"title"`
	Value           string `json:"value"`
	GuiControl      string `json:"guiControl"`
	SelectOptions   string `json:"selectOptions"`
	SelectKey       string `json:"selectKey"`
	SelectValue     string `json:"selectValue"`
	SelectTableName string `json:"selectTableName"`
}

type GuiListTable struct {
	TableName          string  `json:"tableName"`
	AllowAttachedFiles bool    `json:"allowAttachedFiles"`
	GuiList            GuiList `json:"list"`
	GuiEdit            GuiEdit `json:"edit"`
}

type GuiListTables struct {
	Tables []GuiListTable `json:"tables"`
}

func (t *GuiListTable) HasGuiControl(controlName string) bool {
	c := 0
	for range t.GuiEdit.Columns {
		if strings.Contains(t.GuiEdit.Columns[c].GuiControl, controlName) {
			return true
		}
		c++
	}
	return false
}

func (t *GuiListTable) GetEditGuiControlForColumn(columnName string) string {
	response := ""

	c := 0
	for range t.GuiEdit.Columns {
		if t.GuiEdit.Columns[c].Value == columnName {
			response = t.GuiEdit.Columns[c].GuiControl
			break
		}
		c++
	}

	return response
}

func (t *GuiListTable) GetEditSelectKeyForColumn(columnName string) string {
	response := ""

	c := 0
	for range t.GuiEdit.Columns {
		if t.GuiEdit.Columns[c].Value == columnName {
			response = t.GuiEdit.Columns[c].SelectKey
			break
		}
		c++
	}

	return response
}

func (t *GuiListTable) GetEditSelectValueForColumn(columnName string) string {
	response := ""

	c := 0
	for range t.GuiEdit.Columns {
		if t.GuiEdit.Columns[c].Value == columnName {
			response = t.GuiEdit.Columns[c].SelectValue
			break
		}
		c++
	}

	return response
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
