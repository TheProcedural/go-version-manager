package controllers

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gitlab.com/bjerke-tek/gov/utils"
)

/**********************************************************/
/* This is not finsihed, macOS is not properly supported. */
/*     The shitty code quality must be address as well    */
/**********************************************************/

func InstallGoVersion(version string) {

	if version == "latest" {
		fmt.Println("Fetching the latest Go version...")
		version = utils.GetLatestVersion("go")
	}

	if IsInstalled(version) {
		fmt.Printf("Go version %s is already installed.\nUse 'gov reinstall %s' to reinstall it.\n", version, version)
		return
	}

	fmt.Printf("Downloading Go version %s...\n", version)
	osArch := runtime.GOARCH
	var url string

	switch runtime.GOOS {
	case "linux":
		url = fmt.Sprintf("https://golang.org/dl/go%s.%s-%s.tar.gz", version, runtime.GOOS, osArch)
	case "darwin":
		url = fmt.Sprintf("https://golang.org/dl/go%s.%s-%s.tar.gz", version, runtime.GOOS, osArch)
	default:
		fmt.Println("Unsupported OS")
		os.Exit(1)
	}

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error downloading Go version.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory.")
		os.Exit(1)
	}

	installDir := filepath.Join(homeDir, "go-versions", version)
	os.MkdirAll(installDir, os.ModePerm)

	fmt.Printf("Installing Go version %s...\n", version)

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("Error extracting Go archive.")
		os.Exit(1)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error extracting Go archive.")
			os.Exit(1)
		}
		header.Name = strings.TrimPrefix(header.Name, "go/")
		targetPath := filepath.Join(installDir, header.Name)
		if header.FileInfo().IsDir() {
			os.MkdirAll(targetPath, os.ModePerm)
		} else {
			file, err := os.Create(targetPath)
			if err != nil {
				fmt.Println("Error extracting Go archive.")
				os.Exit(1)
			}
			defer file.Close()
			_, err = io.Copy(file, tarReader)
			if err != nil {
				fmt.Println("Error extracting Go archive.")
				os.Exit(1)
			}

			if isExecutable(header) {
				if err := os.Chmod(targetPath, 0755); err != nil {
					fmt.Println("Error setting executable permission on binary file:", err)
					os.Exit(1)
				}
			}
		}
	}

	fmt.Printf("Go version %s installed successfully!\n", version)
}

func isExecutable(header *tar.Header) bool {
	return header.Mode&0111 != 0
}

func IsInstalled(version string) bool {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory.")
		os.Exit(1)
	}

	dir := "go-versions"

	versionsDir := filepath.Join(homeDir, dir)

	if _, err := os.Stat(versionsDir); os.IsNotExist(err) {
		return false
	} else if err != nil {
		fmt.Println("Error checking directory:", err)
		os.Exit(1)
	}

	subDirs, err := os.ReadDir(versionsDir)
	if err != nil {
		fmt.Println("Error listing versions:", err)
		os.Exit(1)
	}

	for _, subDir := range subDirs {
		if subDir.Name() == version {
			return true
		}
	}

	return false
}
