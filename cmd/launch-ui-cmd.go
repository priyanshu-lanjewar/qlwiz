/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/getlantern/systray"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/ui"
	"github.com/spf13/cobra"
)

// uiCmd represents the ui command
var lui = &cobra.Command{
	Use: "launch-ui-cmd",

	Run: func(cmd *cobra.Command, args []string) {

		onExit := func() {
		}
		systray.Run(ui.OnReady, onExit)
	},
}

func init() {
	rootCmd.AddCommand(lui)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
