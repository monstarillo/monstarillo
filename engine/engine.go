package engine

//
//import (
//	"fmt"
//	"os"
//	"strings"
//)
//
//func check(e error) {
//	if e != nil {
//		panic(e)
//	}
//}
//
//func ReadFile(file string) string {
//	data, err := os.ReadFile(file)
//	check(err)
//	return string(data)
//}
//
//func WriteFile(fileName, data string) {
//	d1 := []byte(data)
//	err := os.WriteFile(fileName, d1, 0644)
//	check(err)
//}
//
//func IndexAt(s, sep string, n int) int {
//	idx := strings.Index(s[n:], sep)
//	if idx > -1 {
//		idx += n
//	}
//	return idx
//}
//
//func ProcessTemplate(template, generatedFileName, generatedFolderName,
//	outputDirectory string, appendFile, overwriteFile bool, tables *[]Table) {
//
//	originalTemplateData := ReadFile(template)
//	keywordPosition := 0
//	position := 0
//	nextBeginLoop := 0
//	data := ""
//	keyword := ""
//
//	for {
//		beginKeyword := IndexAt(originalTemplateData, "%%ForEachTable", keywordPosition)
//
//		if beginKeyword > -1 {
//			endKeyword := IndexAt(originalTemplateData, "%%", beginKeyword+2)
//			keyword = originalTemplateData[beginKeyword+2 : endKeyword]
//			nextBeginLoop = IndexAt(originalTemplateData, "%%"+keyword+"%%", position)
//		} else {
//			keyword = ""
//		}
//
//		if nextBeginLoop == -1 {
//			break
//		}
//
//		if nextBeginLoop > -1 && len(keyword) > 1 {
//			lastReturnPosition := strings.LastIndex(originalTemplateData[0:nextBeginLoop], "\n") + 1
//			if lastReturnPosition > position {
//				data += originalTemplateData[position:lastReturnPosition]
//
//			} else {
//				data += originalTemplateData[position:nextBeginLoop]
//			}
//		}
//
//		nextEndLoop := IndexAt(originalTemplateData, "%%EndForEachTable", nextBeginLoop)
//		dataToLoop := originalTemplateData[nextBeginLoop+len(keyword)+6 : nextEndLoop]
//
//		position = nextEndLoop + len("%%EndForEachTable%%") + 1
//
//		processedData := ProcessTemplateData(dataToLoop, generatedFileName, generatedFolderName, outputDirectory, overwriteFile, false, 0, tables)
//
//		if strings.Contains(keyword, "RemoveLastComma") {
//			lastComma := strings.LastIndex(processedData, ",")
//			if lastComma > -1 {
//				processedData = processedData[0:lastComma] + processedData[lastComma+1:len(processedData)]
//			}
//		}
//
//		data += processedData
//
//	}
//
//	data += originalTemplateData[position:len(originalTemplateData)]
//	ProcessTemplateData(data, generatedFileName, generatedFolderName, outputDirectory, overwriteFile, true, 0, tables)
//
//}
//
//func ProcessTemplateData(templateData, generatedFileName, generatedFolderName,
//	outputDirectory string, overwriteFile, writeFile bool, minimumGeneratedFileLength int, tables *[]Table) string {
//
//	var dir = outputDirectory + "/" + generatedFolderName
//	err := os.MkdirAll(dir, 0755)
//	check(err)
//
//	//originalTemplateData := templateData
//	returnData := ""
//	var t = 0
//	var tbls []Table
//	tbls = *tables
//	for range *tables {
//		fmt.Println(tbls[t].TableName)
//		calculatedFileName := CalculatedFileName(generatedFileName, &tbls[t])
//		fileName := dir + "/" + calculatedFileName
//		if overwriteFile == false {
//			if FileExists(fileName) {
//				return ""
//			}
//		}
//
//		data := ""
//		position := 0
//		keywordPosition := 0
//		keyword := ""
//		nextBeginLoop := 0
//
//		for {
//			beginKeyword := IndexAt(templateData, "%%", keywordPosition)
//
//			if beginKeyword > -1 {
//				endKeyword := IndexAt(templateData, "%%", beginKeyword+2)
//				keyword = templateData[beginKeyword+2 : endKeyword]
//
//			} else {
//				keyword = ""
//			}
//
//			nextBeginLoop = IndexAt(templateData, "%%"+keyword+"%%", position)
//			if nextBeginLoop == -1 {
//				break
//			}
//
//			if contains(GetGlobalKeywords(), keyword) || contains(GetColumnKeywords(), keyword) {
//				data += templateData[position : nextBeginLoop+len(keyword)+4]
//				position = nextBeginLoop + len(keyword) + 4
//				keywordPosition = position
//			} else {
//				lastReturnPosition := strings.LastIndex(templateData[0:nextBeginLoop], "\n") + 1
//
//				if lastReturnPosition > position {
//					data += templateData[position:lastReturnPosition]
//
//				} else {
//					data += templateData[position:nextBeginLoop]
//				}
//
//				endLoopString := "%%EndFor%%"
//				if strings.Contains(keyword, "Column") {
//					endLoopString = "%%EndForEachColumn%%"
//				}
//
//				nextEndLoop := IndexAt(templateData, endLoopString, nextBeginLoop)
//
//				dataToLoop := templateData[nextBeginLoop+len(keyword)+6 : nextEndLoop]
//
//				data += ProcessColumnLoop(dataToLoop, keyword, &tbls[t])
//
//				position = nextEndLoop + len(endLoopString) + 1
//				keywordPosition = position
//
//			}
//		}
//
//		data += templateData[position:len(templateData)]
//		data = ProcessGlobals(data, tbls[t])
//		err = os.Remove(fileName)
//
//		if len(data) > minimumGeneratedFileLength && writeFile {
//			WriteFile(fileName, data)
//		}
//		returnData += data
//		t++
//	}
//
//	return returnData
//}
//
//func CalculatedFileName(fileName string, table *Table) string {
//	t := *table
//	fileName = strings.ReplaceAll(fileName, "%%TableName%%", t.TableName)
//	return fileName
//}
//
//func FileExists(fileName string) bool {
//	if _, err := os.Stat(fileName); err == nil {
//		return true
//	} else {
//		return false
//	}
//}
//
//func contains(s []string, str string) bool {
//	for _, v := range s {
//		if v == str {
//			return true
//		}
//	}
//
//	return false
//}
//
//func ProcessGlobals(data string, table Table) string {
//
//	data = strings.ReplaceAll(data, "%%TableName%%", table.TableName)
//	return data
//}
//
//func ProcessColumnLoop(data, keyword string, table *Table) string {
//	var columns []Column
//
//	switch keyword {
//	case "ForEachColumn", "ForEachColumnRemoveLastComma":
//		columns = table.Columns
//	}
//
//	originalData := data
//	returnData := ""
//	c := 0
//	position := 0
//	for range columns {
//		data := ""
//		position = 0
//		keywordPosition := 0
//		keyword := ""
//
//		for {
//			beginKeyword := IndexAt(originalData, "%%", keywordPosition)
//
//			if beginKeyword > -1 {
//				endKeyword := IndexAt(originalData, "%%", beginKeyword+2)
//				keyword = originalData[beginKeyword+2 : endKeyword]
//
//			} else {
//				keyword = ""
//			}
//
//			nextBeginLoop := IndexAt(originalData, "%%"+keyword+"%%", position)
//			if nextBeginLoop == -1 {
//				break
//			}
//
//			if contains(GetGlobalKeywords(), keyword) || contains(GetColumnKeywords(), keyword) {
//				data += originalData[position:nextBeginLoop]
//
//				///data += originalData.substr(nextBeginLoop, keyword.length + 4);
//				position = nextBeginLoop + len(keyword) + 4
//				keywordPosition = position
//			} else {
//
//				lastReturnPosition := strings.LastIndex(originalData[0:nextBeginLoop], "\n") + 1
//				//lastReturnPosition := IndexAt(templateData, "\n", nextBeginLoop )
//				if lastReturnPosition > position {
//					data += originalData[position:lastReturnPosition]
//
//				} else {
//					data += originalData[position:nextBeginLoop]
//				}
//				endLoopString := "%%EndFor%%"
//				if strings.Contains(keyword, "Column") {
//					endLoopString = "%%EndForEachColumn%%"
//				}
//
//				nextEndLoop := IndexAt(originalData, endLoopString, nextBeginLoop)
//
//				dataToLoop := originalData[nextBeginLoop+len(keyword)+6 : nextEndLoop]
//
//				data += dataToLoop
//
//				position = nextEndLoop + len(endLoopString) + 1
//				keywordPosition = position
//			}
//
//		}
//		data += originalData[position:len(originalData)]
//
//		data = strings.ReplaceAll(originalData, "%%ColumnName%%", columns[c].ColumnName)
//
//		returnData += data
//		c++
//	}
//	if strings.Contains(keyword, "RemoveLastComma") {
//		lastComma := strings.LastIndex(returnData, ",")
//		if lastComma > -1 {
//			returnData = returnData[0:lastComma] + returnData[lastComma+1:len(returnData)]
//		}
//	}
//
//	return returnData
//}
//
//func GetColumnKeywords() []string {
//	return []string{
//		"ColumnName",
//		"ColumnNamePascalCase",
//		"ColumnNameCamelCase",
//		"SetString",
//		"JavaSetStringJson",
//		"ColumnDataType",
//		"ColumnJavaDataType",
//		"ColumnCSharpDataType",
//		"ColumnFirstUnitTestValueJava",
//		"ColumnFirstUnitTestValueJavaAsString",
//		"ColumnSecondUnitTestValueJava",
//		"ColumnSecondUnitTestValueJavaAsString",
//		"ColumnIsRequired",
//		"ColumnHibernateGeneratedValue",
//		"ColumnHibernateRelation",
//		"ColumnHibernateRelationProperties",
//		"GetHibernateRelationSetFirstUnitTestValuesDto",
//		"GetHibernateRelationSetSecondUnitTestValuesDto",
//		"GetHibernateRelationSetFirstUnitTestValuesModel",
//		"GetHibernateRelationSetSecondUnitTestValuesModel",
//		"HibernateId",
//		"GetPrimaryHibernateRelationColumn",
//	}
//}
//
//func GetGlobalKeywords() []string {
//	return []string{
//		"CamelTableName",
//		"CamelTableNameEF",
//		"TableName",
//		"PascalTableName",
//		"PascalTableNameEF",
//		"PrimaryColumnJavaType",
//		"PrimaryColumnCSharpType",
//		"JavaImportsForDataTypes",
//		"GetPrimaryCamelCaseColumnNames",
//		"GetPrimaryCamelCaseColumnNamesIdForCompositeKey",
//		"GetPrimaryPascalCaseColumnNames",
//		"GetPrimaryPascalCaseColumnNamesDtoPrefix",
//		"GetPrimaryCSharpTypesItIsAnyPrefix",
//		"GetFirstPrimaryUnitTestValuesJava",
//		"GetFirstPrimaryUnitTestValuesAsStringsJava",
//		"GetSecondPrimaryUnitTestValuesJava",
//		"GetSecondPrimaryUnitTestValuesAsStringsJava",
//		"GetApiMappingString",
//		"GetApiPathVariableString",
//		"GetUrlTemplateVariableString",
//		"GetAllColumnsListWithCSharpTypes",
//		"GetPrimaryColumnsListWithCSharpTypes",
//		"GetHibernateRepositoryMethods",
//		"HibernateManyToManyRelationImports",
//		"HibernateManyToManyRelationPropertiesImports",
//	}
//}
