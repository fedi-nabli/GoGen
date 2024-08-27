package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fedi-nabli/GoGen/src/internal/config"
	"github.com/fedi-nabli/GoGen/src/internal/languages"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type ProjectConfig struct {
	Name  string
	Stack int
	Path  string
}

func GenerateProject(osNum int) {
	projectConfig := ProjectConfig{}
	user_config := config.LoadUserConfig()

	// Get project name
	prompt := promptui.Prompt{
		Label: "Enter project name",
		Validate: func(s string) error {
			if len(s) <= 2 {
				return fmt.Errorf("project name must be at least 3 characters")
			}
			return nil
		},
	}

	name, err := prompt.Run()
	if err != nil {
		color.Red("Error getting name")
	}

	projectConfig.Name = name

	// Get project stack
	projectConfig.Stack = chooseProjectStack()

	// Get project name
	path_prompt := promptui.Prompt{
		Label: "Enter project path",
		Validate: func(s string) error {
			if len(s) == 0 {
				return fmt.Errorf("project path cannot be null")
			}
			return nil
		},
	}

	path, err := path_prompt.Run()
	if err != nil {
		color.Red("Error getting path")
	}

	projectConfig.Path = path
	os.Chdir(projectConfig.Path)

	// Check if the path exists and is a directory
	info, err := os.Stat(projectConfig.Path)
	if os.IsNotExist(err) {
		fmt.Println("Path does not exist:", projectConfig.Path)
		return
	}
	if err != nil {
		fmt.Println("Error accessing the path:", err)
		return
	}
	if !info.IsDir() {
		fmt.Println("The path is not a directory:", projectConfig.Path)
		return
	}

	// Attempt to change directory
	err = os.Chdir(projectConfig.Path)
	if err != nil {
		fmt.Println("Failed to change directory:", err)
		return
	}

	// Confirm current directory change
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Start generating the project
	color.Green("Changed working directory to: %s\n", workingDir)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error getting user home directory: %v\n", err)
	}
	scriptsDir := filepath.Join(homeDir, "scripts")
	var commandFile string

	switch projectConfig.Stack {
	case languages.C:
		commandFile = filepath.Join(scriptsDir, "generate_C_project.txt")
	case languages.MERN:
		CheckAndInstallTool(user_config.PackageManager, osNum)
		commandFile = filepath.Join(scriptsDir, "generate_MERN_project.txt")
	case languages.MEVN:
		CheckAndInstallTool(user_config.PackageManager, osNum)
		commandFile = filepath.Join(scriptsDir, "generate_MEVN_project.txt")
	case languages.MEAN:
		CheckAndInstallTool(user_config.PackageManager, osNum)
		commandFile = filepath.Join(scriptsDir, "generate_MEAN_project.txt")
	}

	absCommandFile, err := filepath.Abs(commandFile)
	if err != nil {
		color.Red("failed to get absolute path of command file: %v", err)
	}

	args := map[string]string{
		"PROJECT_NAME":    projectConfig.Name,
		"PACKAGE_MANAGER": user_config.PackageManager,
	}

	executeCommandsFromFile(absCommandFile, args)

	fmt.Println("Project generated successfully!")
}

func chooseProjectStack() int {
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
	case "C":
		return languages.C
	case "C Makefile":
		return languages.C_MAKEFILE
	case "C CMake":
		return languages.C_CMAKE
	case "C++ Makefile":
		return languages.CXX_MAKEFILE
	case "C++ CMake":
		return languages.CXX_CMAKE
	case "Rust":
		return languages.RUST
	case "Rust Lib":
		return languages.RUST_LIB
	case "Go":
		return languages.GO
	case "Flask":
		return languages.FLASK
	case "Node Express":
		return languages.NODE_EXPRESS
	case "Node Express Typescript":
		return languages.NODE_EXPRESS_TYPESCRIPT
	case "MERN":
		return languages.MERN
	case "MEVN":
		return languages.MEVN
	case "MEAN":
		return languages.MEAN
	default:
		color.Red("Please choose a supported Language/Stack")
	}

	return 0
}

func executeCommandsFromFile(filename string, args map[string]string) {
	file, err := os.Open(filename)
	if err != nil {
		color.Red("failed to open command file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := scanner.Text()

		for key, value := range args {
			placeholder := fmt.Sprintf("{{%s}}", key)
			command = strings.ReplaceAll(command, placeholder, value)
		}

		// Split the command into parts
		parts := strings.Fields(command)
		if len(parts) == 0 {
			continue // Skip empty lines
		}

		// Handle 'cd' command
		if parts[0] == "cd" {
			if len(parts) < 2 {
				color.Red("cd command requires a directory argument")
				return
			}

			err := os.Chdir(parts[1])
			if err != nil {
				color.Red("failed to change directory to %s: %v", parts[1], err)
				return
			}
			color.White("Changed directory to: %s\n", parts[1])
			continue
		}

		cmd := exec.Command(parts[0], parts[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		color.White("Executing: %s\n", command)
		err := cmd.Run()
		if err != nil {
			color.Red("command failed: %s\nError: %v", command, err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		color.Red("error reading command file: %v", err)
		return
	}
}
