/*
Copyright © 2022 Patrick Wright

*/
package cmd

import (
	"fmt"
	"github.com/monstarillo/monstarillo/engine"
	"github.com/monstarillo/monstarillo/models"
	"github.com/spf13/cobra"
	"strconv"
)

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Generate code against a MySql database",
	Long: `Mysql is a CLI library for Go that allows users to generate code for
database applications against a MySql database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mysql is in the building")

		unitTestValuesJson, _ := cmd.Flags().GetString("utv")
		password, _ := cmd.Flags().GetString("p")
		userName, _ := cmd.Flags().GetString("u")
		templateFile, _ := cmd.Flags().GetString("t")
		//	database, _ := cmd.Flags().GetString("db")
		schema, _ := cmd.Flags().GetString("schema")
		gui, _ := cmd.Flags().GetString("gui")
		port, _ := cmd.Flags().GetInt("port")
		host, _ := cmd.Flags().GetString("host")

		database := "tcp(" + host + ":" + strconv.Itoa(port) + ")/" + schema
		models.ConnectDB(userName, password, database)
		tables := models.GetTables(schema)

		engine.ProcessTables(tables, unitTestValuesJson, templateFile, gui)

		models.CloseDB()
	},
}

func init() {
	rootCmd.AddCommand(mysqlCmd)

	mysqlCmd.PersistentFlags().String("utv", "", "Unit Test Values file")
	mysqlCmd.PersistentFlags().String("t", "", "Templates to run")
	mysqlCmd.PersistentFlags().String("u", "", "DB user name")
	mysqlCmd.PersistentFlags().String("p", "", "DB password")
	mysqlCmd.PersistentFlags().String("db", "", "database name")
	mysqlCmd.PersistentFlags().String("schema", "", "database schema name")
	mysqlCmd.PersistentFlags().String("gui", "", "GUI Tables file")
	mysqlCmd.PersistentFlags().Int("port", 3306, "Database port")
	mysqlCmd.PersistentFlags().String("host", "", "Database host")
	mysqlCmd.MarkPersistentFlagRequired("t")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//mysqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
