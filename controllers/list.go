package controllers

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListInstalledVersions() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory.")
		os.Exit(1)
	}

	installDir := filepath.Join(homeDir, "go-versions")

	dirs, err := filepath.Glob(filepath.Join(installDir, "*"))
	if err != nil {
		fmt.Println("Error listing installed Go versions.")
		os.Exit(1)
	}

	fmt.Println("Installed Go versions:")
	for _, dir := range dirs {
		version := filepath.Base(dir)
		fmt.Println(version)
	}
}
