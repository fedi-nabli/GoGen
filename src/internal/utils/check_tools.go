package utils

import (
	"os/exec"

	"github.com/fedi-nabli/GoGen/src/internal/config"

	"github.com/fatih/color"
)

func CheckAndInstallTool(tool string, os int) bool {
	path, err := exec.LookPath(tool)

	if path != "" {
		color.Green("%s is installed at %s", tool, path)
		return true
	}

	if err != nil {
		color.Red("%v is not installed", tool)
		color.Green("Installing %s...", tool)
	}

	var cmd *exec.Cmd

	switch os {
	case config.WINDOWS:
		cmd = chocoInstallTool(tool)
	case config.LINUX:
		cmd = aptInstallTool(tool)
	case config.MACOS:
		cmd = brewInstallTool(tool)
	default:
		cmd = nil
		color.Red("Operating system not supported yet!")
		return false
	}

	if cmd == nil {
		color.Red("No installation method for %s", tool)
		return false
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		color.Red("Failed to install %s: %v\n%s", tool, err, output)
		return false
	}

	color.Green("%s is installed successfully!", tool)
	return true
}

func chocoInstallTool(tool string) *exec.Cmd {
	return exec.Command("choco", "install", tool, "-y")
}

func aptInstallTool(tool string) *exec.Cmd {
	return exec.Command("sudo", "apt-get", "install", "-y", tool)
}

func brewInstallTool(tool string) *exec.Cmd {
	return exec.Command("brew", "install", tool)
}
