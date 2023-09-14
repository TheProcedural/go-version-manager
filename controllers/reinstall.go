package controllers

import "fmt"

func ReinstallGoVersion(version string) {

	if !IsInstalled(version) {
		fmt.Printf("Go version %s is not installed.\nUse 'gov install %s' to install it.\n", version, version)
		return
	}

	RemoveGoVersion(version)
	InstallGoVersion(version)
}
