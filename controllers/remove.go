package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RemoveGoVersion(version string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(1)
	}

	installDir := filepath.Join(homeDir, "go-versions", version)

	if _, err := os.Stat(installDir); os.IsNotExist(err) {
		fmt.Printf("Go version %s is not installed. Use 'gov install %s' to install it.\n", version, version)
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Error checking directory:", err)
		os.Exit(1)
	}

	err = os.RemoveAll(installDir)

	if err != nil {
		if os.IsPermission(err) {
			cmd := fmt.Sprintf("sudo rm -rf %s", installDir)
			_, cmdErr := exec.Command("sh", "-c", cmd).CombinedOutput()
			if cmdErr != nil {
				fmt.Printf("Error removing Go version %s: %v\n", version, cmdErr)
				os.Exit(1)
			}
			fmt.Printf("Removed Go version %s with elevated privileges.\n", version)
		} else {
			fmt.Println("Error removing Go version:", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Removed Go version %s.\n", version)
	}
}
