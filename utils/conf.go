package utils

import (
	"encoding/json"
	"os"

	"github.com/gabrielefagundes/commita/structs"
)

func CreateConfig() error {
	if _, err := os.Stat("./config.json"); err == nil {
		return nil
	}

	conf := structs.DefaultConfigs()

	data, _ := json.MarshalIndent(conf, "", " ")

	return os.WriteFile("./config.json", []byte(data), 0666)
}

func LoadConfig() (structs.Config, error) {
	var conf structs.Config

	data, err := os.ReadFile("./config.json")

	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(data, &conf)

	return conf, err
}

func SaveConfig(conf structs.Config) error {
	data, err := json.MarshalIndent(conf, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile("./config.json", data, 0666)
}
