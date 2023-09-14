package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major         int
	Minor         int
	Patch         int
	PreRelease    string
	PreReleaseNum int
}

func IsNewerVersion(newerStr, olderStr string) bool {

	newer := parseSemanticVersion(newerStr)
	older := parseSemanticVersion(olderStr)

	if newer.Major > older.Major {
		return true
	} else if newer.Major < older.Major {
		return false
	}

	if newer.Minor > older.Minor {
		return true
	} else if newer.Minor < older.Minor {
		return false
	}

	if newer.Patch > older.Patch {
		return true
	} else if newer.Patch < older.Patch {
		return false
	}

	if newer.PreRelease != "" || older.PreRelease != "" {
		return isNewerPreRelease(newer, older)
	}

	return false
}

func parseSemanticVersion(versionStr string) Version {

	parts := strings.SplitN(versionStr, "-", 2)
	versionParts := strings.Split(parts[0], ".")

	// TODO: Handle error
	if len(versionParts) != 3 {
		fmt.Printf("invalid version format: %s", versionStr)
	}

	major, _ := strconv.Atoi(versionParts[0])
	minor, _ := strconv.Atoi(versionParts[1])
	patch, _ := strconv.Atoi(versionParts[2])

	preRelease := ""
	preReleaseNum := 0

	if len(parts) > 1 {
		preReleaseParts := strings.Split(parts[1], ".")
		preRelease = parts[1]

		if len(preReleaseParts) > 1 {
			preRelease = preReleaseParts[0]
			preReleaseNum, _ = strconv.Atoi(preReleaseParts[1])
		}
	}

	return Version{Major: major, Minor: minor, Patch: patch, PreRelease: preRelease, PreReleaseNum: preReleaseNum}
}

func isNewerPreRelease(newer, older Version) bool {

	if newer.PreRelease == "" {
		return false
	} else if older.PreRelease == "" {
		return true
	}

	if newer.PreRelease == older.PreRelease {
		return newer.PreReleaseNum > older.PreReleaseNum
	}

	return newer.PreRelease > older.PreRelease
}
