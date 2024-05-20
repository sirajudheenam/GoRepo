package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	toml "github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
  rootCmd.AddCommand(readConfigCmd)
}

var readConfigCmd = &cobra.Command{
  Use:   "readconfig [configfile]",
  Short: "Reads the configuration file and prints it out",
  Long:  `This reads the various types of configuration files. Currently supports JSON, TOML, YAML.`,
  PreRun: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside readConfigCmd PreRun with args: %v\n", args)
  },
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside readConfigCmd Run with args: %v\n", args)
	readJSON()
	readTOML()
	readYAML()
  },
  PostRun: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside readConfigCmd PostRun with args: %v\n", args)
  },
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
	fmt.Printf("Inside readConfigCmd PersistentPostRun with args: %v\n", args)
  },

}

func readJSON() {
	file, err := os.Open("configFiles/config1.json")
	if err != nil {
	  fmt.Println("error opening file:", err)
	}
	
	defer file.Close()
	jsonDecoder := json.NewDecoder(file)
	jsonConfig := JsonConfig{}
	err = jsonDecoder.Decode(&jsonConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(jsonConfig.Users)
	fmt.Println(jsonConfig.Groups)
}
func readTOML() {
	var tomlConf TomlConfig

	if _, err := toml.DecodeFile("configFiles/config2.toml", &tomlConf); err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(tomlConf.Age) 
	fmt.Println(tomlConf.Cats)

}
func readYAML() {
	data, err := os.ReadFile("configFiles/config3.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	var config CLIConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("\nParsed YAML: \n %+v\n", config)

	params := config.Parameters

	for k, v := range params {
		fmt.Printf("\n %v -> %v\n", k, v)
	}
}