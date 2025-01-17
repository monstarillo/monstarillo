package engine

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
	"github.com/go-easygen/easygen/egFilePath"
	"github.com/go-easygen/easygen/egVar"
	"github.com/iancoleman/strcase"
	"github.com/monstarillo/monstarillo/models"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ProcessJson(templateFile string) {
	const jsondata = `{"something":{"a":"valueofa"}, "somethingElse": [1234, 5678]}`

	jsonObject := map[string]interface{}{}
	if err := json.Unmarshal([]byte(jsondata), &jsonObject); err != nil {
		panic(err)
	}

	ts := models.ReadTemplates(templateFile)
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egFilePath.FuncDefs()).
		Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs()).Funcs(funcMap)
	z := 0
	for range ts.Templates {
		var templatePathBuffer strings.Builder
		if err := easygen.Execute0(tmpl, &templatePathBuffer, ts.Templates[z].TemplateFile, jsonObject); err != nil {
			log.Fatal(err)
		}
		templatePath := templatePathBuffer.String()

		fileName := filepath.Join(ts.Templates[z].OutputPath, ts.Templates[z].GeneratedFolderName, ts.Templates[z].GeneratedFileName)

		if _, err := os.Stat(filepath.Dir(fileName)); os.IsNotExist(err) {
			err := os.MkdirAll(filepath.Dir(fileName), 0777)
			check(err)
		}
		if ts.Templates[z].OverwriteFile {
			deleteFile(fileName)
		}

		if ts.Templates[z].OverwriteFile == false && fileExists(fileName) {
			break
		}

		if ts.Templates[z].CopyOnly == true {
			copyFile(templatePath, fileName)
			break
		}

		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}

		err = easygen.Execute(tmpl, f, templatePath, jsonObject)
		if err != nil {
			return
		}

		if ts.Templates[z].MinimumGeneratedFileLength > 0 {
			fi, err := f.Stat()
			if err != nil {
				// Could not obtain stat, handle error
			}
			if fi.Size() < int64(ts.Templates[z].MinimumGeneratedFileLength) {
				deleteFile(fileName)
			}
		}
		err = f.Close()
		if err != nil {
			return
		}
		fmt.Println("Processing template : " + color.BlueString(templatePath))
		z++
	}
}

func ProcessTables(tables []models.Table, unitTestValuesJson, templateFile, gui, caseModel, caseProperty, modelsFile string) {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egFilePath.FuncDefs()).
		Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs()).Funcs(funcMap)

	ts := models.ReadTemplates(templateFile)

	tablesToProcess := getTablesToProcess(tables, ts.IncludeTables, ts.IgnoreTables, gui)

	dbModels := models.ReadModels(modelsFile)
	tablesToProcess = ProcessModelData(tables, caseModel, caseProperty, dbModels)

	context := new(MonstarilloContext)
	context.Tables = tablesToProcess

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(dirname, ".monstarillo", "tables.json")
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(fileName), 0777)
		check(err) // path/to/whatever exists
	}

	err = WriteTablesToJson(context.Tables, filepath.Join(dirname, ".monstarillo", "tables.json"))
	if err != nil {
		log.Fatal(err)
		return
	}

	context.Tags = ts.Tags
	context.UnitTestValuesFile = unitTestValuesJson
	if len(gui) > 0 {
		context.GuiListTables = models.ReadGuiTables(gui)
	}

	err = WriteContextToJson(context)
	if err != nil {
		log.Fatal(err)
		return
	}

	z := 0
	for range ts.Templates {
		var templatePathBuffer strings.Builder
		if err := easygen.Execute0(tmpl, &templatePathBuffer, ts.Templates[z].TemplateFile, context); err != nil {
			log.Fatal(err)
		}
		templatePath := templatePathBuffer.String()
		fmt.Println("Processing template : " + color.BlueString(templatePath))
		v := 0
		for range tablesToProcess {

			context.CurrentTable = tablesToProcess[v]

			for _, gui := range context.GuiListTables.Tables {
				if strings.ToLower(gui.TableName) == strings.ToLower(context.CurrentTable.TableName) {
					context.CurrentGuiTable = gui
					break
				} else {
					context.CurrentGuiTable = models.GuiListTable{}
				}
			}
			var fileBuffer strings.Builder
			if err := easygen.Execute0(tmpl, &fileBuffer, ts.Templates[z].GeneratedFileName, context); err != nil {
				log.Fatal(err)
			}
			file := fileBuffer.String()

			var outputPathBuffer strings.Builder
			if err := easygen.Execute0(tmpl, &outputPathBuffer, ts.Templates[z].OutputPath, context); err != nil {
				log.Fatal(err)
			}
			outputPath := outputPathBuffer.String()

			var folderBuffer strings.Builder
			folder := ""
			if len(ts.Templates[z].GeneratedFolderName) > 0 {
				if err := easygen.Execute0(tmpl, &folderBuffer, ts.Templates[z].GeneratedFolderName, context); err != nil {
					log.Fatal(err)
				}
				folder = folderBuffer.String()
			}

			fileName := filepath.Join(outputPath, folder, file)

			if _, err := os.Stat(filepath.Dir(fileName)); os.IsNotExist(err) {
				err := os.MkdirAll(filepath.Dir(fileName), 0777)
				check(err)
			}

			if ts.Templates[z].OverwriteFile {
				deleteFile(fileName)
			}

			if ts.Templates[z].OverwriteFile == false && fileExists(fileName) {
				break
			}

			if ts.Templates[z].CopyOnly == true {
				copyFile(templatePath, fileName)
				break
			}

			f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
			if err != nil {
				log.Fatal(err)
			}

			err = easygen.Execute(tmpl, f, templatePath, context)

			if err != nil {
				return
			}
			var fileShouldBeDeleted = false
			if ts.Templates[z].MinimumGeneratedFileLength > 0 {
				fi, err := f.Stat()
				if err != nil {
					// Could not obtain stat, handle error
				}

				if fi.Size() < int64(ts.Templates[z].MinimumGeneratedFileLength) {

					fileShouldBeDeleted = true
				}
			}

			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
			if fileShouldBeDeleted {
				deleteFile(fileName)
			}

			v++
		}

		z++
	}

}

func ProcessModelData(tables []models.Table, caseModel, caseProperty string, userModels []models.Model) []models.Table {

	var dbModels []models.Model
	var useUserModels = len(userModels) > 0

	v := 0
	for range tables {
		if useUserModels {
			tables[v].ModelName = models.GetModelNameForTable(userModels, tables[v].TableName)
		} else {
			tables[v].ModelName = getCaseValue(caseModel, tables[v].TableName)
		}
		var model models.Model

		model.TableName = tables[v].TableName
		model.ModelName = tables[v].ModelName

		col := 0
		for range tables[v].Columns {
			if useUserModels {
				tables[v].Columns[col].PropertyName = models.GetPropertyNameForModelColumn(userModels, tables[v].TableName, tables[v].Columns[col].ColumnName)
			} else {
				tables[v].Columns[col].PropertyName = getCaseValue(caseProperty, tables[v].Columns[col].ColumnName)
			}

			var modelColumn models.ModelColumn
			modelColumn.ColumnName = tables[v].Columns[col].ColumnName
			modelColumn.PropertyName = tables[v].Columns[col].PropertyName
			model.ModelColumns = append(model.ModelColumns, modelColumn)
			col++
		}
		fmt.Println(tables[v].TableName + " " + tables[v].ModelName)
		dbModels = append(dbModels, model)
		v++
	}

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	err = WriteModelsToJson(dbModels, filepath.Join(dirname, ".monstarillo", "models.json"))
	if err != nil {
		return nil
	}
	return tables
}

func getCaseValue(caseToReturn, value string) string {
	switch caseToReturn {
	case "pascal":
		return strcase.ToCamel(value)

	case "camel":
		return strcase.ToLowerCamel(value)
	case "kebab":
		return strcase.ToKebab(value)

	case "snake":
		return strcase.ToSnake(value)

	}
	return strcase.ToCamel(value)
}

func copyFile(src, dst string) {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(dst), 0777)
		check(err) // path/to/whatever exists
	}
	fin, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
}

func WriteContextToJson(context *MonstarilloContext) error {
	strContext, err := json.Marshal(context)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	err = WriteFile(strContext, filepath.Join(dirname, ".monstarillo", "context.json"))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func WriteFile(fileData []byte, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = f.Write(fileData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func WriteTablesToJson(tables []models.Table, fileName string) error {
	strTables, err := json.Marshal(tables)
	if err != nil {
		return err
	}

	err = WriteFile(strTables, fileName)
	if err != nil {
		return err
	}
	return nil
}

func WriteModelsToJson(models []models.Model, fileName string) error {
	modelData, err := json.Marshal(models)
	if err != nil {
		return err
	}

	err = WriteFile(modelData, fileName)
	if err != nil {
		return err
	}
	return nil
}

func getTablesToProcess(tables []models.Table, includeTables, ignoreTables []string, guiFile string) []models.Table {
	var tablesToProcess []models.Table
	if len(includeTables) > 0 {
		for _, iTbl := range includeTables {
			for _, tbl := range tables {
				if tbl.TableName == iTbl {
					tablesToProcess = append(tablesToProcess, tbl)
				}
			}
		}
	} else if len(ignoreTables) > 0 {

		for _, tbl := range tables {
			found := false
			for _, iTbl := range ignoreTables {
				if tbl.TableName == iTbl {
					found = true
					break
				}
			}
			if !found {
				tablesToProcess = append(tablesToProcess, tbl)
			}

		}
	} else {
		tablesToProcess = tables
	}

	if len(guiFile) > 0 {
		guiTables := models.ReadGuiTables(guiFile)

		t := 0
		for range tablesToProcess {
			for _, g := range guiTables.Tables {
				if g.TableName == tablesToProcess[t].TableName {
					tablesToProcess[t].GuiListTable = g
					fmt.Println(g.TableName)
				}
			}
			t++
		}
	}

	return tablesToProcess
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
