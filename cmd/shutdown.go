/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// shutdownCmd represents the shutdown command
var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shutdown the SUFFERING Server",
	Long:  `Shutdown the SUFFERING Server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shutdown called")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(shutdownCmd)
}
