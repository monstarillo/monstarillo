/*
Copyright © 2022 Patrick Wright

*/
package cmd

import (
	"fmt"
	"github.com/monstarillo/monstarillo/engine"
	"github.com/monstarillo/monstarillo/models"
	"github.com/spf13/cobra"
)

var oracleCmd = &cobra.Command{
	Use:   "oracle",
	Short: "Generate code against an Oracle database.",
	Long: `Oracle is a CLI library for Go that allows users to generate code for
database applications against a Oracle database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oracle is in the building")
		unitTestValuesJson, _ := cmd.Flags().GetString("utv")

		templateFile, _ := cmd.Flags().GetString("t")
		server, _ := cmd.Flags().GetString("server")
		schema, _ := cmd.Flags().GetString("schema")
		gui, _ := cmd.Flags().GetString("gui")
		password, _ := cmd.Flags().GetString("p")
		userName, _ := cmd.Flags().GetString("u")
		port, _ := cmd.Flags().GetInt("port")
		service, _ := cmd.Flags().GetString("service")

		models.ConnectOracleDB(userName, password, server, service, port)
		tables := models.GetOracleTables(schema)

		engine.ProcessTables(tables, unitTestValuesJson, templateFile, gui)
		models.CloseDB()
	},
}

func init() {
	rootCmd.AddCommand(oracleCmd)

	oracleCmd.PersistentFlags().String("utv", "", "Unit Test Values file")
	oracleCmd.PersistentFlags().String("t", "", "Templates to run")
	oracleCmd.PersistentFlags().String("u", "", "DB user name")
	oracleCmd.PersistentFlags().String("p", "", "DB password")
	oracleCmd.PersistentFlags().String("server", "", "Database server")
	oracleCmd.PersistentFlags().String("schema", "", "Database schema name")
	oracleCmd.PersistentFlags().String("service", "", "Database service name")
	oracleCmd.PersistentFlags().String("gui", "", "GUI Tables file")
	oracleCmd.PersistentFlags().Int("port", 1521, "Database port")

	oracleCmd.MarkPersistentFlagRequired("t")
}
