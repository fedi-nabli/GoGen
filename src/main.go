package main

import (
	"os"

	"github.com/fedi-nabli/GoGen/src/internal/config"
	"github.com/fedi-nabli/GoGen/src/internal/utils"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func main() {
	targetOS := config.LoadPackageConfig()
	targetOSNum := config.WINDOWS

	switch targetOS {
	case "windows":
		targetOSNum = config.WINDOWS
	case "linux":
		targetOSNum = config.LINUX
	case "macos":
		targetOSNum = config.MACOS
	default:
		targetOSNum = config.WINDOWS
	}

	color.White("Welcome to GoGen, your project generator")

	color.Cyan("Please choose your action \U0001F447")

	prompt := promptui.Select{
		Label: "Action",
		Items: []string{"Change your package manager", "Generate project"},
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

	switch result {
	case "Change your package manager":
		config.ChangePackageManager()
	case "Generate project":
		utils.GenerateProject(targetOSNum)
	}
}
