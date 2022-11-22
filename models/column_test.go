package models

import (
	"errors"
	"testing"
)

func getColumn(tableName, columnName string) (Column, error) {
	table, _ := getTable(tableName)
	for _, c := range table.Columns {
		if c.ColumnName == columnName {
			return c, nil
		}
	}

	var column Column
	return column, errors.New("Column not found")
}
func TestColumn_GetPascalCaseTableName(t *testing.T) {

	column, _ := getColumn("film", "film_id")
	value := column.GetPascalCaseTableName()

	if value != "Film" {
		t.Errorf("Value was incorrect, got %s, want Film", value)
	}
}

func TestColumn_GetPascalCaseColumnName(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetPascalCaseColumnName()

	if value != "FilmId" {
		t.Errorf("Value was incorrect, got %s, want FilmId", value)
	}
}

func TestColumn_GetCamelCaseColumnName(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetCamelCaseColumnName()

	if value != "filmId" {
		t.Errorf("Value was incorrect, got %s, want filmId", value)
	}
}

func TestColumn_Test(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.Test()

	if value != "film_id mysql" {
		t.Errorf("Value was incorrect, got %s, want film_id mysql", value)
	}
}

func TestColumn_GetJavascriptDefaultValue_mysql(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetJavascriptDefaultValue()

	if value != "0" {
		t.Errorf("Value was incorrect, got %s, want 0", value)
	}
}

func TestColumn_GetJavascriptDefaultValue_postgres(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "postgres"
	value := column.GetJavascriptDefaultValue()

	if value != "0" {
		t.Errorf("Value was incorrect, got %s, want 0", value)
	}
}

func TestColumn_GetJavascriptDefaultValue_badDb(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "badDb"
	value := column.GetJavascriptDefaultValue()

	if value != "InvalidDatabaseType" {
		t.Errorf("Value was incorrect, got %s, want InvalidDatabaseType", value)
	}
}

func TestColumn_GetJavascriptDataType_mysql(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetJavascriptDataType()

	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, want Number", value)
	}
}

func TestColumn_GetJavascriptDataType_postgres(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "postgres"
	value := column.GetJavascriptDataType()

	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, want Number", value)
	}
}

func TestColumn_GetJavascriptDataType_badDb(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "badDb"
	value := column.GetJavaDataType()

	if value != "InvalidDatabaseType" {
		t.Errorf("Value was incorrect, got %s, want InvalidDatabaseType", value)
	}
}

func TestColumn_GetJavaDataType_mysql(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetJavaDataType()

	if value != "Integer" {
		t.Errorf("Value was incorrect, got %s, want Ingteger", value)
	}
}

func TestColumn_GetJavaDataType_postgres(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "postgres"
	value := column.GetJavaDataType()

	if value != "Integer" {
		t.Errorf("Value was incorrect, got %s, want Integer", value)
	}
}

func TestColumn_GetJavaDataType_badDb(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "badDb"
	value := column.GetJavaDataType()

	if value != "InvalidDatabaseType" {
		t.Errorf("Value was incorrect, got %s, want InvalidDatabaseType", value)
	}
}

func TestColumn_GetGoDataType_mysql(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetGoDataType()

	if value != "int8" {
		t.Errorf("Value was incorrect, got %s, want int8", value)
	}
}

func TestColumn_GetGoDataType_postgres(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "postgres"
	value := column.GetGoDataType()

	if value != "int" {
		t.Errorf("Value was incorrect, got %s, want int", value)
	}
}

func TestColumn_GetGoDataType_badDb(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "badDb"
	value := column.GetGoDataType()

	if value != "InvalidDatabaseType" {
		t.Errorf("Value was incorrect, got %s, want InvalidDatabaseType", value)
	}
}
func TestColumn_IsBinary_false(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.IsBinary()

	if value != false {
		t.Errorf("Value was incorrect, got %t, want false", value)
	}
}

func TestColumn_IsBinary_trueMySql(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "binary"
	value := column.IsBinary()

	if value != true {
		t.Errorf("Value was incorrect, got %t, want true", value)
	}
}

func TestColumn_IsBinary_truePostgres(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "postgres"
	column.DataType = "binary"
	value := column.IsBinary()

	if value != true {
		t.Errorf("Value was incorrect, got %t, want true", value)
	}
}

func TestColumn_GetCSharpDataType_mysql(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetCSharpDataType()

	if value != "short" {
		t.Errorf("Value was incorrect, got %s, want short", value)
	}
}

func TestColumn_GetCSharpDataType_postgres(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "postgres"
	value := column.GetCSharpDataType()

	if value != "short" {
		t.Errorf("Value was incorrect, got %s, want short", value)
	}
}

func TestColumn_GetCSharpDataType_badDb(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DatabaseType = "badDb"
	value := column.GetCSharpDataType()

	if value != "InvalidDatabaseType" {
		t.Errorf("Value was incorrect, got %s, want InvalidDatabaseType", value)
	}
}

func TestColumn_GetGetSetStringJson_Default(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetSetStringJson()

	if value != "" {
		t.Errorf("Value was incorrect, got %s, want ''", value)
	}
}

func TestColumn_GetGetSetStringJson_String(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "varchar"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_Object(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "enum"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_Date(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "date"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_Timestamp(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "timestamp"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_Time(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "time"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_BigDecimal(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "dec"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_Float(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "float"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_Double(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "double"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_Byte(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "binary"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetGetSetStringJson_Long(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	column.DataType = "bigint"
	value := column.GetSetStringJson()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_NewColumn(t *testing.T) {
	column := NewColumn("id", "mysql", "person")

	if column.ColumnName != "id" {
		t.Errorf("Value was incorrect, got %s, want id", column.ColumnName)
	}

}

func TestColumn_GetCSharpFirstUnitTestValueFromFile(t *testing.T) {
	column, _ := getColumn("actor", "first_name")
	value := column.GetCSharpFirstUnitTestValueFromFile("../test-data/unit-test.json")
	if value != "Apple" {
		t.Errorf("Value was incorrect, got %s, want Apple", value)
	}
}

func TestColumn_GetCSharpFirstUnitTestValue(t *testing.T) {
	column, _ := getColumn("actor", "first_name")
	value := column.GetCSharpFirstUnitTestValue()
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, want A", value)
	}
}
func TestColumn_GetCSharpSecondUnitTestValue(t *testing.T) {
	column, _ := getColumn("actor", "first_name")
	value := column.GetCSharpSecondUnitTestValue()
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, want B", value)
	}
}
func TestColumn_GetCSharpFirstUnitTestValueFromFile_NotIncluded(t *testing.T) {
	column, _ := getColumn("actor", "first_name")
	column.ColumnName = "xxx"
	value := column.GetCSharpFirstUnitTestValueFromFile("../test-data/unit-test.json")
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, want A", value)
	}
}

func TestColumn_GetCSharpSecondUnitTestValueFromFile(t *testing.T) {
	column, _ := getColumn("actor", "first_name")
	value := column.GetCSharpSecondUnitTestValueFromFile("../test-data/unit-test.json")
	if value != "Batman" {
		t.Errorf("Value was incorrect, got %s, want Batman", value)
	}
}

func TestColumn_GetCSharpSecondUnitTestValueFromFile_NotIncluded(t *testing.T) {
	column, _ := getColumn("actor", "first_name")
	column.ColumnName = "xxx"
	value := column.GetCSharpSecondUnitTestValueFromFile("../test-data/unit-test.json")
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, want B", value)
	}
}

func TestColumn_GetGetSetString_Default(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetSetString()

	if value != "" {
		t.Errorf("Value was incorrect, got %s, want ''", value)
	}
}

func TestColumn_GetGetSetString_String(t *testing.T) {
	column, _ := getColumn("film", "title")
	value := column.GetSetString()

	if value != "\"" {
		t.Errorf("Value was incorrect, got %s, want '\"'", value)
	}
}

func TestColumn_GetAspNetRouteConstraintType_default(t *testing.T) {
	column, _ := getColumn("film", "title")
	value := column.GetAspNetRouteConstraintType()

	if value != "string" {
		t.Errorf("Value was incorrect, got %s, want string", value)
	}
}

func TestColumn_GetAspNetRouteConstraintType_int(t *testing.T) {
	column, _ := getColumn("film", "film_id")
	value := column.GetAspNetRouteConstraintType()

	if value != "int" {
		t.Errorf("Value was incorrect, got %s, want int", value)
	}
}

func TestColumn_GetAspNetRouteConstraintType_binary(t *testing.T) {
	column := NewColumn("file_data", "mysql", "person")
	column.DataType = "binary"
	value := column.GetAspNetRouteConstraintType()

	if value != "int" {
		t.Errorf("Value was incorrect, got %s, want int", value)
	}
}

func TestColumn_GetJavaFirstUnitTestValue(t *testing.T) {
	column := NewColumn("file_data", "mysql", "person")
	column.DataType = "varchar"

	value := column.GetJavaFirstUnitTestValue()
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, want A", value)
	}
}

func TestColumn_GetJavaSecondUnitTestValue(t *testing.T) {
	column := NewColumn("file_data", "mysql", "person")
	column.DataType = "varchar"

	value := column.GetJavaSecondUnitTestValue()
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, want B", value)
	}
}

func TestColumn_GetGoFirstUnitTestValue(t *testing.T) {
	column := NewColumn("file_data", "mysql", "person")
	column.DataType = "varchar"

	value := column.GetGoFirstUnitTestValue()
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, want A", value)
	}
}

func TestColumn_GetGoSecondUnitTestValue(t *testing.T) {
	column := NewColumn("file_data", "mysql", "person")
	column.DataType = "varchar"

	value := column.GetGoSecondUnitTestValue()
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, want B", value)
	}
}

func TestColumn_GetJavaSecondUnitTestValue_postgres(t *testing.T) {
	column := NewColumn("file_data", "mysql", "person")
	column.DataType = "varchar"
	column.DatabaseType = "postgres"

	value := column.GetJavaSecondUnitTestValue()
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, want B", value)
	}
}

func TestColumn_GetGoFirstUnitTestValue_postgres(t *testing.T) {
	column := NewColumn("file_data", "mysql", "person")
	column.DataType = "varchar"
	column.DatabaseType = "postgres"

	value := column.GetGoFirstUnitTestValue()
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, want A", value)
	}
}

func TestColumn_GetGoSecondUnitTestValue_postgres(t *testing.T) {
	column := NewColumn("file_data", "mysql", "person")
	column.DataType = "varchar"
	column.DatabaseType = "postgres"

	value := column.GetGoSecondUnitTestValue()
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, want B", value)
	}
}
