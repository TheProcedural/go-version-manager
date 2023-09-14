package controllers

import (
	"fmt"
	"os"
	"path/filepath"
)

func DisplayCurrentVersion() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory.")
		os.Exit(1)
	}

	currentLink := filepath.Join(homeDir, "go")
	linkPath, err := os.Readlink(currentLink)
	if err != nil {
		fmt.Println("Error reading the currently used Go version.")
		os.Exit(1)
	}

	version := filepath.Base(linkPath)
	fmt.Printf("Currently using Go version %s.\n", version)
}
