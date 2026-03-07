package setup

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gabrielefagundes/commita/structs"
)

var (
	dir, _           = os.UserHomeDir()
	commitaConfigDir = fmt.Sprintf("%s\\commita\\config.json", dir)
)

func CreateConfig() error {
	if _, err := os.Stat(commitaConfigDir); err == nil {
		return nil
	}

	conf := structs.DefaultConfigs()

	data, _ := json.MarshalIndent(conf, "", " ")

	return os.WriteFile(commitaConfigDir, []byte(data), 0666)
}

func LoadConfig() (structs.Config, error) {
	var conf structs.Config

	data, err := os.ReadFile(commitaConfigDir)

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

	return os.WriteFile(commitaConfigDir, data, 0666)
}
