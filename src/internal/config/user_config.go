package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/fedi-nabli/GoGen/src/internal/languages"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type UserConfig struct {
	PackageManager string `json:"packageManager"`
}

const configFileName = ".gogen_user_preferences"

func LoadUserConfig() UserConfig {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error getting user home directory: %v\n", err)
	}

	configPath := filepath.Join(homeDir, configFileName)
	userConfig := UserConfig{}

	data, err := os.ReadFile(configPath)
	if err == nil {
		err = json.Unmarshal(data, &userConfig)
		if err == nil && userConfig.PackageManager != "" {
			return userConfig
		}
	}

	packageManager := choosePackageManager()

	switch packageManager {
	case "npm":
		userConfig.PackageManager = "npm"
	case "yarn":
		userConfig.PackageManager = "yarn"
	case "pnpm":
		userConfig.PackageManager = "pnpm"
	default:
		userConfig.PackageManager = "npm"
	}

	// Save the config
	data, err = json.Marshal(userConfig)
	if err != nil {
		color.Red("Error saving config: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		color.Red("Error writing config file: %v\n", err)
		os.Exit(1)
	}

	return userConfig
}

func choosePackageManager() string {
	color.Cyan("Please choose your package manager \U0001F447")

	prompt := promptui.Select{
		Label: "Package Managers",
		Items: languages.NodePackageManagers,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U0001F336 {{ . | red | bold }}",
			Inactive: "  {{ . | blue }}",
			Selected: "\U0001F336 {{ . | red | bold }}",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		color.Red("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
