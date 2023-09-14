package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func Prune() {
	currentVersion := runtime.Version()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(1)
	}

	goVersionsDir := filepath.Join(homeDir, "go-versions")

	if _, err := os.Stat(goVersionsDir); os.IsNotExist(err) {
		fmt.Println("No Go versions installed.")
		return
	} else if err != nil {
		fmt.Println("Error checking directory:", err)
		os.Exit(1)
	}

	subDirs, err := os.ReadDir(goVersionsDir)
	if err != nil {
		fmt.Println("Error listing Go versions:", err)
		os.Exit(1)
	}

	for _, subDir := range subDirs {
		version := subDir.Name()
		if version != currentVersion {
			if isSymbolicLinkToCurrent(homeDir, version) {
				fmt.Printf("Skipped removal of Go version %s as it's currently active.\n", version)
				continue
			}

			RemoveGoVersion(version)
		}
	}

	fmt.Println("Finished pruning Go versions.")
}

func isSymbolicLinkToCurrent(homeDir, version string) bool {
	currentLink := filepath.Join(homeDir, "go")
	linkPath, err := os.Readlink(currentLink)
	if err != nil {
		return false
	}

	linkVersion := filepath.Base(linkPath)
	return linkVersion == version
}
