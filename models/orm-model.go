package models

import (
	"encoding/json"
	"fmt"
	"github.com/gertd/go-pluralize"
	"io"
	"os"
)

// OrmModel represents a database model with metadata for ORM transformations, including table and column details.
type OrmModel struct {
	TableName, ModelName, Orm string
	Columns                   []OrmColumn
}

// GetModelNameInCase returns the ModelName of the OrmModel converted to the specified case format.
func (m *OrmModel) GetModelNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, m.ModelName)
}

// GetModelTableNameInCase returns the TableName of the OrmModel converted to the specified case format.
func (m *OrmModel) GetModelTableNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, m.TableName)
}

// ReadModels reads a JSON file containing ORM model definitions and unmarshals it into a slice of OrmModel.
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

func (m *OrmModel) GetPrimaryModelColumns() []OrmColumn {
	var columns []OrmColumn

	c := 0
	for range m.Columns {
		if m.Columns[c].IsPrimaryKey {
			columns = append(columns, m.Columns[c])
		}
		c++
	}

	return columns
}

// GetModelNamePluralInCase ensures model name is pluralized in the case that is passed.
func (m *OrmModel) GetModelNamePluralInCase(caseToReturn string) string {
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(m.ModelName) {
		return getCaseValue(caseToReturn, m.ModelName)
	}
	return getCaseValue(caseToReturn, pluralize.Plural(m.ModelName))
}

func (m *OrmModel) HasCompositePrimaryKey() bool {
	return len(m.GetPrimaryModelColumns()) > 1
}
