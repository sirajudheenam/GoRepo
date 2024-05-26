package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(cwdCmd)
}

var cwdCmd = &cobra.Command{
  Use:   "cwd",
  Short: "Prints the current working directory",
  Long:  `This prints the current working directory.`,
  // PreRun: func(cmd *cobra.Command, args []string) {
	// fmt.Printf("Inside cwdCmd PreRun with args: %v\n", args)
  // },
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside cwdCmd Run with args: %v\n", args)
	fileStats()
  },
  // PostRun: func(cmd *cobra.Command, args []string) {
	// fmt.Printf("Inside cwdCmd PostRun with args: %v\n", args)
  // },
  // PersistentPostRun: func(cmd *cobra.Command, args []string) {
	// fmt.Printf("Inside cwdCmd PersistentPostRun with args: %v\n", args)
  // },

}

func fileStats() {
  fmt.Println("cwdStats: -")
  arguments := os.Args
  pwd, err := os.Getwd()
	fmt.Println("pwd: -", pwd)
  if err != nil {
      fmt.Println("Error:", err)
  }
  if len(arguments) == 1 {
      return
  }
    // if arguments[1] != "-P" {
    //     return
    // }
  fileinfo, err := os.Lstat(pwd)
  fmt.Println("fileinfo: ", fileinfo)
	fmt.Printf("File name: %s\n", fileinfo.Name())
	if err != nil {
		log.Println(err)
	}
  if fileinfo.Mode()&os.ModeSymlink != 0 {
      realpath, err := filepath.EvalSymlinks(pwd)
      if err != nil {
          log.Println(err)
      }
      fmt.Println(realpath)
  }
}