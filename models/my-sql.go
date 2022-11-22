package models

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Actor struct {
	first_name, last_name string
}

var DB *sql.DB

func ConnectDB(userName, password, dbname string) {
	connectionString := userName + ":" + password + "@" + dbname
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		return
	}
}

type MySqlColumn struct {
	OrdinalPosition                          int
	NumericPrecision, CharacterMaximumLength *int
	IsNullable, DataType, Extra, ColumnKey   string
}
type MySqlForeignKey struct {
	TableName, ColumnName, ConstraintName, ReferencedColumnName, ReferencedTableName string
}

func GetForeignKeys(database, tableName string) []ForeignKey {
	var fks []ForeignKey

	sqlStatement := "SELECT " +
		"CONSTRAINT_NAME, " +
		"TABLE_NAME, " +
		"COLUMN_NAME, " +
		"REFERENCED_TABLE_NAME, " +
		"REFERENCED_COLUMN_NAME " +
		"FROM " +
		"INFORMATION_SCHEMA.KEY_COLUMN_USAGE " +
		"WHERE " +
		"REFERENCED_TABLE_SCHEMA = \"" + database + "\" " +
		"AND TABLE_NAME = \"" + tableName + "\""

	// Execute the query
	results, err := DB.Query(sqlStatement)

	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var fk ForeignKey
		// for each row, scan the result into our tag composite object
		err = results.Scan(&fk.ConstraintName, &fk.FkTableName, &fk.FkColumnName, &fk.PkTableName, &fk.PkColumnName)
		if err != nil {
			log.Fatal(err)
		}

		fk.PkColumn = GetColumn(database, fk.PkColumnName, fk.PkTableName)
		fk.FkColumn = GetColumn(database, fk.FkColumnName, fk.FkTableName)
		fks = append(fks, fk)

	}
	return fks
}

func GetReferencedForeignKeys(database, tableName string) []ForeignKey {
	var fks []ForeignKey

	query := "SELECT " +
		"CONSTRAINT_NAME, " +
		"REFERENCED_TABLE_NAME as TABLE_NAME, " +
		"REFERENCED_COLUMN_NAME as COLUMN_NAME, " +
		"TABLE_NAME as REFERENCED_TABLE_NAME, " +
		"COLUMN_NAME as REFERENCED_COLUMN_NAME " +
		"FROM " +
		"INFORMATION_SCHEMA.KEY_COLUMN_USAGE " +
		"WHERE " +
		"REFERENCED_TABLE_SCHEMA = '" + database + "' " +
		"AND REFERENCED_TABLE_NAME = '" + tableName + "'"

	// Execute the query
	results, err := DB.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var fk ForeignKey
		// for each row, scan the result into our tag composite object
		err = results.Scan(&fk.ConstraintName, &fk.FkTableName, &fk.FkColumnName, &fk.PkTableName, &fk.PkColumnName)
		if err != nil {
			log.Fatal(err)
		}

		fk.PkColumn = GetColumn(database, fk.FkColumnName, fk.FkTableName)
		fk.FkColumn = GetColumn(database, fk.PkColumnName, fk.PkTableName)
		fks = append(fks, fk)

	}
	return fks
}

func GetColumn(schema, columnName, tableName string) Column {
	var column Column

	var col MySqlColumn
	// Execute the query
	err := DB.QueryRow("SELECT extra, ordinal_position, is_nullable, data_type, Numeric_precision, character_maximum_length, COLUMN_KEY FROM INFORMATION_SCHEMA.COLUMNS where TABLE_SCHEMA = \""+schema+"\" and column_name = \""+columnName+"\" and table_name= '"+tableName+"'").Scan(&col.Extra, &col.OrdinalPosition, &col.IsNullable, &col.DataType, &col.NumericPrecision, &col.CharacterMaximumLength, &col.ColumnKey)
	if err != nil {
		log.Printf("SELECT extra, ordinal_position, is_nullable, data_type, Numeric_precision, character_maximum_length, COLUMN_KEY FROM INFORMATION_SCHEMA.COLUMNS where TABLE_SCHEMA = \"" + schema + "\" and column_name = \"" + columnName + "\" and table_name= '" + tableName + "'")
		log.Fatal(err)
	}
	// and then print out the tag"s Name attribute
	//log.Printf(columnName + " " + col.DataType)
	column.ColumnName = columnName
	column.TableName = tableName
	column.OrdinalPosition = col.OrdinalPosition
	column.DataType = col.DataType
	column.DatabaseType = "mysql"
	column.IsForeignKey = false

	if col.NumericPrecision != nil {
		column.NumericPrecision = *col.NumericPrecision
	}

	if col.CharacterMaximumLength != nil {
		column.CharacterMaximumLength = *col.CharacterMaximumLength
	}

	if col.IsNullable == "YES" {
		column.IsNullable = true
	} else {
		column.IsNullable = false
	}

	if col.Extra == "auto_increment" {
		column.IsAutoIncrement = true
	}

	if col.ColumnKey == "PRI" {
		column.IsPrimaryKey = true
	}

	return column
}

func GetTableNames(schema string) []string {
	var tables []string
	//db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/sakila")
	// if there is an error opening the connection, handle it

	// Execute the query
	results, err := DB.Query("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES  WHERE TABLE_SCHEMA =\"" + schema + "\" and TABLE_TYPE =\"BASE TABLE\"")

	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var t string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&t)
		if err != nil {
			log.Fatal(err)
		}
		// and then print out the tag"s Name attribute
		//log.Printf(t)
		tables = append(tables, t)
	}
	return tables
}

func GetColumnNames(schema, tableName string) []string {
	var columns []string

	// Execute the query
	results, err := DB.Query("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS  WHERE TABLE_NAME =\"" + tableName + "\"" + " and TABLE_SCHEMA =\"" + schema + "\"")

	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var t string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&t)
		if err != nil {
			log.Fatal(err)
		}
		// and then print out the tag"s Name attribute

		columns = append(columns, t)
	}
	return columns
}

func IsColumnForeignKey(columnName string, keys []ForeignKey) bool {
	for _, f := range keys {
		if f.FkColumnName == columnName {
			return true

		}
	}
	return false
}

func GetPkTableName(columnName string, keys []ForeignKey) string {
	for _, f := range keys {
		if f.FkColumnName == columnName {
			return f.PkTableName
		}
	}
	return ""
}

func GetPkColumnName(columnName string, keys []ForeignKey) string {
	for _, f := range keys {
		if f.FkColumnName == columnName {
			return f.PkColumnName
		}
	}
	return ""
}

func GetTables(schema string) []Table {
	var tables []Table

	t := 0
	tableNames := GetTableNames(schema)
	for range tableNames {
		c := 0
		tbl := NewTable(tableNames[t], "mysql")
		tbl.ForeignKeys = GetForeignKeys(schema, tableNames[t])
		fmt.Println("Processing table: " + color.BlueString(tableNames[t]))
		tbl.ReferencedForeignKeys = GetReferencedForeignKeys(schema, tableNames[t])
		columnNames := GetColumnNames(schema, tableNames[t])
		for range columnNames {
			col := GetColumn(schema, columnNames[c], tableNames[t])
			col.IsForeignKey = IsColumnForeignKey(col.ColumnName, tbl.ForeignKeys)
			col.PkTableName = GetPkTableName(col.ColumnName, tbl.ForeignKeys)
			col.PkColumnName = GetPkColumnName(col.ColumnName, tbl.ForeignKeys)
			c++
			tbl.AddColumn(col)
		}

		tables = append(
			tables,
			tbl,
		)
		t++
	}

	return tables
}
