package command

import (
	"fmt"
	"path"

	go_format "github.com/pefish/go-format"
	"github.com/pefish/hardhat-abigen/pkg/abigen"
	"github.com/pefish/hardhat-abigen/pkg/global"
	type_ "github.com/pefish/hardhat-abigen/pkg/type"

	"github.com/pefish/go-commander"
	go_file "github.com/pefish/go-file"
)

type PersistenceDataType struct {
}

type DefaultCommand struct {
	persistenceData *PersistenceDataType
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{
		persistenceData: &PersistenceDataType{},
	}
}

func (dc *DefaultCommand) Config() interface{} {
	return &global.GlobalConfig
}

func (dc *DefaultCommand) Data() interface{} {
	return dc.persistenceData
}

func (dc *DefaultCommand) Init(command *commander.Commander) error {
	return nil
}

func (dc *DefaultCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *DefaultCommand) Start(command *commander.Commander) error {
	artifactFile := command.Args["artifactFile"]

	var artifactData type_.ArtifactData
	err := go_file.ReadJsonFile(&artifactData, artifactFile)
	if err != nil {
		return err
	}

	abiDir := path.Join(global.GlobalConfig.OutDir, "abi")
	err = go_file.AssertPathExist(abiDir)
	if err != nil {
		return err
	}
	abiFilename := path.Join(abiDir, fmt.Sprintf("%s.abi", artifactData.ContractName))
	err = go_file.WriteJsonFile(abiFilename, artifactData.ABI)
	if err != nil {
		return err
	}

	bytecodeDir := path.Join(global.GlobalConfig.OutDir, "bytecode")
	err = go_file.AssertPathExist(bytecodeDir)
	if err != nil {
		return err
	}
	bytecodeFilename := path.Join(bytecodeDir, fmt.Sprintf("%s.bin", artifactData.ContractName))
	err = go_file.WriteFile(bytecodeFilename, []byte(artifactData.ByteCode[2:]))
	if err != nil {
		return err
	}

	bindDir := path.Join(global.GlobalConfig.OutDir, "bind")
	err = go_file.AssertPathExist(bindDir)
	if err != nil {
		return err
	}
	gofileFilename := path.Join(
		bindDir,
		fmt.Sprintf("%s.go", go_format.CamelCaseToUnderscore(artifactData.ContractName)),
	)
	err = abigen.Run(
		abiFilename,
		artifactData.ContractName,
		"bind",
		gofileFilename,
		bytecodeFilename,
	)
	if err != nil {
		return err
	}
	return nil
}
