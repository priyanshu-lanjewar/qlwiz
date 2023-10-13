/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/qlwiz/pkg/manager"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "To edit config manually",

	Run: func(cmd *cobra.Command, args []string) {
		manager.EditConfig()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
