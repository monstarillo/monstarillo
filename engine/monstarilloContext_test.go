package engine

import (
	"testing"
)

func TestMonstarilloContext_GetColumn(t *testing.T) {
	context := ReadMonstrilloContext("../test-data/context.json")
	value := context.GetColumn("actor", "first_name")

	if value.TableName != "actor" {
		t.Errorf("Value was incorrect, got %s wanted actor", value.TableName)
	}
}

func TestMonstarilloContext_GetTable(t *testing.T) {
	context := ReadMonstrilloContext("../test-data/context.json")
	value := context.GetTable("actor")

	if value.TableName != "actor" {
		t.Errorf("Value was incorrect, got %s wanted actor", value.TableName)
	}
}

func TestMonstarilloContext_GetFkTableName(t *testing.T) {
	context := ReadMonstrilloContext("../test-data/context.json")
	value := context.GetFkTableName("film_actor", "actor_id")

	if value != "actor" {
		t.Errorf("Value was incorrect, got %s wanted actor", value)
	}
}

func TestMonstarilloContext_GetFkTableNamePlural(t *testing.T) {
	context := ReadMonstrilloContext("../test-data/context.json")
	value := context.GetFkTableNamePlural("film_actor", "actor_id")

	if value != "actors" {
		t.Errorf("Value was incorrect, got %s wanted actor", value)
	}
}

func TestReadMonstrilloContext(t *testing.T) {
	context := ReadMonstrilloContext("../test-data/context.json")
	value := len(context.Tables)
	if value != 12 {
		t.Errorf("Value was incorrect, got %d wanted 12", value)
	}
}
