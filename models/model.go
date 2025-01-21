package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Model struct {
	TableName, ModelName, Orm string
	ModelColumns              []ModelColumn
}

type ModelColumn struct {
	ColumnName, PropertyName, OrmType string
	IsReadOnly                        bool
}

func GetModelNameForTable(dbModels []Model, tableName string) string {

	response := ""

	c := 0
	for range dbModels {
		if dbModels[c].TableName == tableName {
			response = dbModels[c].ModelName
			break
		}
		c++
	}

	return response
}

func GetOrmForTable(dbModels []Model, tableName string) string {

	response := ""

	c := 0
	for range dbModels {
		if dbModels[c].TableName == tableName {
			response = dbModels[c].Orm
			break
		}
		c++
	}

	return response
}

func GetPropertyNameForModelColumn(dbModels []Model, tableName, columnName string) string {

	response := ""

	mod := 0
	for range dbModels {
		if dbModels[mod].TableName == tableName {
			col := 0
			for range dbModels[mod].ModelColumns {
				if dbModels[mod].ModelColumns[col].ColumnName == columnName {
					response = dbModels[mod].ModelColumns[col].PropertyName
					break
				}
				col++
			}
		}
		mod++
	}

	return response
}

func GetOrmTypeForModelColumn(dbModels []Model, tableName, columnName string) string {

	response := ""

	mod := 0
	for range dbModels {
		if dbModels[mod].TableName == tableName {
			col := 0
			for range dbModels[mod].ModelColumns {
				if dbModels[mod].ModelColumns[col].ColumnName == columnName {
					response = dbModels[mod].ModelColumns[col].OrmType
					break
				}
				col++
			}
		}
		mod++
	}

	return response
}

func ReadModels(modelFile string) []Model {
	jsonFile, err := os.Open(modelFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var models []Model

	err = json.Unmarshal(byteValue, &models)
	if err != nil {
		return []Model{}
	}

	return models
}
