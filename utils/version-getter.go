package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"gitlab.com/bjerke-tek/gov/constants"
)

func GetLatestVersion(i string) string {
	var versionURL string

	if i == "go" {
		versionURL = "https://go.dev/VERSION?m=text"
	} else if i == "gov" {
		versionURL = "https://gov.theprocedural.com/version"
	}

	resp, err := http.Get(versionURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching the latest version.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	versionData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the latest Go version.")
		os.Exit(1)
	}

	var latestVersion string

	re := regexp.MustCompile(`(\d+\.\d+\.\d+)`)
	matches := re.FindStringSubmatch(string(versionData))

	if len(matches) < 2 {
		fmt.Println("Error extracting the semantic version.")
		os.Exit(1)
	}

	latestVersion = matches[1]

	return latestVersion
}

func GetLatestInstalledVersion(i string) string {

	if i == "gov" {
		return constants.Version
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory.")
		os.Exit(1)
	}

	dir := "go-versions"

	versionsDir := filepath.Join(homeDir, dir)

	if _, err := os.Stat(versionsDir); os.IsNotExist(err) {
		return ""
	} else if err != nil {
		fmt.Println("Error checking directory:", err)
		os.Exit(1)
	}

	subDirs, err := os.ReadDir(versionsDir)
	if err != nil {
		fmt.Println("Error listing versions:", err)
		os.Exit(1)
	}

	var latestVersion string

	for _, subDir := range subDirs {
		version := subDir.Name()
		if version > latestVersion {
			latestVersion = version
		}
	}

	return latestVersion
}
