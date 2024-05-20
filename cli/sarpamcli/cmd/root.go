package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	var cfgFile string
	var userLicense string
	var projectBase string

	initConfig()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", "Sirajudheen Mohamed Ali", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Sirajudheen Mohamed Ali <sirajudheen.mohamed.ali@proton.me>")
	viper.SetDefault("license", "apache")

	fmt.Println("author:", viper.Get("author"))
	fmt.Println("ProjectBase:", viper.Get("projectbase"))
	fmt.Println("useViper:", viper.Get("useViper"))
	fmt.Println("license:", viper.Get("license"))

  }

func initConfig() {
	var cfgFile string
	cfgFile = "config.yaml"
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		fmt.Println(home)

		if err != nil {
		fmt.Println(err)
		os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "sarpam",
	Short: "sarpam is a demonstration of a Cobra CLI application",
	Long: `This is a long description of Sarpam. 
				It is a tool to demonstrate how to build a 
				CLI application using Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
	//   fmt.Println("First level of femida command")
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }