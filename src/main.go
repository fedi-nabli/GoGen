package main

import (
	"fmt"
	"os"

	"github.com/fedi-nabli/GoGen/src/internal/config"
	"github.com/fedi-nabli/GoGen/src/internal/languages"
	"github.com/fedi-nabli/GoGen/src/internal/utils"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func main() {
	config.LoadUserConfig()
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

	utils.CheckAndInstallTool("npm", targetOSNum)

	color.Cyan("Please choose your programming stack \U0001F447")

	prompt := promptui.Select{
		Label: "Programming Stack",
		Items: languages.LanguagesArray,
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
	default:
		fmt.Println("Please choose a supported Language/Stack")
	}
}
