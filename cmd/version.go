/*
Copyright © 2026 Syrup Studios>
*/
package cmd

import (
	"fmt"
	"syrup-studios/suffering-server/utils"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Outputs the version of the server",
	Long:  `The version command outputs the version of the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
