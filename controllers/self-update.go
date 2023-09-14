package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gitlab.com/bjerke-tek/gov/utils"
)

/***************************************************/
/*  This is not finsihed, macOS is not supported.  */
/* The shitty code quality must be address as well */
/***************************************************/

func SelfUpdate() error {

	latestGovVersion := utils.GetLatestVersion("gov")
	installedGovVersion := utils.GetLatestInstalledVersion("gov")

	if utils.IsNewerVersion(latestGovVersion, installedGovVersion) {
		fmt.Printf("Updating gov to version %s...\n", latestGovVersion)
	} else {
		fmt.Println("You are already running the latest version of gov.")
		return nil
	}

	// Define the URL to download the latest 'gov' binary.
	platform := runtime.GOOS
	arch := runtime.GOARCH
	url := fmt.Sprintf("https://gov.bjerkepedia.com/bin/%s/%s/gov", platform, arch)

	// Define the path to the existing 'gov' binary.
	existingBinaryPath := "/usr/local/bin/gov"

	// Create a temporary file to save the downloaded binary.
	tempBinaryPath := "/tmp/gov_temp"
	output, err := os.Create(tempBinaryPath)
	if err != nil {
		return err
	}
	defer output.Close()

	// Send an HTTP GET request to download the binary.
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check if the response status code is 200 OK.
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download 'gov', status code: %d", response.StatusCode)
	}

	// Copy the downloaded binary to the temporary file.
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	// Make the downloaded binary executable.
	err = os.Chmod(tempBinaryPath, 0755)
	if err != nil {
		return err
	}

	// Use sudo to move the temporary binary to /usr/local/bin (requires elevated privileges).
	cmd := exec.Command("sudo", "mv", tempBinaryPath, existingBinaryPath)
	cmd.Stdin = strings.NewReader("your_sudo_password") // Provide the sudo password here
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Successfully updated 'gov' to the latest version.")
	return nil
}
