package models

import (
	"testing"
)

func ConnectToTestDBPostgres() {
	settings := ReadIntegrationTestSettings("../test-data/IntegrationTestSettings.json")
	ConnectPostgresDB(settings.PostgresUser, settings.PostgresPassword, settings.PostgresDbName, settings.PostgresHost, settings.PostgresPort)

}

func TestPostgres_GetTableNamesPostgres(t *testing.T) {
	ConnectToTestDBPostgres()
	tables := GetTableNamesPostgres("public")

	CloseDB()

	if len(tables) != 6 {
		t.Errorf("Value was incorrect, got %d, wanted 5", len(tables))
	}
}

func TestGetPostgresTables(t *testing.T) {
	ConnectToTestDBPostgres()
	tables := GetPostgresTables("public", "postgres")

	CloseDB()

	if len(tables) != 6 {
		t.Errorf("Value was incorrect, got %d, wanted 5", len(tables))
	}
}

func TestGetPostgresPrimaryKeys(t *testing.T) {
	ConnectToTestDBPostgres()
	pk := GetPostgresPrimaryKeys("employees", "public")

	CloseDB()

	if pk[0] != "emp_no" {
		t.Errorf("Value was incorrect, got %s, wanted emp_no", pk[0])
	}
}

func TestGetPostgresColumns(t *testing.T) {
	ConnectToTestDBPostgres()

	columns := GetPostgresColumns("employees", "public", "postgres", nil)
	CloseDB()

	if columns[0].ColumnName != "emp_no" {
		t.Errorf("Value was incorrect, got %s, wanted emp_no", columns[0].ColumnName)
	}
	if columns[0].DataType != "int4" {
		t.Errorf("Value was incorrect, got %s, wanted int4", columns[0].DataType)
	}

}
