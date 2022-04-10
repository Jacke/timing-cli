/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectsCmd represents the serve command
var timeEntriesCmd = &cobra.Command{
	Use:   "entries",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := getProjects()
		fmt.Println("projects called err:", err)
	},
}

var getEntryCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var newEntryCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var updateEntryCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var deleteEntryCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	timeEntriesCmd.AddCommand(getEntryCmd)
	timeEntriesCmd.AddCommand(newEntryCmd)
	timeEntriesCmd.AddCommand(updateEntryCmd)
	timeEntriesCmd.AddCommand(deleteEntryCmd)
	rootCmd.AddCommand(timeEntriesCmd)
}
