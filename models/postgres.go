package models

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

func ConnectPostgresDB(user, password, dbname, host string, port int) {
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", connection)
	CheckError(err)

	// close database
	//defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	//fmt.Println("Connected!")

	DB = db
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetTableNamesPostgres(schema string) []string {
	var tables []string

	results, err := DB.Query("SELECT table_name from information_schema.tables  where table_schema ='" + schema + "' ")
	if err != nil {
		CheckError(err)
	}

	for results.Next() {
		var t string

		err = results.Scan(&t)
		if err != nil {
			CheckError(err)
		}
		tables = append(tables, t)
	}
	return tables
}

func GetPostgresTables(schema, database string) []Table {
	var tables []Table
	tableNames := GetTableNamesPostgres(schema)

	a := 0
	for range tableNames {

		var table Table
		table.DatabaseType = "postgres"
		table.TableName = tableNames[a]
		fmt.Println("Processing table: " + color.BlueString(tableNames[a]))
		primaryKeys := GetPostgresPrimaryKeys(table.TableName, schema)

		table.ForeignKeys = GetPostgresForeignKeys(schema, database, table.TableName, primaryKeys)
		table.ReferencedForeignKeys = GetPostgresReferencedForeignKeys(schema, database, table.TableName, primaryKeys)
		table.Columns = GetPostgresColumns(tableNames[a], schema, database, table.ForeignKeys)
		tables = append(tables, table)
		a++
	}

	return tables
}

type PostgresColumn struct {
	OrdinalPosition                                        int
	NumericPrecision, CharacterMaximumLength, NumericScale *int
	IsNullable, DataType, IsIdentity, ColumnDefault        string
}

type PostgresPrimaryKey struct {
	TableName, TableSchema, ConstraintName, ColumnName string
	OrdinalPosition                                    int
}

func GetPostgresPrimaryKeys(tableName, schema string) []string {
	var keys []string

	sqlStatement := "select  " +
		"kcu.column_name as key_column " +
		"from information_schema.table_constraints tco " +
		"join information_schema.key_column_usage kcu " +
		"on kcu.constraint_name = tco.constraint_name " +
		"and kcu.constraint_schema = tco.constraint_schema " +
		"and kcu.constraint_name = tco.constraint_name " +
		"where tco.constraint_type = 'PRIMARY KEY' and kcu.table_name = '" + tableName + "' " +
		" and kcu.table_schema = '" + schema + "' " +
		"order by kcu.table_schema, " +
		"kcu.table_name "

	//fmt.Println(sqlStatement)
	// Execute the query
	results, err := DB.Query(sqlStatement)

	if err != nil {
		CheckError(err)
	}

	for results.Next() {
		var pk string

		err = results.Scan(&pk)
		if err != nil {
			CheckError(err)
		}

		keys = append(keys, pk)

	}

	return keys

}

func GetPostgresColumns(tableName, schema, database string, fks []ForeignKey) []Column {
	var columns []Column

	sqlStatement := "select  column_name from information_schema.columns " +
		" where table_schema = '" + schema + "' " +
		" and table_catalog = '" + database + "' " +
		" and table_name = '" + tableName + "'"

	// Execute the query
	results, err := DB.Query(sqlStatement)

	if err != nil {
		CheckError(err)
	}

	primaryKeys := GetPostgresPrimaryKeys(tableName, schema)

	for results.Next() {
		var columnName string

		err = results.Scan(&columnName)
		if err != nil {
			CheckError(err)
		}

		column := GetColumnPostgres(schema, columnName, tableName, database, primaryKeys)
		column.IsForeignKey = IsColumnForeignKey(columnName, fks)
		column.PkTableName = GetPkTableName(column.ColumnName, fks)
		column.PkColumnName = GetPkColumnName(column.ColumnName, fks)
		columns = append(columns, column)

	}

	return columns

}
func GetColumnPostgres(schema, columnName, tableName, database string, primaryKeys []string) Column {
	var column Column

	var col PostgresColumn

	sqlStatement := "select  is_identity, ordinal_position, is_nullable, udt_name, numeric_precision, character_maximum_length, column_default, numeric_scale from information_schema.columns " +
		" where table_schema = '" + schema + "' " +
		" and table_catalog = '" + database + "' " +
		" and table_name = '" + tableName + "'" +
		" and column_name = '" + columnName + "'"

	//fmt.Println("Get Column ---")
	//fmt.Println(sqlStatement)
	// Execute the query
	err := DB.QueryRow(sqlStatement).Scan(&col.IsIdentity, &col.OrdinalPosition, &col.IsNullable, &col.DataType, &col.NumericPrecision, &col.CharacterMaximumLength, &col.NumericScale, &col.NumericScale)
	if err != nil {
		log.Printf(sqlStatement)
		CheckError(err)
	}
	// and then print out the tag"s Name attribute
	//log.Printf(columnName + " " + col.DataType)
	column.ColumnName = columnName
	column.TableName = tableName
	column.OrdinalPosition = col.OrdinalPosition
	column.DataType = col.DataType
	column.DatabaseType = "postgres"

	if col.NumericPrecision != nil {
		column.NumericPrecision = *col.NumericPrecision
	}

	if col.CharacterMaximumLength != nil {
		column.CharacterMaximumLength = *col.CharacterMaximumLength
	}

	if col.NumericScale != nil {
		column.NumericScale = *col.NumericScale
	}

	if col.IsNullable == "YES" {
		column.IsNullable = true
	} else {
		column.IsNullable = false
	}

	if col.IsIdentity == "YES" {
		column.IsAutoIncrement = true
	} else if strings.Contains(strings.ToUpper(col.ColumnDefault), "NEXTVAL") {
		column.IsAutoIncrement = true
	} else {
		column.IsAutoIncrement = false
	}

	if contains(primaryKeys, columnName) {
		column.IsPrimaryKey = true
	} else {
		column.IsPrimaryKey = false
	}

	return column
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func GetPostgresForeignKeys(schema, database, tableName string, primaryKeys []string) []ForeignKey {
	var fks []ForeignKey

	sqlStatement := "SELECT  " +
		"tc.constraint_name, " +
		"ccu.table_name   AS foreign_table_name, " +
		"ccu.column_name  AS foreign_column_name, " +
		"tc.table_name, " +
		"kcu.column_name " +
		"FROM information_schema.table_constraints AS tc " +
		"JOIN information_schema.key_column_usage AS kcu " +
		"ON tc.constraint_name = kcu.constraint_name " +
		"AND tc.table_schema = kcu.table_schema " +
		"JOIN information_schema.constraint_column_usage AS ccu " +
		"ON ccu.constraint_name = tc.constraint_name " +
		"AND ccu.table_schema = tc.table_schema " +
		"WHERE tc.constraint_type = 'FOREIGN KEY' " +
		"AND tc.table_name = '" + tableName + "' " +
		"AND tc.table_schema = '" + schema + "'"

	//fmt.Println(sqlStatement)
	// Execute the query
	results, err := DB.Query(sqlStatement)

	if err != nil {
		CheckError(err)
	}

	for results.Next() {
		var fk ForeignKey
		// for each row, scan the result into our tag composite object
		err = results.Scan(&fk.ConstraintName, &fk.FkTableName, &fk.FkColumnName, &fk.PkTableName, &fk.PkColumnName)
		if err != nil {
			CheckError(err)
		}

		fk.PkColumn = GetColumnPostgres(schema, fk.PkColumnName, fk.PkTableName, database, primaryKeys)
		fk.FkColumn = GetColumnPostgres(schema, fk.FkColumnName, fk.FkTableName, database, primaryKeys)
		fks = append(fks, fk)

	}
	return fks
}

func GetPostgresReferencedForeignKeys(schema, database, tableName string, primaryKeys []string) []ForeignKey {
	var fks []ForeignKey

	sqlStatement := "SELECT tc.constraint_name, " +
		"tc.table_name  AS foreign_table_name, " +
		"kcu.column_name AS foreign_column_name,  " +
		"ccu.table_name  , " +
		"ccu.column_name  " +
		"FROM information_schema.table_constraints AS tc " +
		"JOIN information_schema.key_column_usage AS kcu " +
		"ON tc.constraint_name = kcu.constraint_name " +
		"AND tc.table_schema = kcu.table_schema " +
		"JOIN information_schema.constraint_column_usage AS ccu " +
		"ON ccu.constraint_name = tc.constraint_name " +
		"AND ccu.table_schema = tc.table_schema " +
		"WHERE tc.constraint_type = 'FOREIGN KEY' " +
		"AND ccu.table_name = '" + tableName + "' " +
		"AND ccu.table_schema = '" + schema + "'"

	// Execute the query
	results, err := DB.Query(sqlStatement)

	if err != nil {
		CheckError(err)
	}

	for results.Next() {
		var fk ForeignKey
		// for each row, scan the result into our tag composite object
		err = results.Scan(&fk.ConstraintName, &fk.FkTableName, &fk.FkColumnName, &fk.PkTableName, &fk.PkColumnName)
		if err != nil {
			CheckError(err)
		}

		fk.PkColumn = GetColumnPostgres(schema, fk.PkColumnName, fk.PkTableName, database, primaryKeys)
		fk.FkColumn = GetColumnPostgres(schema, fk.FkColumnName, fk.FkTableName, database, primaryKeys)
		fks = append(fks, fk)

	}
	return fks
}
