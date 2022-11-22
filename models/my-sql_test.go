package models

import "testing"

func ConnectToTestMySql() {
	settings := ReadIntegrationTestSettings("../test-data/IntegrationTestSettings.json")
	ConnectDB(settings.MySqlUser, settings.MySqlPassword, settings.MySqlDB)

}

func TestGetForeignKeys(t *testing.T) {
	ConnectToTestMySql()
	fks := GetForeignKeys("employees", "dept_manager")
	CloseDB()
	if len(fks) != 2 {
		t.Errorf("Value was incorrect, got %d, want 2", len(fks))
	}
}

func TestGetReferencedForeignKeys(t *testing.T) {
	ConnectToTestMySql()
	fks := GetReferencedForeignKeys("employees", "employees")
	CloseDB()
	if len(fks) != 4 {
		t.Errorf("Value was incorrect, got %d, want 4", len(fks))
	}
}

func TestGetColumn(t *testing.T) {
	ConnectToTestMySql()
	column := GetColumn("employees", "emp_no", "employees")
	CloseDB()
	if column.DataType != "int" {
		t.Errorf("Value was incorrect, got %s, want 4", column.DatabaseType)
	}

	if column.IsAutoIncrement != false {
		t.Errorf("Value was incorrect, got %v, want false", column.IsAutoIncrement)
	}

	if column.IsPrimaryKey != true {
		t.Errorf("Value was incorrect, got %v, want true", column.IsPrimaryKey)
	}
}

func TestGetTableNames(t *testing.T) {
	ConnectToTestMySql()
	tables := GetTableNames("employees")
	CloseDB()

	if len(tables) != 6 {
		t.Errorf("Value was incorrect, got %d, want 6", len(tables))

	}
}

func TestGetColumnNames(t *testing.T) {
	ConnectToTestMySql()
	columns := GetColumnNames("employees", "employees")
	CloseDB()

	if len(columns) != 6 {
		t.Errorf("Value was incorrect, got %d, want 6", len(columns))

	}
}

func TestIsColumnForeignKey(t *testing.T) {
	ConnectToTestMySql()
	fks := GetForeignKeys("employees", "dept_manager")
	value := IsColumnForeignKey("emp_no", fks)

	if value != true {
		t.Errorf("Value was incorrect, got %v, want true", value)
	}
	value = IsColumnForeignKey("from_date", fks)

	if value != false {
		t.Errorf("Value was incorrect, got %v, want false", value)
	}

	CloseDB()
}

func TestGetPkTableName(t *testing.T) {
	ConnectToTestMySql()
	fks := GetForeignKeys("employees", "dept_manager")
	value := GetPkTableName("emp_no", fks)

	if value != "employees" {
		t.Errorf("Value was incorrect, got %s, want true", value)
	}

	CloseDB()

}

func TestGetPkColumnName(t *testing.T) {
	ConnectToTestMySql()
	fks := GetForeignKeys("employees", "dept_manager")
	value := GetPkColumnName("emp_no", fks)

	if value != "emp_no" {
		t.Errorf("Value was incorrect, got %s, want emp_no", value)
	}

	CloseDB()

}

func TestGetTables(t *testing.T) {
	ConnectToTestMySql()
	tables := GetTables("employees")
	value := len(tables)

	if value != 6 {
		t.Errorf("Value was incorrect, got %d, want 6", value)
	}

	CloseDB()

}
