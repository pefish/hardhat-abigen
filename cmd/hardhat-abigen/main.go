package main

import (
	"fmt"
	"log"

	"github.com/pefish/hardhat-abigen/cmd/hardhat-abigen/command"
	"github.com/pefish/hardhat-abigen/version"

	"github.com/pefish/go-commander"
)

func main() {
	commanderInstance := commander.NewCommander(
		version.AppName,
		version.Version,
		fmt.Sprintf("%s is a tool to generate go files by hardhat artifacts file. Author: pefish", version.AppName),
	)
	commanderInstance.RegisterDefaultSubcommand(&commander.SubcommandInfo{
		Desc:       "Use this command by default if you don't set subcommand.",
		Args:       []string{"artifactFile"},
		Subcommand: command.NewDefaultCommand(),
	})
	err := commanderInstance.Run()
	if err != nil {
		log.Fatal(err)
	}
}
