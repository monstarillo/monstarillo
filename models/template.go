package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Template struct {
	TemplateFile        string `json:"templateFile"`
	GeneratedFileName   string `json:"generatedFileName"`
	GeneratedFolderName string `json:"generatedFolderName"`
	OutputPath          string `json:"outputPath"`

	AppendFile         bool   `json:"appendFile"`
	OverwriteFile      bool   `json:"overwriteFile"`
	UnitTestJsonJava   string `json:"unitTestJsonJava"`
	UnitTestJsonCSharp string `json:"UnitTestJsonCSharp"`
	ForeignKeyMapping  string `json:"foreignKeyMapping"`

	MinimumGeneratedFileLength int `json:"minimumGeneratedFileLength"`
}

type Templates struct {
	Templates     []Template `json:"templates"`
	Tags          []Tag      `json:"tags"`
	IgnoreTables  []string   `json:"ignoreTables"`
	IncludeTables []string   `json:"includeTables"`
}

func NewTemplate(templateFile, generatedFileName, generatedFolderName, outputPath string) Template {
	var t Template
	t.TemplateFile = templateFile
	t.GeneratedFileName = generatedFileName
	t.GeneratedFolderName = generatedFolderName
	t.OutputPath = outputPath
	t.MinimumGeneratedFileLength = 0
	t.UnitTestJsonJava = ""
	t.UnitTestJsonJava = ""
	t.ForeignKeyMapping = ""
	//t.AppendFile = false
	//t.OverwriteFile = true

	return t
}

func ReadTemplates(templateFile string) Templates {
	jsonFile, err := os.Open(templateFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var templates Templates

	err = json.Unmarshal(byteValue, &templates)
	if err != nil {
		return Templates{}
	}

	return templates
}
