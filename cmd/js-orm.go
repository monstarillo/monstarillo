/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/monstarillo/monstarillo/engine"
	"github.com/monstarillo/monstarillo/models"
	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsorm = &cobra.Command{
	Use:   "js-orm",
	Short: "Execute templates against a Javascript ORM",
	Long:  `Execute templates against a Javascript ORM`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("js-orm is in the building")

		modelFile, _ := cmd.Flags().GetString("m")
		templateFile, _ := cmd.Flags().GetString("t")

		models := models.ReadModels(modelFile)

		a := 0
		for range models {
			fmt.Println(models[a].TableName + " " + models[a].ModelName)
			a++
		}

		engine.ProcessModels(models, templateFile)
	},
}

func init() {
	rootCmd.AddCommand(jsorm)
	//jsonCmd.PersistentFlags().String("t", "", "Templates to run")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	jsorm.PersistentFlags().String("m", "", "models.json file")
	jsorm.PersistentFlags().String("t", "", "Templates to run")

	err := jsorm.MarkPersistentFlagRequired("t")
	if err != nil {
		return
	}
	err = jsorm.MarkPersistentFlagRequired("m")
	if err != nil {
		return
	}
}
