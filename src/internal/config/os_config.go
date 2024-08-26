package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

const (
	WINDOWS = iota
	LINUX   = iota
	MACOS   = iota
)

type GoGenConfig struct {
	TaregtOS string `json:"targetOS"`
}

const gogenFileName = ".gogen_config"

func LoadPackageConfig() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error getting home user directory: %v \n", err)
	}

	configPath := filepath.Join(homeDir, gogenFileName)
	gogenConfig := GoGenConfig{}

	data, err := os.ReadFile(configPath)
	if err == nil {
		err = json.Unmarshal(data, &gogenConfig)
		if err == nil && gogenConfig.TaregtOS != "" {
			return gogenConfig.TaregtOS
		}
	}

	identifiedOS := runtime.GOOS

	switch identifiedOS {
	case "windows":
		gogenConfig.TaregtOS = "windows"
	case "linux":
		gogenConfig.TaregtOS = "linux"
	case "darwin":
		gogenConfig.TaregtOS = "macos"
	default:
		gogenConfig.TaregtOS = "windows"
	}

	// Save the config
	data, err = json.Marshal(gogenConfig)
	if err != nil {
		color.Red("Error saving config: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		color.Red("Error writing config file: %v\n", err)
		os.Exit(1)
	}

	return gogenConfig.TaregtOS
}
