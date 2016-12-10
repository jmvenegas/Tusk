package model

import (
	"encoding/json"
	"os"
)

type TuskConfig struct {
	Socket   string
	DBUser   string
	DBPass   string
	DBSocket string
	DBName   string
}

func NewTuskConfig(configPath string) *TuskConfig {
	file, err := os.Open(configPath)
	HandleError(err)
	decoder := json.NewDecoder(file)
	config := new(TuskConfig)
	err = decoder.Decode(config)
	HandleError(err)
	return config
}
