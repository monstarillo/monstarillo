package models

import (
	"testing"
)

func TestGetJavascriptDefaultValueForMySql_varchar(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("varchar", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_date(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("date", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_enum(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("enum", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_timestamp(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("timestamp", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_bit(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("bit", 1)
	if value != "Boolean" {
		t.Errorf("Value was incorrect, got %s, wanted boolean", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_bit_precision0(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("bit", 0)
	if value != "0" {
		t.Errorf("Value was incorrect, got %s, wanted 0", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_binary(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("binary", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_int(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("int", 0)

	if value != "0" {
		t.Errorf("Value was incorrect, got %s, wanted 0", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_bool(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("bool", 0)
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavascriptDefaultValueForMySql_BadType(t *testing.T) {
	value := GetJavascriptDefaultValueForMySql("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

func TestGetJavascriptDataTypeForMySql_varchar(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("varchar", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDataTypeForMySql_date(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("date", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavascriptDataTypeForMySql_enum(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("enum", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavascriptDataTypeForMySql_timestamp(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("timestamp", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavascriptDataTypeForMySql_bit(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("bit", 1)
	if value != "Boolean" {
		t.Errorf("Value was incorrect, got %s, wanted boolean", value)
	}
}

func TestGetJavascriptDataTypeForMySql_bit_precision0(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("bit", 0)
	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, wanted Number", value)
	}
}

func TestGetJavascriptDataTypeForMySql_binary(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("binary", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavascriptDataTypeForMySql_int(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("int", 0)
	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, wanted Number", value)
	}
}

func TestGetJavascriptDataTypeForMySql_bool(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("bool", 0)
	if value != "Boolean" {
		t.Errorf("Value was incorrect, got %s, wanted Boolean", value)
	}
}

func TestGetJavascriptDataTypeForMySql_BadType(t *testing.T) {
	value := GetJavascriptDataTypeForMySql("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

// GetJavaDataTypeForMySql

func TestGetJavaDataTypeForMySql_varchar(t *testing.T) {
	value := GetJavaDataTypeForMySql("varchar", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavaDataTypeForMySql_bit(t *testing.T) {
	value := GetJavaDataTypeForMySql("bit", 0)
	if value != "Integer" {
		t.Errorf("Value was incorrect, got %s, wanted Integer", value)
	}
}

func TestGetJavaDataTypeForMySql_bit1(t *testing.T) {
	value := GetJavaDataTypeForMySql("bit", 1)
	if value != "Boolean" {
		t.Errorf("Value was incorrect, got %s, wanted Boolean", value)
	}
}

func TestGetJavaDataTypeForMySql_int(t *testing.T) {
	value := GetJavaDataTypeForMySql("int", 0)
	if value != "Integer" {
		t.Errorf("Value was incorrect, got %s, wanted Integer", value)
	}
}
func TestGetJavaDataTypeForMySql_time(t *testing.T) {
	value := GetJavaDataTypeForMySql("time", 0)
	if value != "Time" {
		t.Errorf("Value was incorrect, got %s, wanted Time", value)
	}
}

func TestGetJavaDataTypeForMySql_date(t *testing.T) {
	value := GetJavaDataTypeForMySql("date", 0)
	if value != "Date" {
		t.Errorf("Value was incorrect, got %s, wanted Date", value)
	}
}

func TestGetJavaDataTypeForMySql_set(t *testing.T) {
	value := GetJavaDataTypeForMySql("set", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavaDataTypeForMySql_timestamp(t *testing.T) {
	value := GetJavaDataTypeForMySql("timestamp", 0)
	if value != "Timestamp" {
		t.Errorf("Value was incorrect, got %s, wanted Timestamp", value)
	}
}

func TestGetJavaDataTypeForMySql_bool(t *testing.T) {
	value := GetJavaDataTypeForMySql("bool", 0)
	if value != "Boolean" {
		t.Errorf("Value was incorrect, got %s, wanted Boolean", value)
	}
}

func TestGetJavaDataTypeForMySql_BadType(t *testing.T) {
	value := GetJavaDataTypeForMySql("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

/// CSharpDataType

func TestGetCSharpDataTypeForMySql_varchar(t *testing.T) {
	value := GetCSharpDataTypeForMySql("varchar", 0)
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetCSharpDataTypeForMySql_tinyint(t *testing.T) {
	value := GetCSharpDataTypeForMySql("tinyint", 0)
	if value != "byte" {
		t.Errorf("Value was incorrect, got %s, wanted byte", value)
	}
}

func TestGetCSharpDataTypeForMySql_bigint(t *testing.T) {
	value := GetCSharpDataTypeForMySql("bigint", 0)
	if value != "Int64" {
		t.Errorf("Value was incorrect, got %s, wanted byte", value)
	}
}

func TestGetCSharpDataTypeForMySql_float(t *testing.T) {
	value := GetCSharpDataTypeForMySql("float", 0)
	if value != "Decimal" {
		t.Errorf("Value was incorrect, got %s, wanted Decimal", value)
	}
}

func TestGetCSharpDataTypeForMySql_double(t *testing.T) {
	value := GetCSharpDataTypeForMySql("double", 0)
	if value != "Long" {
		t.Errorf("Value was incorrect, got %s, wanted Long", value)
	}
}
func TestGetCSharpDataTypeForMySql_bit(t *testing.T) {
	value := GetCSharpDataTypeForMySql("bit", 0)
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetCSharpDataTypeForMySql_bit1(t *testing.T) {
	value := GetCSharpDataTypeForMySql("bit", 1)
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetCSharpDataTypeForMySql_int(t *testing.T) {
	value := GetCSharpDataTypeForMySql("int", 0)
	if value != "int" {
		t.Errorf("Value was incorrect, got %s, wanted int", value)
	}
}
func TestGetCSharpDataTypeForMySql_time(t *testing.T) {
	value := GetCSharpDataTypeForMySql("time", 0)
	if value != "DateTime" {
		t.Errorf("Value was incorrect, got %s, wanted DateTime", value)
	}
}

func TestGetCSharpDataTypeForMySql_date(t *testing.T) {
	value := GetCSharpDataTypeForMySql("date", 0)
	if value != "DateTime" {
		t.Errorf("Value was incorrect, got %s, wanted DateTime", value)
	}
}

func TestGetCSharpDataTypeForMySql_set(t *testing.T) {
	value := GetCSharpDataTypeForMySql("set", 0)
	if value != "Object" {
		t.Errorf("Value was incorrect, got %s, wanted Object", value)
	}
}

func TestGetCSharpDataTypeForMySql_timestamp(t *testing.T) {
	value := GetCSharpDataTypeForMySql("timestamp", 0)
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetCSharpDataTypeForMySql_bool(t *testing.T) {
	value := GetCSharpDataTypeForMySql("bool", 0)
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetCSharpDataTypeForMySql_BadType(t *testing.T) {
	value := GetCSharpDataTypeForMySql("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

/// GoDataType

func TestGetGoDataTypeForMySql_varchar(t *testing.T) {
	value := GetGoDataTypeForMySql("varchar", 0)
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetGoDataTypeForMySql_tinyint(t *testing.T) {
	value := GetGoDataTypeForMySql("tinyint", 0)
	if value != "int8" {
		t.Errorf("Value was incorrect, got %s, wanted int8", value)
	}
}

func TestGetGoDataTypeForMySql_bigint(t *testing.T) {
	value := GetGoDataTypeForMySql("bigint", 0)
	if value != "int64" {
		t.Errorf("Value was incorrect, got %s, wanted int64", value)
	}
}

func TestGetGoDataTypeForMySql_float(t *testing.T) {
	value := GetGoDataTypeForMySql("float", 0)
	if value != "float32" {
		t.Errorf("Value was incorrect, got %s, wanted float32", value)
	}
}

func TestGetGoDataTypeForMySql_double(t *testing.T) {
	value := GetGoDataTypeForMySql("double", 0)
	if value != "float64" {
		t.Errorf("Value was incorrect, got %s, wanted float64", value)
	}
}
func TestGetGoDataTypeForMySql_bit(t *testing.T) {
	value := GetGoDataTypeForMySql("bit", 0)
	if value != "int8" {
		t.Errorf("Value was incorrect, got %s, wanted int8", value)
	}
}

func TestGetGoDataTypeForMySql_bit1(t *testing.T) {
	value := GetGoDataTypeForMySql("bit", 1)
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetGoDataTypeForMySql_int(t *testing.T) {
	value := GetGoDataTypeForMySql("int", 0)
	if value != "int" {
		t.Errorf("Value was incorrect, got %s, wanted int", value)
	}
}
func TestGetGoDataTypeForMySql_time(t *testing.T) {
	value := GetGoDataTypeForMySql("time", 0)
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted time.Time", value)
	}
}

func TestGetGoDataTypeForMySql_date(t *testing.T) {
	value := GetGoDataTypeForMySql("date", 0)
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted datatypes.Date", value)
	}
}

func TestGetGoDataTypeForMySql_set(t *testing.T) {
	value := GetGoDataTypeForMySql("set", 0)
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetGoDataTypeForMySql_timestamp(t *testing.T) {
	value := GetGoDataTypeForMySql("timestamp", 0)
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted time.Time", value)
	}
}

func TestGetGoDataTypeForMySql_bool(t *testing.T) {
	value := GetGoDataTypeForMySql("bool", 0)
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetGoDataTypeForMySql_BadType(t *testing.T) {
	value := GetGoDataTypeForMySql("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

// CSharpSecond

func TestGetCSharpSecondUnitTestValueForMySql_varchar(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("varchar")
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_bit(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("bit")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_bit1(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("bit")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_int(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("int")
	if value != "3" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}
func TestGetCSharpSecondUnitTestValueForMySql_time(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("time")
	if value != "new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)" {
		t.Errorf("Value was incorrect, got %s, wanted new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_date(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("date")
	if value != "new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)" {
		t.Errorf("Value was incorrect, got %s, wanted new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_set(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("set")
	if value != "BB" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_timestamp(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("timestamp")
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_bool(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("bool")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_BadType(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_bigint(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("bigint")
	if value != "2000" {
		t.Errorf("Value was incorrect, got %s, wanted byte", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_float(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("float")
	if value != "3500F" {
		t.Errorf("Value was incorrect, got %s, wanted 3500F", value)
	}
}

func TestGetCSharpSecondUnitTestValueForMySql_double(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForMySql("double")
	if value != "4500D" {
		t.Errorf("Value was incorrect, got %s, wanted 4500D", value)
	}
}

//  CSharpFirst

func TestGetCSharpFirstUnitTestValueForMySql_varchar(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("varchar")
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_bit(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("bit")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_bit1(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("bit")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_int(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("int")
	if value != "2" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_date(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("date")
	if value != "new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day)" {
		t.Errorf("Value was incorrect, got %s, wanted new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day)", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_set(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("set")
	if value != "AA" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_timestamp(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("timestamp")
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_bool(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("bool")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_bigint(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("bigint")
	if value != "1000" {
		t.Errorf("Value was incorrect, got %s, wanted byte", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_float(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("float")
	if value != "3000F" {
		t.Errorf("Value was incorrect, got %s, wanted 3500F", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_double(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("double")
	if value != "4000D" {
		t.Errorf("Value was incorrect, got %s, wanted 4500D", value)
	}
}

func TestGetCSharpFirstUnitTestValueForMySql_BadType(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForMySql("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

//  GoFirst

func TestGetGoFirstUnitTestValueForMySql_varchar(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("varchar", 0)
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_bit(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("bit", 0)
	if value != "1" {
		t.Errorf("Value was incorrect, got %s, wanted 1", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_bit1(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("bit", 1)
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_int(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("int", 0)
	if value != "2" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_date(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("date", 0)
	if value != "datatypes.Date" {
		t.Errorf("Value was incorrect, got %s, wanted datatypes.Date)", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_set(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("set", 0)
	if value != "AA" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_timestamp(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("timestamp", 0)
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted time.Time", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_bool(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("bool", 0)
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_bigint(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("bigint", 0)
	if value != "22" {
		t.Errorf("Value was incorrect, got %s, wanted 22", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_float(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("float", 0)
	if value != "33" {
		t.Errorf("Value was incorrect, got %s, wanted 33", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_double(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("double", 0)
	if value != "33" {
		t.Errorf("Value was incorrect, got %s, wanted 33", value)
	}
}

func TestGetGoFirstUnitTestValueForMySql_BadType(t *testing.T) {
	value := GetGoFirstUnitTestValueForMySql("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

//  GoSecond

func TestGetGoSecondUnitTestValueForMySql_varchar(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("varchar", 0)
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_bit(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("bit", 0)
	if value != "3" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_bit1(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("bit", 1)
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_int(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("int", 0)
	if value != "3" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_date(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("date", 0)
	if value != "datatypes.Date" {
		t.Errorf("Value was incorrect, got %s, wanted datatypes.Date)", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_set(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("set", 0)
	if value != "BB" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_timestamp(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("timestamp", 0)
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted time.Time", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_bool(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("bool", 0)
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_bigint(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("bigint", 0)
	if value != "33" {
		t.Errorf("Value was incorrect, got %s, wanted 33", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_float(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("float", 0)
	if value != "44" {
		t.Errorf("Value was incorrect, got %s, wanted 44", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_double(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("double", 0)
	if value != "44" {
		t.Errorf("Value was incorrect, got %s, wanted 44", value)
	}
}

func TestGetGoSecondUnitTestValueForMySql_BadType(t *testing.T) {
	value := GetGoSecondUnitTestValueForMySql("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

//  JavaFirst

func TestGetJavaFirstUnitTestValueForMySql_varchar(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("varchar")
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_bit(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("bit")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_bit1(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("bit")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_int(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("int")
	if value != "2" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_set(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("set")
	if value != "AA" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_bool(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("bool")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_bigint(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("bigint")
	if value != "1000L" {
		t.Errorf("Value was incorrect, got %s, wanted 1000L", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_float(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("float")
	if value != "new Float(3)" {
		t.Errorf("Value was incorrect, got %s, wanted new Float(3)", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_double(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("double")
	if value != "new Double(4)" {
		t.Errorf("Value was incorrect, got %s, wanted new Double(4)", value)
	}
}

func TestGetJavaFirstUnitTestValueForMySql_BadType(t *testing.T) {
	value := GetJavaFirstUnitTestValueForMySql("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

//  JavaSecond

func TestGetJavaSecondUnitTestValueForMySql_varchar(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("varchar")
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_bit(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("bit")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_bit1(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("bit")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_int(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("int")
	if value != "3" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_set(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("set")
	if value != "BB" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_bool(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("bool")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_bigint(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("bigint")
	if value != "2000L" {
		t.Errorf("Value was incorrect, got %s, wanted 2000L", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_float(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("float")
	if value != "new Float(33)" {
		t.Errorf("Value was incorrect, got %s, wanted new Float(33)", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_double(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("double")
	if value != "new Double(44)" {
		t.Errorf("Value was incorrect, got %s, wanted new Double(44)", value)
	}
}

func TestGetJavaSecondUnitTestValueForMySql_BadType(t *testing.T) {
	value := GetJavaSecondUnitTestValueForMySql("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}
