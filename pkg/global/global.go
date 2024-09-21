package global

import "github.com/pefish/go-commander"

type Config struct {
	commander.BasicConfig
	OutDir string `json:"out-dir" default:"./contract" usage:"Destination directory for extracted ABI json files and generated bindings."`
}

var GlobalConfig Config
