/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get is Palette that contains getting resources based commands",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
