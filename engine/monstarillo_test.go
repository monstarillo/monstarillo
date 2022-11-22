package engine

import (
	"github.com/monstarillo/monstarillo/models"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestGetTablesToProcess_AllTables(t *testing.T) {
	tables := models.ReadTables("../test-data/tables.json")
	var includeTables, ignoreTables []string
	value := getTablesToProcess(tables, includeTables, ignoreTables, "")

	if len(value) != 12 {
		t.Errorf("Value was incorrect, got %d, wanted 12", len(value))
	}
}

func TestGetTablesToProcess_Include(t *testing.T) {
	tables := models.ReadTables("../test-data/tables.json")
	var includeTables, ignoreTables []string
	includeTables = append(includeTables, "fool_date")
	value := getTablesToProcess(tables, includeTables, ignoreTables, "")

	if len(value) != 1 {
		t.Errorf("Value was incorrect, got %d, wanted 1", len(value))
	}
}

func TestGetTablesToProcess_Ignore(t *testing.T) {
	tables := models.ReadTables("../test-data/tables.json")
	var includeTables, ignoreTables []string
	ignoreTables = append(ignoreTables, "fool_date")
	value := getTablesToProcess(tables, includeTables, ignoreTables, "")

	if len(value) != 11 {
		t.Errorf("Value was incorrect, got %d, wanted 11", len(value))
	}
}

func TestProcessTables(t *testing.T) {
	CopyTestTablesAndContextJson()
	tables := models.ReadTables("../test-data/tables.json")
	deleteFile("../UnitTestResults/test.txt")
	ProcessTables(tables, "", "../test-data/test.json", "")
	content, err := ioutil.ReadFile("../UnitTestResults/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	if !strings.Contains(text, "Language") {
		t.Errorf("Value was incorrect, did not find language, got %s", text)
	}

}
