package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ForeignKeyDisplay struct {
	ForeignKeyColumn string `json:"foreignKeyColumn"`
	ReferencedTable  string `json:"referencedTable"`
	DisplayField     string `json:"displayField"`
	DtoFieldName     string `json:"dtoFieldName"`
	DtoFieldType     string `json:"dtoFieldType"`
	ShowInList       bool   `json:"showInList"`
	ShowInDetail     bool   `json:"showInDetail"`
}

type SearchConfig struct {
	IsSearchable              bool   `json:"isSearchable"`
	SearchField               string `json:"searchField"`
	SearchFieldJavascriptType string `json:"searchFieldJavascriptType"`
	SearchFieldJavaType       string `json:"searchFieldJavaType"`
	SearchEndpoint            string `json:"searchEndpoint"`
}

type ServiceConfig struct {
	UseJoinedQueries bool `json:"useJoinedQueries"`
	EnableSearch     bool `json:"enableSearch"`
}

type GuiList struct {
	FirstColumn GuiListColumn   `json:"firstColumn"`
	Columns     []GuiListColumn `json:"columns"`
}

type GuiEdit struct {
	Columns []GuiListColumn `json:"columns"`
}

type GuiListColumn struct {
	ColumnName          string `json:"columnName"`
	Title               string `json:"title"`
	Value               string `json:"value"`
	GuiControl          string `json:"guiControl"`
	SelectOptions       string `json:"selectOptions"`
	SelectKey           string `json:"selectKey"`
	SelectValue         string `json:"selectValue"`
	SelectTableName     string `json:"selectTableName"`
	IsForeignKeyDisplay bool   `json:"isForeignKeyDisplay,omitempty"`
	IsNullable          bool   `json:"isNullable"`
}

type GuiListTable struct {
	TableName          string              `json:"tableName"`
	AllowAttachedFiles bool                `json:"allowAttachedFiles"`
	ForeignKeyDisplays []ForeignKeyDisplay `json:"foreignKeyDisplays"`
	SearchConfig       *SearchConfig       `json:"searchConfig"`
	ServiceConfig      ServiceConfig       `json:"serviceConfig"`
	GuiList            GuiList             `json:"list"`
	GuiEdit            GuiEdit             `json:"edit"`
	GuiView            GuiEdit             `json:"view"`
	GuiCreate          GuiEdit             `json:"create"`
}

type Endpoints struct {
	GetAll  string `json:"getAll"`
	GetById string `json:"getById"`
	Create  string `json:"create"`
	Update  string `json:"update"`
	Delete  string `json:"delete"`
	Search  string `json:"search"`
}

type ApiConfig struct {
	BaseUrl   string    `json:"baseUrl"`
	Endpoints Endpoints `json:"endpoints"`
}

type RouteNames struct {
	List   string `json:"list"`
	View   string `json:"view"`
	Edit   string `json:"edit"`
	Create string `json:"create"`
}

type RouteConfig struct {
	BasePath   string     `json:"basePath"`
	RouteNames RouteNames `json:"routeNames"`
}

type ComponentConfig struct {
	BasePath string            `json:"basePath"`
	Imports  map[string]string `json:"imports"`
}

type DataTypeConfig struct {
	Date       string `json:"date"`
	DateTime   string `json:"datetime"`
	Boolean    string `json:"boolean"`
	Text       string `json:"text"`
	ForeignKey string `json:"foreignKey"`
}

type GuiListTables struct {
	Tables          []GuiListTable  `json:"tables"`
	ApiConfig       ApiConfig       `json:"apiConfig"`
	RouteConfig     RouteConfig     `json:"routeConfig"`
	ComponentConfig ComponentConfig `json:"componentConfig"`
	DataTypeConfig  DataTypeConfig  `json:"dataTypeConfig"`
}

func (t *GuiListTable) HasEditGuiControl(controlName string) bool {
	c := 0
	for range t.GuiEdit.Columns {
		if strings.Contains(t.GuiEdit.Columns[c].GuiControl, controlName) {
			return true
		}
		c++
	}
	return false
}

func (t *GuiListTable) HasCreateGuiControl(controlName string) bool {
	c := 0
	for range t.GuiCreate.Columns {
		if strings.Contains(t.GuiCreate.Columns[c].GuiControl, controlName) {
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
