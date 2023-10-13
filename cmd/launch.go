/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/links"
	"github.com/spf13/cobra"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:     "launch",
	Short:   "launch a link/group",
	Aliases: []string{"ll"},
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flag("url").Value.String()
		if url == "" {
			link, _ := links.GetAllLinks()
			linkToLaunch, _ := links.GetLinksToLaunch(link)
			links.Launch(linkToLaunch)
		} else {
			links.Launch(helpers.Link{
				Link: url,
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
	launchCmd.Flags().String("url", "", "")
}
