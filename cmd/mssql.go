/*
Copyright © 2022 Patrick Wright
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/monstarillo/monstarillo/engine"
	"github.com/monstarillo/monstarillo/models"
	"github.com/spf13/cobra"
)

// mssqlCmd represents the mssql command
var mssqlCmd = &cobra.Command{
	Use:   "mssql",
	Short: "Generate code against a Microsoft SQL Server database.",
	Long: `Mssql is a CLI library for Go that allows users to generate code for
database applications against a Microsoft SQL Server database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mssql is in the building")
		unitTestValuesJson, _ := cmd.Flags().GetString("utv")

		templateFile, _ := cmd.Flags().GetString("t")
		database, _ := cmd.Flags().GetString("db")
		schema, _ := cmd.Flags().GetString("schema")
		gui, _ := cmd.Flags().GetString("gui")
		password, _ := cmd.Flags().GetString("p")
		userName, _ := cmd.Flags().GetString("u")
		port, _ := cmd.Flags().GetInt("port")
		host, _ := cmd.Flags().GetString("host")

		models.ConnectMsSqlServerDB(userName, password, database, host, strconv.Itoa(port))
		tables := models.GetMssqlTables(schema, database)
		fmt.Println("Found " + color.BlueString(strconv.Itoa(len(tables))) + " tables")
		engine.ProcessTables(tables, unitTestValuesJson, templateFile, gui)
		models.CloseDB()
	},
}

func init() {
	rootCmd.AddCommand(mssqlCmd)

	mssqlCmd.PersistentFlags().String("utv", "", "Unit Test Values file")
	mssqlCmd.PersistentFlags().String("t", "", "Templates to run")
	mssqlCmd.PersistentFlags().String("u", "", "DB user name")
	mssqlCmd.PersistentFlags().String("p", "", "DB password")
	mssqlCmd.PersistentFlags().String("db", "", "Database name")
	mssqlCmd.PersistentFlags().String("schema", "dbo", "Database schema name")
	mssqlCmd.PersistentFlags().String("gui", "", "GUI Tables file")
	mssqlCmd.PersistentFlags().Int("port", 1433, "Database port")
	mssqlCmd.PersistentFlags().String("host", "localhost", "Database host")

	mssqlCmd.MarkPersistentFlagRequired("t")
}
