/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"github.com/monstarillo/monstarillo/engine"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Create a templates.json file from an existing directory",
	Long:  `Create a templates.json file from an existing directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("importing is fun!!")
		ignoreFiles, _ := cmd.Flags().GetString("iFiles")
		ignoreDirectories, _ := cmd.Flags().GetString("iFolders")
		templates := getTemplates("/Users/patrickwright/code-gen-output/vue-g01", ignoreFiles, ignoreDirectories)
		templateRoot, _ := cmd.Flags().GetString("tFolder") //"/Users/patrickwright/code-gen-output/templatesX"
		if _, err := os.Stat(templateRoot); os.IsNotExist(err) {
			err := os.MkdirAll(templateRoot, 0755)
			check(err)
		}
		fmt.Println("=========")
		var data bytes.Buffer
		data.WriteString("{\n")
		data.WriteString("\t\"templates\": [\n")
		count := 0
		for _, t := range templates {
			fmt.Println(getTemplateJson(t))
			if count > 0 {
				data.WriteString(",")
			}
			data.WriteString("\n")
			data.WriteString(getTemplateJson(t))
			templateFile := filepath.Join(templateRoot, t.GeneratedFolderName, t.GeneratedFileName)

			fmt.Println(templateFile)
			copyFile(t.RawTemplateFile, templateFile)
			count++
		}
		data.WriteString("\n")
		data.WriteString("\t],\n")

		data.WriteString("\t\"tags\": [\n    {\n      \"tagName\": \"TemplateRoot\",\n      \"value\": \"" + templateRoot + "\"\n    },\n    {\n      \"tagName\": \"OutputPath\",\n      \"value\": \"\"\n    }\n ]\n")
		data.WriteString("}")
		engine.WriteFile(data.Bytes(), filepath.Join(templateRoot, "template.json"))
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.PersistentFlags().String("sFolder", "", "Source folder")
	importCmd.PersistentFlags().String("tFolder", "", "Template folder")
	importCmd.PersistentFlags().String("iFiles", "", "Files to ignore")
	importCmd.PersistentFlags().String("iFolders", "", "Folders to ignore")

	err := importCmd.MarkPersistentFlagRequired("tFolder")
	if err != nil {
		return
	}
	err = importCmd.MarkPersistentFlagRequired("sFolder")
	if err != nil {
		return
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func shouldIgnore(path, ignoreFiles, ignoreDirectories string) bool {

	files := strings.Split(ignoreFiles, ",")
	for _, f := range files {
		if filepath.Base(path) == f {
			return true
		}
	}
	dirs := strings.Split(ignoreDirectories, ",")
	for _, d := range dirs {
		if len(d) > 0 && strings.Contains(path, d) {
			fmt.Println("Ignoring " + d)
			return true
		}
	}
	return false
}

func getTemplates(directory, ignoreFiles, ignoreDirectories string) []TemplateJson {
	var templates []TemplateJson
	fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !shouldIgnore(path, ignoreFiles, ignoreDirectories) && !info.IsDir() {
			var template TemplateJson
			template.RawTemplateFile = path
			relPath := strings.Replace(path, directory, "", 1)
			dir, file := filepath.Split(relPath)

			template.GeneratedFolderName = dir
			template.GeneratedFileName = file
			template.OverwriteFile = true
			template.MinimumGeneratedFileLength = 0
			template.TemplateFile = "{{getTag .Tags \\\"TemplateRoot\\\"}}" + relPath
			template.OutputPath = "{{getTag .Tags \\\"OutputPath\\\"}}"
			fmt.Printf("dir: %v: name: %s: rel: %s dir: %s file: %s\n", info.IsDir(), path, relPath, dir, file)
			templates = append(templates, template)

		}
		return nil
	}
	err := filepath.Walk(directory, fn)
	if err != nil {
		fmt.Println(err)
	}
	return templates
}

func getTemplateJson(template TemplateJson) string {
	var data bytes.Buffer

	data.WriteString("\t{\n")
	data.WriteString("\t\t\"templateFile\": \"" + template.TemplateFile + "\",\n")
	data.WriteString("\t\t\"generatedFileName\": \"" + template.GeneratedFileName + "\",\n")
	data.WriteString("\t\t\"generatedFolderName\": \"" + template.GeneratedFolderName + "\",\n")
	data.WriteString("\t\t\"minimumGeneratedFileLength\": 0,\n")
	data.WriteString("\t\t\"outputPath\": \"" + template.OutputPath + "\",\n")
	data.WriteString("\t\t\"overwriteFile\": true\n,")
	data.WriteString("\t\t\"copyOnly\": true\n")
	data.WriteString("\t}")

	return data.String()
}

type TemplateJson struct {
	TemplateFile               string `json:"templateFile"`
	GeneratedFileName          string `json:"generatedFileName"`
	GeneratedFolderName        string `json:"generatedFolderName"`
	OutputPath                 string `json:"outputPath"`
	OverwriteFile              bool   `json:"overwriteFile"`
	MinimumGeneratedFileLength int    `json:"minimumGeneratedFileLength"`
	RawTemplateFile            string
}

func copyFile(src, dst string) {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(dst), 0755)
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

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
