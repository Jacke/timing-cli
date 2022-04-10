/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"strings"
	// "math"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/hako/durafmt"
)

func FormatFloat(num float64, prc int) string {
	var (
			zero, dot = "0", "."

			str = fmt.Sprintf("%."+strconv.Itoa(prc)+"f", num)
	)

	return strings.TrimRight(strings.TrimRight(str, zero), dot)
}

var timersCmd = &cobra.Command{
	Use:   "timers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := GetTimersArguments{}
		response, err := getTimers(arguments)

		var data = [][]string{}
		for _, s := range response.Data {
			var isActive string = "-"
			if (s.IsRunning != nil && *s.IsRunning) {
				isActive = "+"
			}
			toAdd := []string{s.Self, s.Title, s.StartDate.Format(time.RFC1123), s.EndDate.Format(time.RFC1123), durafmt.Parse(time.Duration(s.Duration) * time.Second).String(), isActive}
			data = append(data, toAdd)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Start", "End", "Duration (min)", "Active"})
		for _, v := range data {
			 table.Append(v)
		 }

		if (err == nil) {
			table.Render()
		} else {
			fmt.Println("create called", err, response)
		}
	},
}

var startTimerCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var stopTimerCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var statusTimerCmd = &cobra.Command{
	Use:   "status",
	Short: "Get if any timers is running",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		var arguments GetTimersArguments
		isRunning := true
		arguments.IsRunning = &isRunning

		resp, err := getTimers(arguments)
		if (err != nil) {
			os.Exit(1)
		} else { 
			// for _, v := range resp.Data {
				// if v.IsRunning {
					// fmt.Println("timer is running", v)
					// os.Exit(0)
				// }
			// }
			if (resp.Meta.Total > 0) {
				fmt.Println("timer is running: ", strconv.Itoa(len(resp.Data)))
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
	},
}

func init() {
	timersCmd.AddCommand(startTimerCmd)
	timersCmd.AddCommand(stopTimerCmd)
	timersCmd.AddCommand(statusTimerCmd)
	rootCmd.AddCommand(timersCmd)
}
