package engine

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func CopyTestTablesAndContextJson() {
	_, err := copyFile("../test-data/tables.json", "tables.json")
	if err != nil {
		return
	}
	_, err = copyFile("../test-data/context.json", "context.json")
	if err != nil {
		return
	}
}

func TestGetTable(t *testing.T) {
	CopyTestTablesAndContextJson()
	table := getTable("address")
	if table.TableName != "address" {
		t.Errorf("Value was incorrect, got %s, wanted address", table.TableName)
	}
}

func TestGetTableFirstPk(t *testing.T) {
	CopyTestTablesAndContextJson()
	pk := getTableFirstPk("address")
	if pk != "address_id" {
		t.Errorf("Value was incorrect, got %s, wanted address_id", pk)
	}
}

func TestGetTableSecondColumn(t *testing.T) {
	CopyTestTablesAndContextJson()
	pk := getTableSecondColumn("address")
	if pk != "address" {
		t.Errorf("Value was incorrect, got %s, wanted address", pk)
	}
}

func TestMakePlural(t *testing.T) {
	pk := makePlural("address")
	if pk != "addresses" {
		t.Errorf("Value was incorrect, got %s, wanted addresses", pk)
	}
}

func TestMakeCamelCase(t *testing.T) {
	pk := makeCamelCase("Address")
	if pk != "address" {
		t.Errorf("Value was incorrect, got %s, wanted address", pk)
	}
}

func TestMakePascalCase(t *testing.T) {
	pk := makePascalCase("address")
	if pk != "Address" {
		t.Errorf("Value was incorrect, got %s, wanted address", pk)
	}
}

func TestGetTag(t *testing.T) {
	CopyTestTablesAndContextJson()
	context := ReadMonstrilloContext("../test-data/context.json")
	tag := getTag(context.Tags, "Namespace")
	if tag != "Foo" {
		t.Errorf("Value was incorrect, got %s, wanted Foo", tag)
	}
}

func TestGetFkTableNamePlural(t *testing.T) {
	CopyTestTablesAndContextJson()
	value := GetFkTableNamePlural("film_actor", "actor_id")
	if value != "actors" {
		t.Errorf("Value was incorrect, got %s, wanted Foo", value)
	}
}

func TestGetEditSelectKeyForColumn(t *testing.T) {
	CopyTestTablesAndContextJson()
	value := GetEditSelectKeyForColumn("film_actor", "actorId")
	if value != "actorId" {
		t.Errorf("Value was incorrect, got %s, wanted actorId", value)
	}
}

func TestGetEditSelectValueForColumn(t *testing.T) {
	CopyTestTablesAndContextJson()
	value := GetEditSelectValueForColumn("film_actor", "actorId")
	if value != "firstName" {
		t.Errorf("Value was incorrect, got %s, wanted firstName", value)
	}
}

//getColumnCountByDataType
func TestGetColumnCountByDataType(t *testing.T) {
	value := getColumnCountByDataType("address", "varchar")
	if value != 5 {
		t.Errorf("Value was incorrect, got %d, wanted 5", value)
	}
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
