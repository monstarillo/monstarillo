package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
	"github.com/go-easygen/easygen/egFilePath"
	"github.com/go-easygen/easygen/egVar"
	"github.com/monstarillo/monstarillo/models"
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

func ProcessTables(tables []models.Table, unitTestValuesJson, templateFile, gui string) {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egFilePath.FuncDefs()).
		Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs()).Funcs(funcMap)

	ts := models.ReadTemplates(templateFile)

	tablesToProcess := getTablesToProcess(tables, ts.IncludeTables, ts.IgnoreTables, gui)
	fmt.Println("Processing " + color.BlueString(strconv.Itoa(len(tablesToProcess))) + " tables")
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

	processTemplates(ts, tmpl, context, tablesToProcess)

}

func ProcessModels(modelsToProcess []models.OrmModel, templateFile string) {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egFilePath.FuncDefs()).
		Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs()).Funcs(funcMap)

	ts := models.ReadTemplates(templateFile)

	context := new(MonstarilloOrmContext)
	context.Models = modelsToProcess

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(dirname, ".monstarillo", "models.json")
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(fileName), 0777)
		check(err)
	}

	err = WriteModelsToJson(context.Models, filepath.Join(dirname, ".monstarillo", "tables.json"))
	if err != nil {
		log.Fatal(err)
		return
	}

	context.Tags = ts.Tags

	err = WriteOrmContextToJson(context)
	if err != nil {
		log.Fatal(err)
		return
	}

	processOrmTemplates(ts, tmpl, context, modelsToProcess)

}

func processOrmTemplates(ts models.Templates, tmpl *template.Template, context *MonstarilloOrmContext, modelsToProcess []models.OrmModel) {
	z := 0
	for range ts.Templates {
		var templatePathBuffer strings.Builder
		if err := easygen.Execute0(tmpl, &templatePathBuffer, ts.Templates[z].TemplateFile, context); err != nil {
			log.Fatal(err)
		}
		templatePath := templatePathBuffer.String()
		fmt.Println("Processing template : " + color.BlueString(templatePath))
		v := 0
		for range modelsToProcess {

			context.CurrentModel = modelsToProcess[v]

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
func processTemplates(ts models.Templates, tmpl *template.Template, ctx *MonstarilloContext, tablesToProcess []models.Table) {
	for _, t := range ts.Templates {
		// Resolve template path
		var tplPathBuf strings.Builder
		if err := easygen.Execute0(tmpl, &tplPathBuf, t.TemplateFile, ctx); err != nil {
			log.Printf("template path render failed: %v", err)
			continue
		}
		templatePath := tplPathBuf.String()
		fmt.Println("Processing template : " + color.BlueString(templatePath))

		for _, table := range tablesToProcess {
			// Set current table/gui
			ctx.CurrentTable = table
			ctx.CurrentGuiTable = models.GuiListTable{} // default clear
			for _, gui := range ctx.GuiListTables.Tables {
				if strings.EqualFold(gui.TableName, ctx.CurrentTable.TableName) {
					ctx.CurrentGuiTable = gui
					break
				}
			}

			// Resolve dynamic file name
			var fileBuf strings.Builder
			if err := easygen.Execute0(tmpl, &fileBuf, t.GeneratedFileName, ctx); err != nil {
				log.Printf("filename render failed (%s): %v", ctx.CurrentTable.TableName, err)
				continue
			}
			file := fileBuf.String()

			// Resolve output path
			var outPathBuf strings.Builder
			if err := easygen.Execute0(tmpl, &outPathBuf, t.OutputPath, ctx); err != nil {
				log.Printf("output path render failed (%s): %v", ctx.CurrentTable.TableName, err)
				continue
			}
			outputPath := outPathBuf.String()

			// Optional folder
			folder := ""
			if t.GeneratedFolderName != "" {
				var folderBuf strings.Builder
				if err := easygen.Execute0(tmpl, &folderBuf, t.GeneratedFolderName, ctx); err != nil {
					log.Printf("folder render failed (%s): %v", ctx.CurrentTable.TableName, err)
					continue
				}
				folder = folderBuf.String()
			}

			fileName := filepath.Join(outputPath, folder, file)

			// Ensure directory exists
			if err := os.MkdirAll(filepath.Dir(fileName), 0o755); err != nil {
				log.Printf("mkdir failed (%s): %v", fileName, err)
				continue
			}

			// Skip if we shouldn't overwrite and file exists
			if !t.OverwriteFile && fileExists(fileName) {
				// move on to next table, not next template
				continue
			}

			// Copy-only path
			if t.CopyOnly {
				if err := copyFile(templatePath, fileName); err != nil {
					log.Printf("copy failed %s -> %s: %v", templatePath, fileName, err)
				}
				continue
			}

			// Render once to memory
			var rendered strings.Builder
			if err := easygen.Execute(tmpl, &rendered, templatePath, ctx); err != nil {
				log.Printf("render failed (%s): %v", ctx.CurrentTable.TableName, err)
				continue
			}
			result := rendered.String()

			// Enforce minimum length (on the buffer we already have)
			if t.MinimumGeneratedFileLength > 0 && len(result) <= t.MinimumGeneratedFileLength {
				//fmt.Println(templatePath, ctx.CurrentTable.TableName, len(result))
				continue
			}

			// Write (truncate by default)
			flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
			f, err := os.OpenFile(fileName, flags, 0o644)
			if err != nil {
				log.Printf("open for write failed (%s): %v", fileName, err)
				continue
			}
			if _, err := f.WriteString(result); err != nil {
				log.Printf("write failed (%s): %v", fileName, err)
				_ = f.Close()
				continue
			}
			if err := f.Close(); err != nil {
				log.Printf("close failed (%s): %v", fileName, err)
				continue
			}
		}
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
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

func WriteOrmContextToJson(context *MonstarilloOrmContext) error {
	strContext, err := json.Marshal(context)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	err = WriteFile(strContext, filepath.Join(dirname, ".monstarillo", "orm-context.json"))

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

func WriteModelsToJson(models []models.OrmModel, fileName string) error {
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
