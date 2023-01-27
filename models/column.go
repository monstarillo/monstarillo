package models

import (
	strcase "github.com/iancoleman/strcase"
	"unicode"
	//"monstarillo/mysql"
)

type Column struct {
	ColumnName, DataType, DatabaseType, TableName, PkTableName, PkColumnName string
	IsPrimaryKey, IsNullable, IsAutoIncrement, IsForeignKey                  bool
	OrdinalPosition, NumericPrecision, NumericScale, CharacterMaximumLength  int
}

func NewColumn(columnName, databaseType, tableName string) Column {
	var c Column
	c.ColumnName = columnName
	c.DatabaseType = databaseType
	c.TableName = tableName
	c.DataType = ""
	c.IsPrimaryKey = false
	c.OrdinalPosition = 0
	c.NumericScale = 0
	c.NumericPrecision = 0
	c.CharacterMaximumLength = 0
	c.IsNullable = false
	c.IsAutoIncrement = false

	return c
}

func (c *Column) Test() string {
	return c.ColumnName + " " + c.DatabaseType
}

func (c *Column) GetCamelCaseColumnName() string {
	return strcase.ToLowerCamel(c.ColumnName)
}

func (c *Column) GetTitleCaseColumnName() string {
	return MakeTitle(strcase.ToCamel(c.ColumnName))
}

func (c *Column) GetPascalCaseColumnName() string {
	return strcase.ToCamel(c.ColumnName)
}

func (c *Column) GetPascalCaseTableName() string {
	return strcase.ToCamel(c.TableName)
}

func (c *Column) GetTitleCaseTableName() string {
	return MakeTitle(strcase.ToCamel(c.TableName))
}

func (c *Column) GetJavascriptDefaultValue() string {
	if c.DatabaseType == "mysql" {
		return GetJavascriptDefaultValueForMySql(c.DataType, c.NumericPrecision)
	} else if c.DatabaseType == "postgres" {
		return GetJavascriptDefaultValueForPostgres(c.DataType, c.NumericPrecision)
	}

	return "InvalidDatabaseType"
}

func (c *Column) GetJavascriptDataType() string {

	if c.DatabaseType == "mysql" {
		return GetJavascriptDataTypeForMySql(c.DataType, c.NumericPrecision)
	} else if c.DatabaseType == "postgres" {
		return GetJavascriptDataTypeForPostgres(c.DataType, c.NumericPrecision)
	}

	return "InvalidDatabaseType"
}

func (c *Column) GetJavaDataType() string {

	if c.DatabaseType == "mysql" {
		return GetJavaDataTypeForMySql(c.DataType, c.NumericPrecision)
	} else if c.DatabaseType == "postgres" {
		return GetJavaDataTypeForPostgres(c.DataType, c.NumericPrecision)
	}

	return "InvalidDatabaseType"
}

func (c *Column) IsBinary() bool {
	return c.GetJavaDataType() == "byte[]"
}

func (c *Column) IsGoIntFamilyType() bool {
	datatype := c.GetGoDataType()
	switch datatype {
	case "uint8":
		fallthrough
	case "int8":
		fallthrough
	case "uint16":
		fallthrough
	case "int16":
		fallthrough
	case "uint32":
		fallthrough
	case "int32":
		fallthrough
	case "uint64":
		fallthrough
	case "int":
		fallthrough
	case "int64":
		return true
	default:
		return false
	}
}

func (c *Column) GetAspNetRouteConstraintType() string {
	cSharpType := c.GetCSharpDataType()

	if cSharpType == "short" {
		return "int"
	}
	if cSharpType == "Byte[]" {
		return "int"
	}
	return cSharpType
}

func (c *Column) GetCSharpDataType() string {
	if c.DatabaseType == "mysql" {
		return GetCSharpDataTypeForMySql(c.DataType, c.NumericPrecision)
	} else if c.DatabaseType == "postgres" {
		return GetCSharpDataTypeForPostgres(c.DataType, c.NumericPrecision)
	}

	return "InvalidDatabaseType"
}

func (c *Column) GetGoDataType() string {
	if c.DatabaseType == "mysql" {
		return GetGoDataTypeForMySql(c.DataType, c.NumericPrecision)
	} else if c.DatabaseType == "postgres" {
		return GetGoDataTypeForPostgres(c.DataType)
	}

	return "InvalidDatabaseType"
}
func (c *Column) GetCSharpFirstUnitTestValueFromFile(valuesFile string) string {

	values := ReadUnitTestValues(valuesFile)

	a := 0
	for range values.Tables {
		if values.Tables[a].TableName == c.GetPascalCaseTableName() {
			b := 0
			for range values.Tables[a].Columns {
				if values.Tables[a].Columns[b].ColumnName == c.GetPascalCaseColumnName() {
					return values.Tables[a].Columns[b].UnitTestValues[0].FirstUnitTestValue
				}
				b++
			}
		}
		a++
	}

	if c.DatabaseType == "mysql" {
		return GetCSharpFirstUnitTestValueForMySql(c.DataType)
	} else if c.DatabaseType == "postgres" {
		return GetCSharpFirstUnitTestValueForPostgres(c.DataType)
	}
	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetJavaFirstUnitTestValueFromFile(valuesFile string) string {

	values := ReadUnitTestValues(valuesFile)

	a := 0
	for range values.Tables {
		if values.Tables[a].TableName == c.GetPascalCaseTableName() {
			b := 0
			for range values.Tables[a].Columns {
				if values.Tables[a].Columns[b].ColumnName == c.GetPascalCaseColumnName() {
					return values.Tables[a].Columns[b].UnitTestValues[0].FirstUnitTestValue
				}
				b++
			}
		}
		a++
	}

	if c.DatabaseType == "mysql" {
		return GetJavaFirstUnitTestValueForMySql(c.DataType)
	} else if c.DatabaseType == "postgres" {
		return GetJavaFirstUnitTestValueForPostgres(c.DataType)
	}
	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetCSharpFirstUnitTestValue() string {

	if c.DatabaseType == "mysql" {
		return GetCSharpFirstUnitTestValueForMySql(c.DataType)
	} else if c.DatabaseType == "postgres" {
		return GetCSharpFirstUnitTestValueForPostgres(c.DataType)
	}
	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetCSharpSecondUnitTestValue() string {

	if c.DatabaseType == "mysql" {
		return GetCSharpSecondUnitTestValueForMySql(c.DataType)
	} else if c.DatabaseType == "postgres" {
		return GetCSharpSecondUnitTestValueForPostgres(c.DataType)
	}
	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetJavaFirstUnitTestValue() string {

	if c.DatabaseType == "mysql" {
		return GetJavaFirstUnitTestValueForMySql(c.DataType)
	} else if c.DatabaseType == "postgres" {
		return GetJavaFirstUnitTestValueForPostgres(c.DataType)
	}
	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetJavaSecondUnitTestValue() string {

	if c.DatabaseType == "mysql" {
		return GetJavaSecondUnitTestValueForMySql(c.DataType)
	} else if c.DatabaseType == "postgres" {
		return GetJavaSecondUnitTestValueForPostgres(c.DataType)
	}

	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetGoFirstUnitTestValue() string {

	if c.DatabaseType == "mysql" {
		return GetGoFirstUnitTestValueForMySql(c.DataType, c.NumericPrecision)
	} else if c.DatabaseType == "postgres" {
		return GetGoFirstUnitTestValueForPostgres(c.DataType)
	}
	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetGoSecondUnitTestValue() string {

	if c.DatabaseType == "mysql" {
		return GetGoSecondUnitTestValueForMySql(c.DataType, c.NumericPrecision)
	} else if c.DatabaseType == "postgres" {
		return GetGoSecondUnitTestValueForPostgres(c.DataType)
	}

	return "Unknown" + "_" + c.DataType
}
func (c *Column) GetCSharpSecondUnitTestValueFromFile(valuesFile string) string {

	values := ReadUnitTestValues(valuesFile)

	a := 0
	for range values.Tables {
		if values.Tables[a].TableName == c.GetPascalCaseTableName() {
			b := 0
			for range values.Tables[a].Columns {
				if values.Tables[a].Columns[b].ColumnName == c.GetPascalCaseColumnName() {
					return values.Tables[a].Columns[b].UnitTestValues[1].SecondUnitTestValue
				}
				b++
			}
		}
		a++
	}
	if c.DatabaseType == "mysql" {
		return GetCSharpSecondUnitTestValueForMySql(c.DataType)
	} else if c.DatabaseType == "postgres" {
		return GetCSharpSecondUnitTestValueForPostgres(c.DataType)
	}
	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetJavaSecondUnitTestValueFromFile(valuesFile string) string {

	values := ReadUnitTestValues(valuesFile)

	a := 0
	for range values.Tables {
		if values.Tables[a].TableName == c.GetPascalCaseTableName() {
			b := 0
			for range values.Tables[a].Columns {
				if values.Tables[a].Columns[b].ColumnName == c.GetPascalCaseColumnName() {
					return values.Tables[a].Columns[b].UnitTestValues[1].SecondUnitTestValue
				}
				b++
			}
		}
		a++
	}
	if c.DatabaseType == "mysql" {
		return GetJavaSecondUnitTestValueForMySql(c.DataType)
	} else if c.DatabaseType == "postgres" {
		return GetJavaSecondUnitTestValueForPostgres(c.DataType)
	}
	return "Unknown" + "_" + c.DataType
}

func (c *Column) GetSetString() string {
	switch c.GetJavaDataType() {

	case "String":
		fallthrough

	case "Object":
		return "\""

	default:
		return ""
	}
}
func (c *Column) GetSetStringJson() string {
	switch c.GetJavaDataType() {
	case "String":
		fallthrough
	case "Object":
		fallthrough
	case "Date":
		fallthrough
	case "Timestamp":
		fallthrough
	case "Time":
		fallthrough
	case "BigDecimal":
		fallthrough
	case "Float":
		fallthrough
	case "Double":
		fallthrough
	case "byte[]":
		fallthrough
	case "Long":
		return "\""

	default:
		return ""
	}
}

func MakeTitle(s string) string {
	var parts []string
	start := 0
	for end, r := range s {
		if end != 0 && unicode.IsUpper(r) {
			parts = append(parts, s[start:end])
			start = end
		}
	}
	if start != len(s) {
		parts = append(parts, s[start:])
	}
	var title string
	for i, part := range parts {
		if i > 0 {
			title += " "
		}
		title += part
	}
	return title
}
