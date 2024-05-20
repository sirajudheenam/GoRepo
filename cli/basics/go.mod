module github.com/sirajudheenam/GoRepo/cli/basics

go 1.22.2

replace github.com/sirajudheenam/GoRepo/cli/basics/cmd => ./cmd

replace github.com/sirajudheenam/GoRepo/cli/basics/cmd/root => ./cmd/root

require github.com/sirajudheenam/GoRepo/cli/basics/cmd v0.0.0-00010101000000-000000000000

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
