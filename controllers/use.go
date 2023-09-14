package controllers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/***************************************************/
/*  This is not finsihed, macOS is not supported.  */
/* The shitty code quality must be address as well */
/***************************************************/

func UseGoVersion(version string) {

	if !IsInstalled(version) {
		fmt.Printf("Go version %s is not installed.\nUse 'gov install %s' to install it.\n", version, version)
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory.")
		os.Exit(1)
	}

	installDir := filepath.Join(homeDir, "go-versions", version)

	if _, err := os.Stat(installDir); err != nil {
		fmt.Printf("Go version %s is not installed.\nUse 'gov install %s' to install it.\n", version, version)
		os.Exit(1)
	}

	// Update the symbolic link to point to the selected Go version
	currentLink := filepath.Join(homeDir, "go")
	err = os.RemoveAll(currentLink)
	if err != nil {
		fmt.Println("Error switching Go version.")
		os.Exit(1)
	}

	err = os.Symlink(installDir, currentLink)
	if err != nil {
		fmt.Println("Error switching Go version.")
		os.Exit(1)
	}

	// Persist environment variable changes
	modifyShellConfigFile()
	reloadShell()

	fmt.Printf("Switched to Go version %s.\n", version)
}

func modifyShellConfigFile() {
	shell := os.Getenv("SHELL")
	configFile := ""

	switch shell {
	case "/bin/bash":
		configFile = filepath.Join(os.Getenv("HOME"), ".bashrc")
	case "/bin/zsh":
		configFile = filepath.Join(os.Getenv("HOME"), ".zshrc")
	}

	if configFile != "" {
		file, err := os.Open(configFile)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", configFile, err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		exists := false
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), `export PATH="$HOME/go/bin:$PATH"`) {
				exists = true
				break
			}
		}

		if !exists {
			appendToFile(configFile, `export PATH="$HOME/go/bin:$PATH"`)
		}
	}
}

func appendToFile(file, content string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error appending to file %s: %v\n", file, err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		fmt.Printf("Error writing to file %s: %v\n", file, err)
		os.Exit(1)
	}
}

func reloadShell() {
	exec.Command("sh", "-c", "exec $SHELL -l").CombinedOutput()
}
