package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Basics",
  Long:  `All software has versions. This is Basic's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Basics v0.1 -- HEAD")
  },
}