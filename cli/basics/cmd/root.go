package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type JsonConfig struct {
    Users    []string
    Groups   []string
}
type TomlConfig struct {
    Age int
    Cats []string
    Pi float64
    Perfection []int
    DOB time.Time
}

// type Config struct {
// 	Config []CLIConfig `yaml:"yamlconfig"`
// }

type CLIConfig struct {
	Name        string      `yaml:"name"`
	Path        string      `yaml:"path"`
	Description string      `yaml:"description"`
	Tags        []string    `yaml:"tags"`
	Parameters  []Parameter `yaml:"parameters"`
}

type Parameter struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
}

var rootCmd = &cobra.Command{
	Use:   "basics",
	Short: "basics is a demonstration of a Cobra CLI application",
	Long: `This is a long description of Basics. 
				It is a tool to demonstrate how to build a 
				CLI application using Cobra.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	  },
	PreRun: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside rootCmd Run with args: %v\n", args)
	fmt.Println("First level of basics command")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  log.Fatalf("error: %v", err)
	  os.Exit(1)
	}
  }