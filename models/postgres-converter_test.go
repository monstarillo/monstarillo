package models

import (
	"testing"
)

func TestGetJavascriptDefaultValueForPostgres_varchar(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("varchar", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_date(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("date", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_bytea(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("bytea", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_timestamp(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("timestamp", 0)
	stringDefault := "\"\""
	if value != stringDefault {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_int(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("int", 0)

	if value != "0" {
		t.Errorf("Value was incorrect, got %s, wanted 0", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_bigint(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("bigint", 0)
	if value != "0" {
		t.Errorf("Value was incorrect, got %s, wanted 0", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_decimal(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("decimal", 0)
	if value != "0" {
		t.Errorf("Value was incorrect, got %s, wanted 0", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_float4(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("float4", 0)
	if value != "0" {
		t.Errorf("Value was incorrect, got %s, wanted 0", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_bool(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("bool", 0)
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavascriptDefaultValueForPostgres_BadType(t *testing.T) {
	value := GetJavascriptDefaultValueForPostgres("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_varchar(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("varchar", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted ''", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_date(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("date", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_bytea(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("bytea", 0)
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_timestamp(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("timestamp", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_int(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("int", 0)
	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, wanted Number", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_bool(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("bool", 0)
	if value != "Boolean" {
		t.Errorf("Value was incorrect, got %s, wanted Boolean", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_bigint(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("bigint", 0)
	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, wanted Number", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_decimal(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("decimal", 0)
	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, wanted Number", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_float4(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("float4", 0)
	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, wanted Number", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_double(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("double", 0)
	if value != "Number" {
		t.Errorf("Value was incorrect, got %s, wanted Number", value)
	}
}

func TestGetJavascriptDataTypeForPostgres_BadType(t *testing.T) {
	value := GetJavascriptDataTypeForPostgres("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

// GetJavaDataTypeForPostgres

func TestGetJavaDataTypeForPostgres_varchar(t *testing.T) {
	value := GetJavaDataTypeForPostgres("varchar", 0)
	if value != "String" {
		t.Errorf("Value was incorrect, got %s, wanted String", value)
	}
}

func TestGetJavaDataTypeForPostgres_int(t *testing.T) {
	value := GetJavaDataTypeForPostgres("int", 0)
	if value != "Integer" {
		t.Errorf("Value was incorrect, got %s, wanted Integer", value)
	}
}
func TestGetJavaDataTypeForPostgres_time(t *testing.T) {
	value := GetJavaDataTypeForPostgres("time", 0)
	if value != "Time" {
		t.Errorf("Value was incorrect, got %s, wanted Time", value)
	}
}

func TestGetJavaDataTypeForPostgres_date(t *testing.T) {
	value := GetJavaDataTypeForPostgres("date", 0)
	if value != "Date" {
		t.Errorf("Value was incorrect, got %s, wanted Date", value)
	}
}

func TestGetJavaDataTypeForPostgres_int8(t *testing.T) {
	value := GetJavaDataTypeForPostgres("int8", 0)
	if value != "Long" {
		t.Errorf("Value was incorrect, got %s, wanted Object", value)
	}
}

func TestGetJavaDataTypeForPostgres_decimal(t *testing.T) {
	value := GetJavaDataTypeForPostgres("decimal", 0)
	if value != "BigDecimal" {
		t.Errorf("Value was incorrect, got %s, wanted BigDecimal", value)
	}
}

func TestGetJavaDataTypeForPostgres_float4(t *testing.T) {
	value := GetJavaDataTypeForPostgres("float4", 0)
	if value != "Float" {
		t.Errorf("Value was incorrect, got %s, wanted Float", value)
	}
}

func TestGetJavaDataTypeForPostgres_double(t *testing.T) {
	value := GetJavaDataTypeForPostgres("double", 0)
	if value != "Double" {
		t.Errorf("Value was incorrect, got %s, wanted Object", value)
	}
}

func TestGetJavaDataTypeForPostgres_binary(t *testing.T) {
	value := GetJavaDataTypeForPostgres("binary", 0)
	if value != "byte[]" {
		t.Errorf("Value was incorrect, got %s, wanted byte[]", value)
	}
}

func TestGetJavaDataTypeForPostgres_timestamp(t *testing.T) {
	value := GetJavaDataTypeForPostgres("timestamp", 0)
	if value != "Timestamp" {
		t.Errorf("Value was incorrect, got %s, wanted Timestamp", value)
	}
}

func TestGetJavaDataTypeForPostgres_bool(t *testing.T) {
	value := GetJavaDataTypeForPostgres("bool", 0)
	if value != "Boolean" {
		t.Errorf("Value was incorrect, got %s, wanted Boolean", value)
	}
}

func TestGetJavaDataTypeForPostgres_BadType(t *testing.T) {
	value := GetJavaDataTypeForPostgres("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

// GetGoDataTypeForPostgres

func TestGetGoDataTypeForPostgres_varchar(t *testing.T) {
	value := GetGoDataTypeForPostgres("varchar")
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetGoDataTypeForPostgres_int(t *testing.T) {
	value := GetGoDataTypeForPostgres("int")
	if value != "int" {
		t.Errorf("Value was incorrect, got %s, wanted int", value)
	}
}
func TestGetGoDataTypeForPostgres_time(t *testing.T) {
	value := GetGoDataTypeForPostgres("time")
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted time.Time", value)
	}
}

func TestGetGoDataTypeForPostgres_date(t *testing.T) {
	value := GetGoDataTypeForPostgres("date")
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted datatypes.Date", value)
	}
}

func TestGetGoDataTypeForPostgres_int8(t *testing.T) {
	value := GetGoDataTypeForPostgres("int8")
	if value != "int64" {
		t.Errorf("Value was incorrect, got %s, wanted int64", value)
	}
}

func TestGetGoDataTypeForPostgres_decimal(t *testing.T) {
	value := GetGoDataTypeForPostgres("decimal")
	if value != "float64" {
		t.Errorf("Value was incorrect, got %s, wanted float64", value)
	}
}

func TestGetGoDataTypeForPostgres_float4(t *testing.T) {
	value := GetGoDataTypeForPostgres("float4")
	if value != "float32" {
		t.Errorf("Value was incorrect, got %s, wanted float32", value)
	}
}

func TestGetGoDataTypeForPostgres_double(t *testing.T) {
	value := GetGoDataTypeForPostgres("double")
	if value != "float64" {
		t.Errorf("Value was incorrect, got %s, wanted float64", value)
	}
}

func TestGetGoDataTypeForPostgres_binary(t *testing.T) {
	value := GetGoDataTypeForPostgres("binary")
	if value != "[]byte" {
		t.Errorf("Value was incorrect, got %s, wanted []byte", value)
	}
}

func TestGetGoDataTypeForPostgres_timestamp(t *testing.T) {
	value := GetGoDataTypeForPostgres("timestamp")
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted time.Time", value)
	}
}

func TestGetGoDataTypeForPostgres_bool(t *testing.T) {
	value := GetGoDataTypeForPostgres("bool")
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetGoDataTypeForPostgres_BadType(t *testing.T) {
	value := GetGoDataTypeForPostgres("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

/// CSharpDataType

func TestGetCSharpDataTypeForPostgres_varchar(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("varchar", 0)
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetCSharpDataTypeForPostgres_tinyint(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("tinyint", 0)
	if value != "byte" {
		t.Errorf("Value was incorrect, got %s, wanted byte", value)
	}
}

func TestGetCSharpDataTypeForPostgres_bigint(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("bigint", 0)
	if value != "Int64" {
		t.Errorf("Value was incorrect, got %s, wanted byte", value)
	}
}

func TestGetCSharpDataTypeForPostgres_float(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("float4", 0)
	if value != "Decimal" {
		t.Errorf("Value was incorrect, got %s, wanted Decimal", value)
	}
}

func TestGetCSharpDataTypeForPostgres_double(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("double", 0)
	if value != "Long" {
		t.Errorf("Value was incorrect, got %s, wanted Long", value)
	}
}
func TestGetCSharpDataTypeForPostgres_bit(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("bit", 0)
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetCSharpDataTypeForPostgres_bit1(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("bit", 1)
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetCSharpDataTypeForPostgres_int(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("int", 0)
	if value != "int" {
		t.Errorf("Value was incorrect, got %s, wanted int", value)
	}
}
func TestGetCSharpDataTypeForPostgres_time(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("time", 0)
	if value != "DateTime" {
		t.Errorf("Value was incorrect, got %s, wanted DateTime", value)
	}
}

func TestGetCSharpDataTypeForPostgres_date(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("date", 0)
	if value != "DateTime" {
		t.Errorf("Value was incorrect, got %s, wanted DateTime", value)
	}
}

//func TestGetCSharpDataTypeForPostgres_set(t *testing.T) {
//	value := GetCSharpDataTypeForPostgres("set", 0)
//	if value != "Object" {
//		t.Errorf("Value was incorrect, got %s, wanted Object", value)
//	}
//}

func TestGetCSharpDataTypeForPostgres_timestamp(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("timestamp", 0)
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetCSharpDataTypeForPostgres_bool(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("bool", 0)
	if value != "bool" {
		t.Errorf("Value was incorrect, got %s, wanted bool", value)
	}
}

func TestGetCSharpDataTypeForPostgres_BadType(t *testing.T) {
	value := GetCSharpDataTypeForPostgres("xxxx", 0)
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

// CSharpSecond

func TestGetCSharpSecondUnitTestValueForPostgres_varchar(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("varchar")
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_bit(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("bit")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_bit1(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("bit")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_int(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("int")
	if value != "3" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}
func TestGetCSharpSecondUnitTestValueForPostgres_time(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("time")
	if value != "new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)" {
		t.Errorf("Value was incorrect, got %s, wanted new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_date(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("date")
	if value != "new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)" {
		t.Errorf("Value was incorrect, got %s, wanted new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)", value)
	}
}

//func TestGetCSharpSecondUnitTestValueForPostgres_set(t *testing.T) {
//	value := GetCSharpSecondUnitTestValueForPostgres("set")
//	if value != "BB" {
//		t.Errorf("Value was incorrect, got %s, wanted BB", value)
//	}
//}

func TestGetCSharpSecondUnitTestValueForPostgres_timestamp(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("timestamp")
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_bool(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("bool")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted false", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_BadType(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_bigint(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("bigint")
	if value != "2000" {
		t.Errorf("Value was incorrect, got %s, wanted byte", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_float(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("float4")
	if value != "3500F" {
		t.Errorf("Value was incorrect, got %s, wanted 3500F", value)
	}
}

func TestGetCSharpSecondUnitTestValueForPostgres_double(t *testing.T) {
	value := GetCSharpSecondUnitTestValueForPostgres("double")
	if value != "4500D" {
		t.Errorf("Value was incorrect, got %s, wanted 4500D", value)
	}
}

//  CSharpFirst

func TestGetCSharpFirstUnitTestValueForPostgres_varchar(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("varchar")
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_bit(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("bit")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_int(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("int")
	if value != "2" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_date(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("date")
	if value != "new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day)" {
		t.Errorf("Value was incorrect, got %s, wanted new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day)", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_set(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("set")
	if value != "AA" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_timestamp(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("timestamp")
	if value != "string" {
		t.Errorf("Value was incorrect, got %s, wanted string", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_bool(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("bool")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_bigint(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("bigint")
	if value != "1000" {
		t.Errorf("Value was incorrect, got %s, wanted byte", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_float(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("float4")
	if value != "3000F" {
		t.Errorf("Value was incorrect, got %s, wanted 3500F", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_double(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("double")
	if value != "4000D" {
		t.Errorf("Value was incorrect, got %s, wanted 4500D", value)
	}
}

func TestGetCSharpFirstUnitTestValueForPostgres_BadType(t *testing.T) {
	value := GetCSharpFirstUnitTestValueForPostgres("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

//  GoFirst

func TestGetGoFirstUnitTestValueForPostgres_varchar(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("varchar")
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_bit(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("bit")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_int(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("int")
	if value != "2" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_date(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("date")
	if value != "datatypes.Date" {
		t.Errorf("Value was incorrect, got %s, wanted datatypes.Date)", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_set(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("set")
	if value != "AA" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_timestamp(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("timestamp")
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted time.Time", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_bool(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("bool")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_bigint(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("bigint")
	if value != "33" {
		t.Errorf("Value was incorrect, got %s, wanted 33", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_float(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("float4")
	if value != "55" {
		t.Errorf("Value was incorrect, got %s, wanted 55", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_double(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("double")
	if value != "44" {
		t.Errorf("Value was incorrect, got %s, wanted 44", value)
	}
}

func TestGetGoFirstUnitTestValueForPostgres_BadType(t *testing.T) {
	value := GetGoFirstUnitTestValueForPostgres("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

//  GoSecond

func TestGetGoSecondUnitTestValueForPostgres_varchar(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("varchar")
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_bit(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("bit")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_int(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("int")
	if value != "3" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_date(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("date")
	if value != "datatypes.Date" {
		t.Errorf("Value was incorrect, got %s, wanted datatypes.Date)", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_set(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("set")
	if value != "BB" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_timestamp(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("timestamp")
	if value != "time.Time" {
		t.Errorf("Value was incorrect, got %s, wanted time.Time", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_bool(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("bool")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_bigint(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("bigint")
	if value != "55" {
		t.Errorf("Value was incorrect, got %s, wanted 55", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_float(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("float4")
	if value != "77" {
		t.Errorf("Value was incorrect, got %s, wanted 77", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_double(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("double")
	if value != "66" {
		t.Errorf("Value was incorrect, got %s, wanted 66", value)
	}
}

func TestGetGoSecondUnitTestValueForPostgres_BadType(t *testing.T) {
	value := GetGoSecondUnitTestValueForPostgres("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

//  JavaFirst

func TestGetJavaFirstUnitTestValueForPostgres_varchar(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("varchar")
	if value != "A" {
		t.Errorf("Value was incorrect, got %s, wanted A", value)
	}
}

func TestGetJavaFirstUnitTestValueForPostgres_bit(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("bit")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavaFirstUnitTestValueForPostgres_int(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("int")
	if value != "2" {
		t.Errorf("Value was incorrect, got %s, wanted 2", value)
	}
}

func TestGetJavaFirstUnitTestValueForPostgres_set(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("set")
	if value != "AA" {
		t.Errorf("Value was incorrect, got %s, wanted AA", value)
	}
}

func TestGetJavaFirstUnitTestValueForPostgres_bool(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("bool")
	if value != "true" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavaFirstUnitTestValueForPostgres_bigint(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("bigint")
	if value != "1000L" {
		t.Errorf("Value was incorrect,got%s, wanted 1000L", value)
	}
}

func TestGetJavaFirstUnitTestValueForPostgres_float(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("float4")
	if value != "new Float(3)" {
		t.Errorf("Value was incorrect, got%s, wanted new Float(3)", value)
	}
}

func TestGetJavaFirstUnitTestValueForPostgres_double(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("double")
	if value != "new Double(4)" {
		t.Errorf("Value was incorrect, got%s, wanted new Double(4)", value)
	}
}

func TestGetJavaFirstUnitTestValueForPostgres_BadType(t *testing.T) {
	value := GetJavaFirstUnitTestValueForPostgres("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}

//  JavaSecond

func TestGetJavaSecondUnitTestValueForPostgres_varchar(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("varchar")
	if value != "B" {
		t.Errorf("Value was incorrect, got %s, wanted B", value)
	}
}

func TestGetJavaSecondUnitTestValueForPostgres_bit(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("bit")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavaSecondUnitTestValueForPostgres_int(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("int")
	if value != "3" {
		t.Errorf("Value was incorrect, got %s, wanted 3", value)
	}
}

func TestGetJavaSecondUnitTestValueForPostgres_set(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("set")
	if value != "BB" {
		t.Errorf("Value was incorrect, got %s, wanted BB", value)
	}
}

func TestGetJavaSecondUnitTestValueForPostgres_bool(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("bool")
	if value != "false" {
		t.Errorf("Value was incorrect, got %s, wanted true", value)
	}
}

func TestGetJavaSecondUnitTestValueForPostgres_bigint(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("bigint")
	if value != "2000L" {
		t.Errorf("Value was incorrect,got%s, wanted 2000L", value)
	}
}

func TestGetJavaSecondUnitTestValueForPostgres_float(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("float4")
	if value != "new Float(33)" {
		t.Errorf("Value was incorrect, got%s, wanted new Float(33)", value)
	}
}

func TestGetJavaSecondUnitTestValueForPostgres_double(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("double")
	if value != "new Double(44)" {
		t.Errorf("Value was incorrect, got%s, wanted new Double(44)", value)
	}
}

func TestGetJavaSecondUnitTestValueForPostgres_BadType(t *testing.T) {
	value := GetJavaSecondUnitTestValueForPostgres("xxxx")
	if value != "Unknown_xxxx" {
		t.Errorf("Value was incorrect, got %s, wanted Unknown_xxxx", value)
	}
}
