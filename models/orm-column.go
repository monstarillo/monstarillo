package models

import "github.com/iancoleman/strcase"

type OrmColumn struct {
	ColumnName, OrmType, DatabaseType, PropertyName string
	IsPrimaryKey, IsNullable, IsAutoIncrement       bool
}

func (c *OrmColumn) GetPropertyNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, c.PropertyName)
}

func (c *OrmColumn) GetColumnNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, c.ColumnName)
}

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
