package controllers

import (
	"fmt"

	"gitlab.com/bjerke-tek/gov/utils"
)

func CheckForGoUpdates() {
	latestGoVersion := utils.GetLatestVersion("go")
	installedGoVersion := utils.GetLatestInstalledVersion("go")

	if utils.IsNewerVersion(latestGoVersion, installedGoVersion) {
		fmt.Printf("Version %s of Go is available\n", latestGoVersion)
	}
}

func CheckForGovUpdates() {
	latestGovVersion := utils.GetLatestVersion("gov")
	installedGovVersion := utils.GetLatestInstalledVersion("gov")

	if utils.IsNewerVersion(latestGovVersion, installedGovVersion) {
		fmt.Printf("Version %s of gov is available\n", latestGovVersion)
	}
}
