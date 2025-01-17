/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/monstarillo/monstarillo/models"
	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var mssqlCmd = &cobra.Command{
	Use:   "mssql",
	Short: "Execute templates against a Microsoft SQL Server",
	Long:  `Execute templates against a Microsoft SQL Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mssql is in the building")

		//templateFile, _ := cmd.Flags().GetString("t")

		//engine.ProcessJson(templateFile)
		models.ConnectMsSqlServerDB("sa", "12345678abcABC1!", "northwind", "1433")
		tables := models.GetMssqlTables("dbo", "northwind")

		a := 0
		for range tables {
			fmt.Println(tables[a].TableName)
			a++
		}
	},
}

func init() {
	rootCmd.AddCommand(mssqlCmd)
	//jsonCmd.PersistentFlags().String("t", "", "Templates to run")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
