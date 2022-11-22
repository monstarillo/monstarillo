package models

import (
	"errors"
	"testing"
)

func getTable(tableName string) (Table, error) {
	tables := ReadTables("../test-data/tables.json")
	for _, t := range tables {
		if t.TableName == tableName {
			return t, nil
		}
	}

	var tbl Table
	return tbl, errors.New("Table not found")
}
func TestCamelCase(t *testing.T) {
	table := NewTable("actor", "mysql")
	camelCase := table.GetCamelCaseTableName()
	if camelCase != "actor" {
		t.Errorf("Camel case was incorrect, got: %s, want: %s", camelCase, "actor")
	}
}

func TestPascalCase(t *testing.T) {
	table := NewTable("actor", "mysql")
	camelCase := table.GetPascalCaseTableName()
	if camelCase != "Actor" {
		t.Errorf("Pascal case was incorrect, got: %s, want: %s", camelCase, "Actor")
	}
}

func TestPascalCaseEf(t *testing.T) {
	table := NewTable("actors", "mysql")
	camelCase := table.GetPascalCaseTableNameEF()
	if camelCase != "Actor" {
		t.Errorf("Pascal case was incorrect, got: %s, want: %s", camelCase, "Actor")
	}
}

func TestCamelCaseEf(t *testing.T) {
	table := NewTable("actors", "mysql")
	camelCase := table.GetCamelCaseTableNameEF()
	if camelCase != "actor" {
		t.Errorf("Pascal case was incorrect, got: %s, want: %s", camelCase, "actor")
	}
}

func GetPersonTable() Table {
	t := NewTable("Person", "mysql")
	id_col := NewColumn("id", "mysql", "Person")
	id_col.IsPrimaryKey = true
	id_col.IsNullable = false
	id_col.DataType = "int"
	t.AddColumn(id_col)

	fName := NewColumn("fname", "mysql", "Person")
	fName.IsNullable = true
	fName.DataType = "varchar"
	t.AddColumn((fName))

	lName := NewColumn("lname", "mysql", "Person")
	lName.IsNullable = true
	lName.DataType = "varchar"
	t.AddColumn((lName))

	startDate := NewColumn("start-date", "mysql", "Person")
	startDate.DataType = "datetime"

	t.AddColumn(startDate)
	return t
}

func TestPrimaryKeyColumns(t *testing.T) {
	table := GetPersonTable()
	primaryColumns := table.GetPrimaryColumns()

	if len(primaryColumns) != 1 {
		t.Errorf("Primary column count was incorrect, got : %d, want 1", len(primaryColumns))
	}
}

func TestNullableColumns(t *testing.T) {
	table := GetPersonTable()
	columns := table.GetNullableColumns()

	if len(columns) != 2 {
		t.Errorf("Primary column count was incorrect, got : %d, want 2", len(columns))
	}
}

func TestTable_HasCompositePrimaryKey(t *testing.T) {
	table := GetPersonTable()
	hasCompositePrimaryKey := table.HasCompositePrimaryKey()

	if hasCompositePrimaryKey != false {
		t.Errorf("Has composite primary key was incorrect, got : %t, want false", hasCompositePrimaryKey)
	}
}

func TestTable_HasAutoIncrementColumn_false(t *testing.T) {
	table := GetPersonTable()
	hasAutoIncrementColumn := table.HasAutoIncrementColumn()

	if hasAutoIncrementColumn != false {
		t.Errorf("Has auto increment Column was incorrect, got : %t, want false", hasAutoIncrementColumn)
	}
}

func TestTable_HasAutoIncrementColumn_true(t *testing.T) {
	table, _ := getTable("address")
	hasAutoIncrementColumn := table.HasAutoIncrementColumn()

	if hasAutoIncrementColumn != true {
		t.Errorf("Has auto increment Column was incorrect, got : %t, want true", hasAutoIncrementColumn)
	}
}

func TestTable_HasJavascriptStringColumn(t *testing.T) {
	table := GetPersonTable()
	value := table.HasJavascriptStringColumn()

	if value != true {
		t.Errorf("Value was incorrect, got : %t, want true", value)
	}
}

func TestTable_HasJavascriptNumberColumn(t *testing.T) {
	table := GetPersonTable()
	value := table.HasJavascriptNumberColumn()

	if value != true {
		t.Errorf("Value was incorrect, got : %t, want true", value)
	}
}

func TestTable_HasDateColumn(t *testing.T) {
	table := GetPersonTable()
	value := table.HasDateColumn()

	if value != true {
		t.Errorf("Value was incorrect, got : %t, want true", value)
	}
}

func TestTable_GetFkTableNameForColumn(t *testing.T) {
	table, _ := getTable("film_actor")
	fkTableName := table.GetFkTableNameForColumn("actor_id")

	if fkTableName != "actor" {
		t.Errorf("Value was incorrect, got %s, want actor", fkTableName)
	}

}

func TestTable_GetCamelCaseTableName(t *testing.T) {
	table, _ := getTable("film_actor")
	value := table.GetCamelCaseTableName()

	if value != "filmActor" {
		t.Errorf("Value was incorrect, got %s, want filmActor", value)
	}
}

func TestTable_GetCamelCaseTableNamePlural(t *testing.T) {
	table, _ := getTable("film_actor")
	value := table.GetCamelCaseTableNamePlural()

	if value != "filmActors" {
		t.Errorf("Value was incorrect, got %s, want filmActors", value)
	}
}

func TestTable_GetTableName(t *testing.T) {
	table, _ := getTable("film_actor")
	value := table.GetTableName()

	if value != "film_actor" {
		t.Errorf("Value was incorrect, got %s, want film_actor", value)
	}
}

func TestTable_GetPascalCaseTableName(t *testing.T) {
	table, _ := getTable("film_actor")
	value := table.GetPascalCaseTableNamePlural()

	if value != "FilmActors" {
		t.Errorf("Value was incorrect, got %s, want FilmActors", value)
	}
}

func TestTable_GetColumnListWithCSharpTypes(t *testing.T) {
	table, _ := getTable("film_actor")
	value := table.GetColumnListWithCSharpTypes()

	if value != "short actorId, short filmId, string lastUpdate " {
		t.Errorf("Value was incorrect, got '%s', want 'short actorId, short filmId, string lastUpdate '", value)
	}
}
