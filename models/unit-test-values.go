package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UnitTestValueTables struct {
	Tables []UnitTestValueTable `json:"tables"`
}
type UnitTestValueTable struct {
	TableName string                `json:"tableName"`
	Columns   []UnitTestValueColumn `json:"columns"`
}

type UnitTestValueColumn struct {
	ColumnName     string          `json:"columnName"`
	UnitTestValues []UnitTestValue `json:"unitTestValues"`
}

type UnitTestValue struct {
	FirstUnitTestValue  string `json:"firstUnitTestValue"`
	SecondUnitTestValue string `json:"secondUnitTestValue"`
}

func ReadUnitTestValues(unitTestFile string) UnitTestValueTables {
	jsonFile, err := os.Open(unitTestFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var templates UnitTestValueTables

	err = json.Unmarshal(byteValue, &templates)
	if err != nil {
		return UnitTestValueTables{}
	}

	return templates
}
