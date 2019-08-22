package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	SlackToken string `toml:"slack_token"`
	Commands   map[string]command
}
type command struct {
	Keyword    string
	Command    string
	Aliases    []string
	Monospaced bool
}

func main() {
	var config tomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SLACK_TOKEN: %s\n", config.SlackToken)
	for _, v := range config.Commands {
		fmt.Printf("==========\n")
		fmt.Printf("KEYWORD: %s\n", v.Keyword)
		fmt.Printf("COMMAND: %s\n", v.Command)
		fmt.Printf("ALIASES: %v\n", v.Aliases)
		fmt.Printf("MONOSPACED: %v\n", v.Monospaced)
	}
	fmt.Printf("==========\n")
}
