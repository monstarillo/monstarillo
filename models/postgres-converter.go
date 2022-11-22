package models

import (
	"strconv"
	"time"
)

func GetJavascriptDefaultValueForPostgres(dataType string, numericPrecision int) string {
	stringDefault := "\"\""
	switch dataType {

	case "varchar", "char", "text", "bpchar", "date", "timestamp", "time":
		return stringDefault

	case "int", "integer", "tinyint", "smallint", "mediumint", "int2", "int4":
		return "0"

	case "bool", "boolean":
		return "true"

	case "bigint", "int8":
		return "0"

	case "decimal", "dec":
		return "0"

	case "float4", "float8", "numeric":
		return "0"

	case "double":
		return "0"

	case "bytea":
		return stringDefault

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavascriptDataTypeForPostgres(dataType string, numericPrecision int) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "bpchar", "longtext":
		return "String"

	case "date", "year", "datetime":
		return "String"

	case "timestamp":
		return "String"

	case "time":
		return "String"

	case "int", "integer", "int2", "smallint", "int4":
		return "Number"

	case "bool", "boolean":
		return "Boolean"

	case "bigint", "int8":
		return "Number"

	case "decimal", "dec":
		return "Number"

	case "float4", "float8":
		return "Number"

	case "double":
		return "Number"

	case "binary", "varbinary", "blob", "tinyblob", "bytea":
		return "string"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaDataTypeForPostgres(dataType string, numericPrecision int) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar":
		return "String"

	case "date", "year", "datetime":
		return "Date"

	case "timestamp":
		return "Timestamp"

	case "time":
		return "Time"

	case "int", "integer", "int2", "int4", "smallint", "mediumint":
		return "Integer"

	case "bool", "boolean":
		return "Boolean"

	case "bigint", "int8":
		return "Long"

	case "decimal", "numeric":
		return "BigDecimal"

	case "float4", "float8":
		return "Float"

	case "double":
		return "Double"

	case "binary", "varbinary", "blob", "tinyblob", "bytea":
		return "byte[]"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoDataTypeForPostgres(dataType string) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar", "time":
		return "string"

	case "timestamp", "year", "datetime", "date":
		return "time.Time"

	case "int", "integer", "int2", "int4", "smallint":
		return "int"

	case "mediumint":
		return "int32"

	case "bool", "boolean":
		return "bool"

	case "bigint", "int8":
		return "int64"

	case "decimal", "numeric", "float8", "double":
		return "float64"

	case "float4":
		return "float32"

	case "binary", "varbinary", "blob", "tinyblob", "bytea":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpDataTypeForPostgres(dataType string, numericPrecision int) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar":
		return "string"

	case "date", "datetime", "year", "time":
		return "DateTime"

	case "timestamp":
		return "string"

	case "int", "int4", "integer", "mediumint":
		return "int"

	case "tinyint":
		if numericPrecision == 1 {
			return "bool"
		} else {
			return "byte"
		}

	case "smallint", "int2":
		return "short"

	case "bool", "boolean", "bit":
		return "bool"

	case "bigint", "int8":
		return "Int64"

	case "float4", "float8", "numeric":
		return "Decimal"

	case "double":
		return "Long"

	case "binary", "varbinary", "blob", "tinyblob", "bytea":
		return "Byte[]"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpFirstUnitTestValueForPostgres(dataType string) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar":
		return "A"

	case "date", "datetime", "year", "time":
		return "new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day)"

	case "timestamp":
		return "string"

	case "set", "enum", "geometry":
		return "AA"

	case "int", "integer", "tinyint", "smallint", "int2", "int4":
		return "2"

	case "bool", "boolean", "bit":
		return "true"

	case "bigint", "int8":
		return "1000"

	case "float4", "float8", "dec":
		return "3000F"

	case "numeric":
		return "new BigDecimal(5)"

	case "double":
		return "4000D"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "new byte[]{(byte) 8, (byte) 2}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoFirstUnitTestValueForPostgres(dataType string) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar":
		return "A"

	case "set", "enum", "geometry":
		return "AA"

	case "date":
		return "datatypes.Date"

	case "timestamp", "year", "datetime", "time":
		return "time.Time"

	case "int", "integer", "int2", "int4", "smallint":
		return "2"

	case "mediumint":
		return "2"

	case "bool", "boolean", "bit":
		return "true"

	case "bigint", "int8":
		return "33"

	case "decimal", "numeric", "float8", "double":
		return "44"

	case "float4":
		return "55"

	case "binary", "varbinary", "blob", "tinyblob", "bytea":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoSecondUnitTestValueForPostgres(dataType string) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar":
		return "B"

	case "set", "enum", "geometry":
		return "BB"

	case "date":
		return "datatypes.Date"

	case "timestamp", "year", "datetime", "time":
		return "time.Time"

	case "int", "integer", "int2", "int4", "smallint":
		return "3"

	case "mediumint":
		return "3"

	case "bool", "boolean", "bit":
		return "false"

	case "bigint", "int8":
		return "55"

	case "decimal", "numeric", "float8", "double":
		return "66"

	case "float4":
		return "77"

	case "binary", "varbinary", "blob", "tinyblob", "bytea":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaFirstUnitTestValueForPostgres(dataType string) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar":
		return "A"
	case "date", "datetime", "year":
		return "new Date(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"
	case "time":
		return "new Time(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "timestamp":
		return "new Timestamp(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "set", "enum", "geometry":
		return "AA"

	case "int", "integer", "tinyint", "smallint", "int2", "int4":
		return "2"

	case "bool", "boolean", "bit":
		return "true"

	case "bigint", "int8":
		return "1000L"

	case "float4", "float8":
		return "new Float(3)"

	case "dec":
		return "new BigDecimal(5)"

	case "numeric":
		return "new BigDecimal(6)"

	case "double":
		return "new Double(4)"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "new byte[]{(byte) 8, (byte) 2}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaSecondUnitTestValueForPostgres(dataType string) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar":
		return "B"
	case "date", "datetime", "year":
		return "new Date(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"
	case "time":
		return "new Time(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "timestamp":
		return "new Timestamp(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "set", "enum", "geometry":
		return "BB"

	case "int", "integer", "tinyint", "smallint", "int2", "int4":
		return "3"

	case "bool", "boolean", "bit":
		return "false"

	case "bigint", "int8":
		return "2000L"

	case "float4", "float8":
		return "new Float(33)"

	case "dec":
		return "new BigDecimal(55)"

	case "numeric":
		return "new BigDecimal(66)"

	case "double":
		return "new Double(44)"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "new byte[]{(byte) 88, (byte) 22}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpSecondUnitTestValueForPostgres(dataType string) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "bpchar":
		return "B"

	case "date", "datetime", "year", "time":
		return "new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)"

	case "timestamp":
		return "string"

	case "int", "integer", "int2", "int4", "smallint", "mediumint":
		return "3"

	case "bool", "boolean", "bit":
		return "false"

	case "bigint":
		return "2000"

	case "float4", "float8", "dec":
		return "3500F"

	case "numeric":
		return "new BigDecimal(55)"

	case "double":
		return "4500D"

	case "binary", "varbinary", "blob", "tinyblob", "bytea":
		return "new byte[]{(byte) 18, (byte) 12}"

	default:
		return "Unknown" + "_" + dataType
	}
}
