package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type OrmModel struct {
	TableName, ModelName, Orm string
	Columns                   []OrmColumn
}

func (m *OrmModel) GetModelNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, m.ModelName)
}

func (m *OrmModel) GetModelTableNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, m.TableName)
}

func ReadModels(modelFile string) []OrmModel {
	jsonFile, err := os.Open(modelFile)
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	var models []OrmModel

	err = json.Unmarshal(byteValue, &models)
	if err != nil {
		return []OrmModel{}
	}

	return models
}
