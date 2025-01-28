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

// postgresCmd represents the postgres command
var postgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Generate code against a Postgres database.",
	Long: `Postgres is a CLI library for Go that allows users to generate code for
database applications against a Postgres database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("postgres is in the building")
		unitTestValuesJson, _ := cmd.Flags().GetString("utv")

		templateFile, _ := cmd.Flags().GetString("t")
		database, _ := cmd.Flags().GetString("db")
		schema, _ := cmd.Flags().GetString("schema")
		gui, _ := cmd.Flags().GetString("gui")
		password, _ := cmd.Flags().GetString("p")
		userName, _ := cmd.Flags().GetString("u")
		port, _ := cmd.Flags().GetInt("port")
		host, _ := cmd.Flags().GetString("host")

		models.ConnectPostgresDB(userName, password, database, host, port)
		tables := models.GetPostgresTables(schema, database)
		engine.ProcessTables(tables, unitTestValuesJson, templateFile, gui)
		models.CloseDB()
	},
}

func init() {
	rootCmd.AddCommand(postgresCmd)

	postgresCmd.PersistentFlags().String("utv", "", "Unit Test Values file")
	postgresCmd.PersistentFlags().String("t", "", "Templates to run")
	postgresCmd.PersistentFlags().String("u", "", "DB user name")
	postgresCmd.PersistentFlags().String("p", "", "DB password")
	postgresCmd.PersistentFlags().String("db", "", "Database name")
	postgresCmd.PersistentFlags().String("schema", "", "Database schema name")
	postgresCmd.PersistentFlags().String("gui", "", "GUI Tables file")
	postgresCmd.PersistentFlags().Int("port", 5432, "Database port")
	postgresCmd.PersistentFlags().String("host", "postgres", "Database host")

	postgresCmd.MarkPersistentFlagRequired("t")

}
