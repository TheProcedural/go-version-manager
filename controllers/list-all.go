package controllers

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
)

func ListAllVersions() {
	fmt.Println("Fetching data...")

	versionPattern := regexp.MustCompile(`go(\d+\.\d+(\.\d+)?(rc\d+|beta\d+)?)`)
	uniqueVersions := make(map[string]bool)

	resp, err := http.Get("https://golang.org/dl/")
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching Go download page.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	htmlContent := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			htmlContent = append(htmlContent, buffer[:n]...)
		}
		if err != nil {
			break
		}
	}

	matches := versionPattern.FindAllStringSubmatch(string(htmlContent), -1)
	if len(matches) == 0 {
		fmt.Println("No Go versions found.")
		return
	}

	fmt.Println("All Go versions; including, unsupported, betas and release candidates (Oldest to Newest):")
	for i := len(matches) - 1; i >= 0; i-- {
		versionStr := matches[i][0]
		if !uniqueVersions[versionStr] {
			fmt.Println(versionStr[2:])
			uniqueVersions[versionStr] = true
		}
	}
}
