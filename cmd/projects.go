/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectsCmd represents the serve command
var projectsCmd = &cobra.Command{
	Use:   "projects",
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

var getProjectCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var newProjectCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var deleteProjectCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	projectsCmd.AddCommand(getProjectCmd)
	projectsCmd.AddCommand(newProjectCmd)
	projectsCmd.AddCommand(updateProjectCmd)
	projectsCmd.AddCommand(deleteProjectCmd)
	rootCmd.AddCommand(projectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
