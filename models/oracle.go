package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/sijms/go-ora/v2"
	go_ora "github.com/sijms/go-ora/v2"
	"strconv"
)

func ConnectOracleDB(user, password, server, service string, port int) {
	connStr := `(DESCRIPTION=
    (ADDRESS_LIST=
    	(LOAD_BALANCE=OFF)
      	(address=(protocol=tcp)(host=` + server + `)(port=` + strconv.Itoa(port) + `))
    )
    (CONNECT_DATA=
    	(SID=OraDoc)
        (SERVER=DEDICATED)
    )
    (SOURCE_ROUTE=yes)
    )`

	databaseUrl := go_ora.BuildJDBC(user, password, connStr, nil)

	db, err := sql.Open("oracle", databaseUrl)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected to Oracle")

	DB = db

}

type OracleColumn struct {
	ColumnName, Nullable, DataType, IdentityColumn, DataDefault sql.NullString
	DataPrecision, DataLength, DataScale                        sql.NullInt64
}

type OraclePrimaryKey struct {
	TableName, TableSchema, ConstraintName, Status, ColumnName string
	OrdinalPosition                                            int
}

func GetOracleTables(schema string) []Table {
	var tables []Table
	tableNames := GetTableNamesOracle()

	a := 0
	for range tableNames {

		var table Table
		table.DatabaseType = "oracle"
		table.TableName = tableNames[a]
		fmt.Println("Processing table: " + color.BlueString(tableNames[a]))
		primaryKeys := GetOraclePrimaryKeys(table.TableName, schema)

		table.ForeignKeys = GetOracleForeignKeys(table.TableName, schema)
		table.ReferencedForeignKeys = GetOracleReferencedForeignKeys(table.TableName, schema)
		table.Columns = GetColumnsOracle(tableNames[a], schema, primaryKeys, table.ForeignKeys)
		tables = append(tables, table)
		a++
	}

	return tables
}

func GetTableNamesOracle() []string {
	var tables []string

	stmt, err := DB.Prepare("SELECT TABLE_NAME FROM USER_TABLES")
	CheckError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	defer rows.Close()

	for rows.Next() {
		var t string

		err = rows.Scan(&t)
		if err != nil {
			CheckError(err)
		}
		tables = append(tables, t)
	}
	return tables
}

func GetColumnsOracle(tableName, schema string, primaryKeys []OraclePrimaryKey, fks []ForeignKey) []Column {
	var oracleColumns []OracleColumn
	var columns []Column

	stmt, err := DB.Prepare("select COLUMN_NAME, NULLABLE, DATA_TYPE, DATA_PRECISION, DATA_SCALE, IDENTITY_COLUMN, DATA_DEFAULT, DATA_LENGTH from sys.ALL_TAB_COLUMNS where owner = :1 and TABLE_NAME= :2")
	CheckError(err)
	defer stmt.Close()

	rows, err := stmt.Query(schema, tableName)

	for rows.Next() {
		var column OracleColumn
		err = rows.Scan(&column.ColumnName, &column.Nullable, &column.DataType, &column.DataPrecision, &column.DataScale, &column.IdentityColumn, &column.DataDefault, &column.DataLength)
		CheckError(err)
		oracleColumns = append(oracleColumns, column)
	}

	p := 0
	for range oracleColumns {
		var col Column
		col.ColumnName = oracleColumns[p].ColumnName.String
		col.TableName = tableName
		col.DatabaseType = "oracle"

		if oracleColumns[p].DataPrecision.Valid == true {
			col.NumericPrecision = int(oracleColumns[p].DataPrecision.Int64)
		}
		col.DataType = oracleColumns[p].DataType.String
		if oracleColumns[p].DataLength.Valid == true {
			col.CharacterMaximumLength = int(oracleColumns[p].DataLength.Int64)
		}

		if oracleColumns[p].DataScale.Valid == true {
			col.NumericScale = int(oracleColumns[p].DataScale.Int64)
		}

		if oracleColumns[p].Nullable.String == "Y" {
			col.IsNullable = true
		} else {
			col.IsNullable = false
		}

		if oracleColumns[p].IdentityColumn.String == "YES" {
			col.IsAutoIncrement = true
		} else {
			col.IsAutoIncrement = false
		}

		pk := 0
		col.IsPrimaryKey = false
		for range primaryKeys {
			if primaryKeys[pk].ColumnName == col.ColumnName {
				col.IsPrimaryKey = true
			}
			pk++
		}

		col.IsForeignKey = IsColumnForeignKey(col.ColumnName, fks)
		col.PkTableName = GetPkTableName(col.ColumnName, fks)
		col.PkColumnName = GetPkColumnName(col.ColumnName, fks)

		columns = append(columns, col)
		p++
	}

	return columns
}

func GetOraclePrimaryKeys(tableName, schema string) []OraclePrimaryKey {
	var primaryKeys []OraclePrimaryKey

	stmt, err := DB.Prepare("SELECT cols.table_name, cols.column_name, cols.position, cons.status, cons.CONSTRAINT_NAME  FROM all_constraints cons, all_cons_columns cols      WHERE cols.table_name = :1      and cons.owner = :2      AND cons.constraint_type = 'P'  AND cons.constraint_name =  cols.constraint_name      AND cons.owner = cols.owner      ORDER BY cols.table_name, cols.position")
	CheckError(err)
	defer stmt.Close()

	rows, err := stmt.Query(tableName, schema)

	for rows.Next() {
		var pk OraclePrimaryKey
		err = rows.Scan(&pk.TableName, &pk.ColumnName, &pk.OrdinalPosition, &pk.Status, &pk.ConstraintName)
		CheckError(err)
		primaryKeys = append(primaryKeys, pk)
	}

	return primaryKeys
}

func GetOracleForeignKeys(tableName, schema string) []ForeignKey {
	var fks []ForeignKey

	sqlStatement := "SELECT a.constraint_name, a.table_name, a.column_name,  " +
		"c_pk.table_name r_table_name, b.column_name r_column_name " +
		"FROM user_cons_columns a " +
		"JOIN user_constraints c ON a.owner = c.owner " +
		"AND a.constraint_name = c.constraint_name " +
		"JOIN user_constraints c_pk ON c.r_owner = c_pk.owner " +
		"AND c.r_constraint_name = c_pk.constraint_name " +
		"JOIN user_cons_columns b ON C_PK.owner = b.owner " +
		"AND  C_PK.CONSTRAINT_NAME = b.constraint_name AND b.POSITION = a.POSITION " +
		"WHERE c.constraint_type = 'R' and a.table_name = :1 " +
		"AND a.owner= :2"

	stmt, err := DB.Prepare(sqlStatement)
	CheckError(err)
	defer stmt.Close()

	rows, err := stmt.Query(tableName, schema)

	for rows.Next() {
		var fk ForeignKey
		err = rows.Scan(&fk.ConstraintName, &fk.FkTableName, &fk.FkColumnName, &fk.PkTableName, &fk.PkColumnName)
		CheckError(err)
		fks = append(fks, fk)
	}

	return fks
}

func GetOracleReferencedForeignKeys(tableName, schema string) []ForeignKey {
	var fks []ForeignKey

	sqlStatement := "SELECT a.constraint_name, a.table_name r_table_name, a.column_name r_column_name,  " +
		"c_pk.table_name , b.column_name  " +
		"FROM user_cons_columns a " +
		"JOIN user_constraints c ON a.owner = c.owner " +
		"AND a.constraint_name = c.constraint_name " +
		"JOIN user_constraints c_pk ON c.r_owner = c_pk.owner " +
		"AND c.r_constraint_name = c_pk.constraint_name " +
		"JOIN user_cons_columns b ON C_PK.owner = b.owner " +
		"AND  C_PK.CONSTRAINT_NAME = b.constraint_name AND b.POSITION = a.POSITION " +
		"WHERE c.constraint_type = 'R' and c_pk.table_name = :1 and a.owner = :2"

	stmt, err := DB.Prepare(sqlStatement)
	CheckError(err)
	defer stmt.Close()

	rows, err := stmt.Query(tableName, schema)

	for rows.Next() {
		var fk ForeignKey
		err = rows.Scan(&fk.ConstraintName, &fk.FkTableName, &fk.FkColumnName, &fk.PkTableName, &fk.PkColumnName)
		CheckError(err)
		fks = append(fks, fk)
	}

	return fks
}
