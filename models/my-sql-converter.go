package models

import (
	"strconv"
	"time"
)

func GetJavascriptDefaultValueForMySql(dataType string, numericPrecision int) string {
	stringDefault := "\"\""
	switch dataType {

	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return stringDefault

	case "date", "year", "datetime":
		return stringDefault

	case "set", "enum", "geometry":
		return stringDefault

	case "timestamp", "time":
		return stringDefault

	case "bit":
		if numericPrecision == 1 {
			return "Boolean"
		} else {
			return "0"
		}

	case "int", "integer", "tinyint", "smallint", "mediumint", "bigint", "decimal", "dec", "float", "double":
		return "0"

	case "bool", "boolean":
		return "true"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return stringDefault

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavascriptDataTypeForMySql(dataType string, numericPrecision int) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "String"

	case "date", "year", "datetime":
		return "String"

	case "set", "enum", "geometry":
		return "String"

	case "timestamp", "time":
		return "String"

	case "bit":
		if numericPrecision == 1 {
			return "Boolean"
		} else {
			return "Number"
		}

	case "int", "integer", "tinyint", "smallint", "mediumint", "bigint", "decimal", "dec", "float", "double":
		return "Number"

	case "bool", "boolean":
		return "Boolean"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "String"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaDataTypeForMySql(dataType string, numericPrecision int) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "String"

	case "date", "year", "datetime":
		return "Date"

	case "set", "enum", "geometry":
		return "String"

	case "timestamp":
		return "Timestamp"

	case "time":
		return "Time"

	case "bit":
		if numericPrecision == 1 {
			return "Boolean"
		} else {
			return "Integer"
		}

	case "int", "integer", "tinyint", "smallint", "mediumint":
		return "Integer"

	case "bool", "boolean":
		return "Boolean"

	case "bigint":
		return "Long"

	case "decimal", "dec":
		return "BigDecimal"

	case "float":
		return "Float"

	case "double":
		return "Double"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "byte[]"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoDataTypeForMySql(dataType string, numericPrecision int) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "string"

	case "date":
		return "time.Time"

	case "set", "enum", "geometry":
		return "string"

	case "time", "datetime", "year", "timestamp":
		return "time.Time"

	case "bit", "tinyint":
		if numericPrecision == 1 {
			return "bool"
		} else {
			return "int8"
		}

	case "json":
		return "datatypes.JSON"

	case "int", "integer":
		return "int"

	case "smallint", "mediumint":
		return "int8"

	case "bool", "boolean":
		return "bool"

	case "bigint":
		return "int64"

	case "decimal", "dec", "numeric", "double":
		return "float64"

	case "float", "real":
		return "float32"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpDataTypeForMySql(dataType string, numericPrecision int) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "string"

	case "date", "datetime", "year", "time":
		return "DateTime"

	case "timestamp":
		return "string"

	case "set", "enum", "geometry":
		return "Object"

	case "int", "integer", "mediumint":
		return "int"

	case "tinyint":
		return "byte"

	case "smallint":
		return "short"

	case "bool", "boolean", "bit":
		return "bool"

	case "bigint":
		return "Int64"

	case "float", "decimal", "dec":
		return "Decimal"

	case "double":
		return "Long"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "Byte[]"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpFirstUnitTestValueForMySql(dataType string) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "A"

	case "date", "datetime", "year", "time":
		return "new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day)"

	case "timestamp":
		return "string"

	case "set", "enum", "geometry":
		return "AA"

	case "int", "integer", "tinyint", "smallint", "mediumint":
		return "2"

	case "bool", "boolean", "bit":
		return "true"

	case "bigint":
		return "1000"

	case "float", "decimal", "dec":
		return "3000F"

	case "double":
		return "4000D"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "new byte[]{(byte) 8, (byte) 2}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoFirstUnitTestValueForMySql(dataType string, numericPrecision int) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "A"

	case "date":
		return "datatypes.Date"

	case "set", "enum", "geometry":
		return "AA"

	case "time", "datetime", "year", "timestamp":
		return "time.Time"

	case "bit", "tinyint":
		if numericPrecision == 1 {
			return "true"
		} else {
			return "1"
		}

	case "json":
		return "datatypes.JSON"

	case "int", "integer":
		return "2"

	case "smallint", "mediumint":
		return "2"

	case "bool", "boolean":
		return "true"

	case "bigint":
		return "22"

	case "decimal", "dec", "numeric", "double":
		return "33"

	case "float", "real":
		return "33"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoSecondUnitTestValueForMySql(dataType string, numericPrecision int) string {

	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "B"

	case "date":
		return "datatypes.Date"

	case "set", "enum", "geometry":
		return "BB"

	case "time", "datetime", "year", "timestamp":
		return "time.Time"

	case "bit", "tinyint":
		if numericPrecision == 1 {
			return "false"
		} else {
			return "3"
		}

	case "json":
		return "datatypes.JSON"

	case "int", "integer":
		return "3"

	case "smallint", "mediumint":
		return "3"

	case "bool", "boolean":
		return "false"

	case "bigint":
		return "33"

	case "decimal", "dec", "numeric", "double":
		return "44"

	case "float", "real":
		return "44"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaFirstUnitTestValueForMySql(dataType string) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "A"

	case "date", "datetime", "year":
		return "new Date(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"
	case "time":
		return "new Time(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "timestamp":
		return "new Timestamp(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "set", "enum", "geometry":
		return "AA"

	case "int", "integer", "tinyint", "smallint", "mediumint":
		return "2"

	case "bool", "boolean", "bit":
		return "true"

	case "bigint":
		return "1000L"

	case "float":
		return "new Float(3)"

	case "decimal", "dec":
		return "new BigDecimal(5)"

	case "double":
		return "new Double(4)"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "new byte[]{(byte) 8, (byte) 2}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaSecondUnitTestValueForMySql(dataType string) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "B"

	case "date", "datetime", "year":
		return "new Date(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"
	case "time":
		return "new Time(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "timestamp":
		return "new Timestamp(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "set", "enum", "geometry":
		return "BB"

	case "int", "integer", "tinyint", "smallint", "mediumint":
		return "3"

	case "bool", "boolean", "bit":
		return "false"

	case "bigint":
		return "2000L"

	case "float":
		return "new Float(33)"

	case "decimal", "dec":
		return "new BigDecimal(55)"

	case "double":
		return "new Double(44)"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "new byte[]{(byte) 88, (byte) 22}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpSecondUnitTestValueForMySql(dataType string) string {
	switch dataType {
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "B"

	case "date", "datetime", "year", "time":
		return "new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)"

	case "timestamp":
		return "string"

	case "set", "enum", "geometry":
		return "BB"

	case "int", "integer", "tinyint", "smallint", "mediumint":
		return "3"

	case "bool", "boolean", "bit":
		return "false"

	case "bigint":
		return "2000"

	case "float", "decimal", "dec":
		return "3500F"

	case "double":
		return "4500D"

	case "binary", "varbinary", "blob", "tinyblob", "mediumblob":
		return "new byte[]{(byte) 18, (byte) 12}"

	default:
		return "Unknown" + "_" + dataType
	}
}
