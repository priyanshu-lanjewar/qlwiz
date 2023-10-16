/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os/exec"
	"strings"
	"syscall"

	ps "github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

// uiCmd represents the ui command
var uiCmd = &cobra.Command{
	Use: "ui",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("qlwiz", "launch-ui-cmd")

		command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		pl, err := ps.Processes()
		if err != nil {

		}
		count := 0
		for _, p := range pl {
			if strings.Contains(p.Executable(), "qlwiz") {
				count += 1
			}
		}

		if count <= 1 {
			_ = command.Start()
		} else {
			log.Fatalln("Already Runnin")
		}
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
