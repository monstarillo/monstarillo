package models

import (
	"github.com/iancoleman/strcase"
	"strings"
)

type ForeignKey struct {
	ConstraintName, FkTableName, FkColumnName, PkTableName, PkColumnName, Relation string
	FkColumn, PkColumn                                                             Column
}

func NewForeignKey(constraintName, fkTableName, fkColumnName,
	pkTableName, pkColumnName string) ForeignKey {
	var fk ForeignKey
	fk.ConstraintName = constraintName
	fk.FkTableName = fkTableName
	fk.FkColumnName = fkColumnName
	fk.PkTableName = pkTableName
	fk.PkColumnName = pkColumnName

	return fk
}

func (f *ForeignKey) GetFkTableNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, f.FkTableName)
}

func (f *ForeignKey) GetFkColumnNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, f.FkColumnName)
}

func (f *ForeignKey) GetPkColumnNameInCase(caseToReturn string) string {
	return getCaseValue(caseToReturn, f.PkColumnName)
}

func (f *ForeignKey) GetCamelCaseFKTableName() string {
	return strcase.ToLowerCamel(f.FkTableName)
}

func (f *ForeignKey) GetPascalCaseFKTableName() string {
	return strcase.ToCamel(f.FkTableName)
}

func (f *ForeignKey) GetCamelCaseFKColumnName() string {
	return strcase.ToLowerCamel(f.FkColumnName)
}

func (f *ForeignKey) GetPascalCaseFKColumnName() string {
	return strcase.ToCamel(f.FkColumnName)
}

func (f *ForeignKey) GetCamelCasePKTableName() string {
	return strcase.ToLowerCamel(f.PkTableName)
}

func (f *ForeignKey) GetPascalCasePKTableName() string {
	return strcase.ToCamel(f.PkTableName)
}

func (f *ForeignKey) GetCamelCasePKColumnName() string {
	return strcase.ToLowerCamel(f.PkColumnName)
}

func (f *ForeignKey) GetPascalCasePKColumnName() string {
	return strcase.ToCamel(f.PkColumnName)
}

func (f *ForeignKey) GetCamelCaseFKTableNamePlural() string {
	if strings.HasSuffix(f.FkTableName, "s") {
		return strcase.ToLowerCamel(f.FkTableName) + "es"
	} else {
		return strcase.ToLowerCamel(f.FkTableName) + "s"
	}
}

func (f *ForeignKey) GetFKTableNamePluralInCase() string {
	if strings.HasSuffix(f.FkTableName, "s") {
		return strcase.ToCamel(f.FkTableName) + "es"
	} else {
		return strcase.ToCamel(f.FkTableName) + "s"
	}
}

func (f *ForeignKey) GetPascalCaseFKTableNamePlural(caseToReturn string) string {
	if strings.HasSuffix(f.FkTableName, "s") {
		return getCaseValue(caseToReturn, f.FkTableName) + "es"
	} else {
		return strcase.ToCamel(f.FkTableName) + "s"
	}
}

func (f *ForeignKey) GetCamelCasePKTableNamePlural() string {
	if strings.HasSuffix(f.PkTableName, "s") {
		return strcase.ToLowerCamel(f.PkTableName) + "es"
	} else {
		return strcase.ToLowerCamel(f.PkTableName) + "s"
	}
}

func (f *ForeignKey) GetPascalCasePKTableNamePlural() string {
	if strings.HasSuffix(f.PkTableName, "s") {
		return strcase.ToCamel(f.PkTableName) + "es"
	} else {
		return strcase.ToCamel(f.PkTableName) + "s"
	}
}
