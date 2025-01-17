package models

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/fatih/color"
	_ "github.com/lib/pq"
	"log"
)

func ConnectMsSqlServerDB(user, password, database string, port string) {

	connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, database)

	db, connectionError := sql.Open("mssql", connectionString)
	if connectionError != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
	}

	DB = db
}

func GetMssqlTables(schema, database string) []Table {
	var tables []Table
	tableNames := GetTableNamesMssql(database, schema)

	a := 0
	for range tableNames {

		var table Table
		table.DatabaseType = "mssql"
		table.TableName = tableNames[a]
		fmt.Println("Processing table: " + color.BlueString(tableNames[a]))
		primaryKeys := GetMssqlPrimaryKeys(table.TableName, schema)
		identityColumns := GetMssqlIdentityColumns(table.TableName)

		fmt.Println(primaryKeys)

		table.ForeignKeys = GetMssqlForeignKeys(table.TableName, primaryKeys, identityColumns)
		table.ReferencedForeignKeys = GetMssqlReferencedForeignKeys(table.TableName, primaryKeys, identityColumns)
		table.Columns = GetMssqlColumns(tableNames[a], schema, database, table.ForeignKeys)
		tables = append(tables, table)
		a++
	}

	return tables
}

func GetTableNamesMssql(database string, schema string) []string {
	var tables []string

	results, err := DB.Query("select TABLE_NAME from INFORMATION_SCHEMA.tables where TABLE_CATALOG = '" + database + "' AND TABLE_SCHEMA ='" + schema + "' AND table_type='BASE TABLE'")
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

func GetMssqlForeignKeys(tableName string, primaryKeys []string, identityColumns []string) []ForeignKey {
	var fks []ForeignKey

	sqlStatement := "SELECT " +
		"f.name AS constraint_name " +
		", OBJECT_NAME(f.parent_object_id) AS table_name " +
		", COL_NAME(fc.parent_object_id, fc.parent_column_id) AS column_name " +
		", OBJECT_NAME(f.referenced_object_id) AS referenced_table " +
		", COL_NAME(fc.referenced_object_id, fc.referenced_column_id) AS referenced_column_name " +
		"FROM sys.foreign_keys AS f " +
		"INNER JOIN sys.foreign_key_columns AS fc " +
		"ON f.object_id = fc.constraint_object_id " +
		"WHERE f.parent_object_id = OBJECT_ID('" + tableName + "')"

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

		fk.PkColumn = GetColumnMssql(fk.PkColumnName, fk.PkTableName, primaryKeys, identityColumns)
		fk.FkColumn = GetColumnMssql(fk.FkColumnName, fk.FkTableName, primaryKeys, identityColumns)
		fks = append(fks, fk)

	}
	return fks
}

func GetMssqlColumns(tableName, schema, database string, fks []ForeignKey) []Column {
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

	primaryKeys := GetMssqlPrimaryKeys(tableName, schema)
	identityColumns := GetMssqlIdentityColumns(tableName)

	for results.Next() {
		var columnName string

		err = results.Scan(&columnName)
		if err != nil {
			CheckError(err)
		}

		column := GetColumnMssql(columnName, tableName, primaryKeys, identityColumns)
		column.IsForeignKey = IsColumnForeignKey(columnName, fks)
		column.PkTableName = GetPkTableName(column.ColumnName, fks)
		column.PkColumnName = GetPkColumnName(column.ColumnName, fks)
		columns = append(columns, column)

	}

	return columns

}

func GetColumnMssql(columnName, tableName string, primaryKeys []string, identityColumns []string) Column {
	var column Column

	var col PostgresColumn

	sqlStatement := "select ORDINAL_POSITION, IS_NULLABLE, DATA_TYPE, NUMERIC_PRECISION, CHARACTER_MAXIMUM_LENGTH, NUMERIC_SCALE   from INFORMATION_SCHEMA.COLUMNS where TABLE_NAME = '" + tableName + "' " +
		"and COLUMN_NAME = '" + columnName + "'"

	err := DB.QueryRow(sqlStatement).Scan(&col.OrdinalPosition, &col.IsNullable, &col.DataType, &col.NumericPrecision, &col.CharacterMaximumLength, &col.NumericScale)
	if err != nil {
		log.Printf(sqlStatement)
		CheckError(err)
	}

	column.ColumnName = columnName
	column.TableName = tableName
	column.OrdinalPosition = col.OrdinalPosition
	column.DataType = col.DataType.String
	column.DatabaseType = "mssql"
	if col.NumericPrecision != nil {
		column.NumericPrecision = *col.NumericPrecision
	}

	if col.CharacterMaximumLength != nil {
		column.CharacterMaximumLength = *col.CharacterMaximumLength
	}

	if col.NumericScale != nil {
		column.NumericScale = *col.NumericScale
	}

	if col.IsNullable.String == "YES" {
		column.IsNullable = true
	} else {
		column.IsNullable = false
	}

	if contains(primaryKeys, columnName) {
		column.IsPrimaryKey = true
	} else {
		column.IsPrimaryKey = false
	}

	if contains(identityColumns, columnName) {
		column.IsAutoIncrement = true
	} else {
		column.IsAutoIncrement = false
	}

	return column
}
func GetMssqlReferencedForeignKeys(tableName string, primaryKeys []string, identityColumns []string) []ForeignKey {
	var fks []ForeignKey

	sqlStatement := "SELECT " +
		"f.name AS constraint_name " +
		", OBJECT_NAME(f.parent_object_id) AS referenced_table " +
		", COL_NAME(fc.parent_object_id, fc.parent_column_id) AS referenced_column_name " +
		", OBJECT_NAME(f.referenced_object_id) AS table_name " +
		", COL_NAME(fc.referenced_object_id, fc.referenced_column_id) AS column_name " +
		"FROM sys.foreign_keys AS f " +
		"INNER JOIN sys.foreign_key_columns AS fc " +
		"ON f.object_id = fc.constraint_object_id " +
		"WHERE f.referenced_object_id = OBJECT_ID('" + tableName + "')"

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

		fk.PkColumn = GetColumnMssql(fk.PkColumnName, fk.PkTableName, primaryKeys, identityColumns)
		fk.FkColumn = GetColumnMssql(fk.FkColumnName, fk.FkTableName, primaryKeys, identityColumns)
		fks = append(fks, fk)

	}
	return fks
}
func GetMssqlIdentityColumns(tableName string) []string {
	var keys []string

	sqlStatement := "SELECT NAME AS COLUMN_NAME FROM     SYS.IDENTITY_COLUMNS " +
		"WHERE OBJECT_NAME(OBJECT_ID) = '" + tableName + "'"

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

func GetMssqlPrimaryKeys(tableName, schema string) []string {
	var keys []string

	sqlStatement := "SELECT COLUMN_NAME " +
		"FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE " +
		"WHERE table_name = '" + tableName + "' AND TABLE_SCHEMA = '" + schema + "' " +
		"AND OBJECTPROPERTY(OBJECT_ID(CONSTRAINT_SCHEMA + '.' + QUOTENAME(CONSTRAINT_NAME)), 'IsPrimaryKey') = 1 "

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
