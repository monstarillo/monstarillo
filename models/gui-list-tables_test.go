package models

import (
	"errors"
	"testing"
)

func getTestGuiListTables() GuiListTables {
	return ReadGuiTables("../test-data/gui-tables.json")
}

func getGuiTable(tableName string) (GuiListTable, error) {
	tables := getTestGuiListTables()

	for _, t := range tables.Tables {
		if t.TableName == tableName {
			return t, nil
		}
	}
	var table GuiListTable
	return table, errors.New("Table not found")

}

func TestGuiListTable_HasGuiControl_true(t *testing.T) {
	table, _ := getGuiTable("fool_date")
	value := table.HasGuiControl("BaseDatepicker")

	if value != true {
		t.Errorf("Value was incorrect, got %t, want true", value)
	}
}

func TestGuiListTable_HasGuiControl_false(t *testing.T) {
	table, _ := getGuiTable("fool_date")
	value := table.HasGuiControl("xxxxx")

	if value != false {
		t.Errorf("Value was incorrect, got %t, want false", value)
	}
}

func TestGuiListTable_GetEditGuiControlForColumn(t *testing.T) {
	table, _ := getGuiTable("fool_date")
	value := table.GetEditGuiControlForColumn("endDate")

	if value != "BaseDatepicker" {
		t.Errorf("Value was incorrect, got %s, want BaseDatepicker", value)
	}
}

func TestGuiListTable_GetEditSelectKeyForColumn(t *testing.T) {
	table, _ := getGuiTable("film_actor")
	value := table.GetEditSelectKeyForColumn("filmId")

	if value != "filmId" {
		t.Errorf("Value was incorrect, got %s, want filmId", value)
	}
}

func TestGuiListTable_GetEditSelectValueForColumn(t *testing.T) {
	table, _ := getGuiTable("film_actor")
	value := table.GetEditSelectValueForColumn("filmId")

	if value != "description" {
		t.Errorf("Value was incorrect, got %s, want description", value)
	}
}
