package models

import "github.com/iancoleman/strcase"

// OrmColumn represents a database column with metadata for ORM transformations.
// It includes details such as column name, data types, property name, and constraints.
type OrmColumn struct {
	ColumnName, OrmType, DatabaseType, PropertyName string
	IsPrimaryKey, IsNullable, IsAutoIncrement       bool
}

// GetPropertyNameInCase returns the PropertyName of the OrmColumn converted to the specified case format.
func (c *OrmColumn) GetPropertyNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, c.PropertyName)
}

// GetColumnNameInCase returns the column name formatted in the specified case style (e.g., camel, snake, kebab, pascal).
func (c *OrmColumn) GetColumnNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, c.ColumnName)
}

// getCaseValue transforms a string to a specified case format such as pascal, camel, kebab, or snake based on the input.
func getCaseValue(caseToReturn, value string) string {
	switch caseToReturn {
	case "pascal":
		return strcase.ToCamel(value)

	case "camel":
		return strcase.ToLowerCamel(value)
	case "kebab":
		return strcase.ToKebab(value)

	case "snake":
		return strcase.ToSnake(value)

	}
	return strcase.ToCamel(value)
}
