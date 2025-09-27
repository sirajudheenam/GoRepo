package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nameOfSoftware string
var shortName string
var longName string

func init() {
	nameOfSoftware = "sarpam"
	shortName = fmt.Sprintf("Print the version number of %s", nameOfSoftware)
	longName = fmt.Sprintf("All software has versions. This is %s's", nameOfSoftware)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: shortName,
	Long:  longName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nversionCmd.Run")
		fmt.Printf("%s v0.1 -- HEAD\n", nameOfSoftware)
	},
}
