package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ListSupportedVersions() {
	fmt.Println("Fetching data...")

	resp, err := http.Get("https://golang.org/dl/?mode=json")
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching remote Go versions.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	var versions []struct {
		Version string `json:"version"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&versions); err != nil {
		fmt.Println("Error parsing remote Go versions:", err)
		os.Exit(1)
	}

	fmt.Println("Available remote Go versions:")
	for _, v := range versions {
		fmt.Println(v.Version[2:])
	}

}
