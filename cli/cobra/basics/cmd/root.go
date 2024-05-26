package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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


func init() {
	var cfgFile string

	var Verbose bool
	var Source string

	initConfig()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/.basics.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "Sirajudheen Mohamed Ali <sirajudheen.mohamed.ali@proton.me>")
	viper.SetDefault("license", "apache")
	fmt.Println("author:", viper.Get("author"))
	fmt.Println("license:", viper.Get("license"))

  }

func initConfig() {
	var cfgFile string
	cfgFile = "basics.yaml"
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		fmt.Println("home:", home)

		if err != nil {
		fmt.Println(err)
		os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".basics")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	viper.ReadInConfig()
	
}

var rootCmd = &cobra.Command{
	Use:   "basics",
	Short: "basics is a demonstration of a Cobra CLI application",
	Long: `This is a long description of Basics. 
				It is a tool to demonstrate how to build a 
				CLI application using Cobra.`,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	//   },
	// PreRun: func(cmd *cobra.Command, args []string) {
	// fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	// },
	Run: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside rootCmd Run with args: %v\n", args)
	fmt.Println("First level of basics command")
	},
	// PostRun: func(cmd *cobra.Command, args []string) {
	// fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	// },
	// PersistentPostRun: func(cmd *cobra.Command, args []string) {
	// fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	// },
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  log.Fatalf("error: %v", err)
	  os.Exit(1)
	}
  }