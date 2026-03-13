package setup

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gabrielefagundes/commita/structs"
)

var (
	// //dir, _ = os.UserHomeDir()
	// //commitaGlobalConfigDir = fmt.Sprintf("%s\\commita\\config.json", dir)
	// TODO: implement global configs aswell, merge them with the local configs.
	commitaLocalConfigDir, dirErr = getLocalDir()
)

func CreateConfig() error {
	if dirErr != nil {
		return dirErr
	}

	if _, err := os.Stat(commitaLocalConfigDir); err == nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(commitaLocalConfigDir), os.ModePerm); err != nil {
		return err
	}

	conf := structs.DefaultConfigs()
	data, _ := json.MarshalIndent(conf, "", "\t")

	return os.WriteFile(commitaLocalConfigDir, []byte(data), 0666)
}

func LoadConfig() (*structs.Config, error) {
	var conf structs.Config

	if dirErr != nil {
		return nil, dirErr
	}

	data, err := os.ReadFile(commitaLocalConfigDir)

	if err != nil {
		return &conf, err
	}

	err = json.Unmarshal(data, &conf)

	return &conf, err
}

func SaveConfig(conf structs.Config) error {
	if dirErr != nil {
		return dirErr
	}

	data, err := json.MarshalIndent(conf, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(commitaLocalConfigDir, data, 0666)
}

func getLocalDir() (string, error) {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("\nnot a git repository")
	}

	root := strings.TrimSpace(string(output))
	return filepath.Join(root, ".commita", "config.json"), nil
}
