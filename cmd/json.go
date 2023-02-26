/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/monstarillo/monstarillo/engine"
	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Execute templates against a json file",
	Long:  `Execute templates against a json file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("json is in the building")

		templateFile, _ := cmd.Flags().GetString("t")

		engine.ProcessJson(templateFile)
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.PersistentFlags().String("t", "", "Templates to run")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
