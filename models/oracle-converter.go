package models

import (
	"strconv"
	"time"
)

func GetJavascriptDefaultValueForOracle(dataType string, numericPrecision int) string {
	stringDefault := "\"\""
	switch dataType {

	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "DATE", "TIMESTAMP", "NCHAR", "NVARCHAR2":
		return stringDefault

	case "BINARY_INTEGER",
		"NATURAL",
		"NATURALN",
		"PLS_INTEGER",
		"POSITIVE",
		"POSITIVEN",
		"INT",
		"INTEGER",
		"SMALLINT",
		"DOUBLE PRECISION", "FLOAT",

		"NUMBER", "DEC", "DECIMAL", "NUMERIC":
		return "0"

	case "BOOLEAN":
		return "true"

	case "RAW",
		"LONG RAW",
		"BFILE",
		"BLOB":

		return stringDefault

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavascriptDataTypeForOracle(dataType string, numericPrecision int) string {

	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "String"

	case "DATE", "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "String"

	case "BINARY_INTEGER", "NATURAL", "NATURALN",
		"PLS_INTEGER",
		"POSITIVE",
		"POSITIVEN",
		"INT",
		"INTEGER",
		"SMALLINT",
		"DOUBLE PRECISION", "FLOAT",
		"NUMBER", "DEC", "DECIMAL", "NUMERIC":
		return "Number"

	case "BOOLEAN":
		return "Boolean"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "string"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaDataTypeForOracle(dataType string, numericPrecision int) string {

	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "String"

	case "DATE", "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "Date"

	case "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER", "SMALLINT":
		return "Integer"

	case "BOOLEAN":
		return "Boolean"

	case "DEC", "DECIMAL", "NUMBER", "NUMERIC":
		return "BigDecimal"

	case "REAL", "float8":
		return "Float"

	case "DOUBLE PRECISION", "FLOAT":
		return "Double"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "byte[]"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoDataTypeForOracle(dataType string) string {

	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "string"

	case "DATE":
		return "datatypes.Date"

	case "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "time.Time"

	case "SMALLINT", "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER":
		return "int"

	case "BOOLEAN":
		return "bool"

	case "DOUBLE PRECISION", "FLOAT", "REAL", "NUMBER", "DEC", "DECIMAL", "NUMERIC":
		return "float64"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpDataTypeForOracle(dataType string, numericPrecision int) string {
	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "string"

	case "DATE":
		return "DateTime"

	case "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "string"

	case "SMALLINT", "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER":
		return "int"

	case "BOOLEAN":
		return "bool"

	case "DOUBLE PRECISION", "FLOAT", "NUMBER", "DEC", "DECIMAL", "NUMERIC":
		return "Decimal"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "Byte[]"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpFirstUnitTestValueForOracle(dataType string) string {
	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "A"

	case "DATE":
		return "new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day)"

	case "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "string"

	case "SMALLINT", "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER":
		return "2"

	case "BOOLEAN":
		return "true"

	case "REAL":
		return "3000F"

	case "NUMBER", "DEC", "DECIMAL", "NUMERIC", "DOUBLE PRECISION", "FLOAT":
		return "new BigDecimal(5)"

	case "double":
		return "4000D"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "new byte[]{(byte) 8, (byte) 2}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoFirstUnitTestValueForOracle(dataType string) string {

	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "A"

	case "DATE":
		return "datatypes.Date"

	case "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "time.Time"

	case "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER", "SMALLINT", "FLOAT", "NUMBER":
		return "2"

	case "BOOL":
		return "true"

	case "DEC", "DECIMAL", "NUMERIC", "DOUBLE PRECISION":
		return "44"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetGoSecondUnitTestValueForOracle(dataType string) string {

	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "B"

	case "DATE":
		return "datatypes.Date"

	case "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "time.Time"

	case "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER", "SMALLINT", "FLOAT", "NUMBER":
		return "3"

	case "BOOLEAN":
		return "false"

	case "DEC", "DECIMAL", "NUMERIC", "DOUBLE PRECISION":
		return "66"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "[]byte"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaFirstUnitTestValueForOracle(dataType string) string {
	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "A"
	case "DATE":
		return "new Date(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"
	//case "time":
	//	return "new Time(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "new Timestamp(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER", "SMALLINT", "NUMBER":
		return "2"

	case "BOOLEAN":
		return "true"

	case "REAL":
		return "new Float(3)"

	case "DEC", "DECIMAL", "NUMERIC":
		return "new BigDecimal(5)"

	case "numeric":
		return "new BigDecimal(6)"

	case "DOUBLE PRECISION", "FLOAT":
		return "new Double(4)"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "new byte[]{(byte) 8, (byte) 2}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetJavaSecondUnitTestValueForOracle(dataType string) string {
	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "B"
	case "DATE":
		return "new Date(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"
	//case "time":
	//	return "new Time(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "new Timestamp(" + strconv.FormatInt(time.Now().UnixMilli()/int64(time.Millisecond), 10) + ")"

	case "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER", "SMALLINT":
		return "3"

	case "BOOLEAN":
		return "false"

	case "REAL":
		return "new Float(33)"

	case "DEC", "DECIMAL", "NUMBER", "NUMERIC":
		return "new BigDecimal(55)"

	case "DOUBLE PRECISION", "FLOAT":
		return "new Double(44)"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "new byte[]{(byte) 88, (byte) 22}"

	default:
		return "Unknown" + "_" + dataType
	}
}

func GetCSharpSecondUnitTestValueForOracle(dataType string) string {
	switch dataType {
	case "VARCHAR", "VARCHAR2", "LONG", "CLOB", "NCLOB", "CHAR", "CHARACTER", "NCHAR", "NVARCHAR2":
		return "B"

	case "DATE":
		return "new DateTime(DateTime.Now.AddDays(1).Year, DateTime.Now.AddDays(1).Month, DateTime.Now.AddDays(1).Day)"

	case "TIMESTAMP", "TIMESTAMP WITH TZ", "TIMESTAMP WITH LOCAL TZ":
		return "string"

	case "BINARY_INTEGER", "NATURAL", "NATURALN", "PLS_INTEGER", "POSITIVE", "POSITIVEN", "INT", "INTEGER", "SMALLINT":
		return "3"

	case "BOOLEAN":
		return "false"

	case "REAL":
		return "3500F"

	case "DEC", "DECIMAL", "NUMBER", "NUMERIC":
		return "new BigDecimal(55)"

	case "DOUBLE PRECISION", "FLOAT":
		return "4500D"

	case "RAW", "LONG RAW", "BFILE", "BLOB":
		return "new byte[]{(byte) 18, (byte) 12}"

	default:
		return "Unknown" + "_" + dataType
	}
}
