/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/qlwiz/pkg/manager"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset Config File",

	Run: func(cmd *cobra.Command, args []string) {
		manager.Reset()
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
