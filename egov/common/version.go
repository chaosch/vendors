package common

import "strings"

func MakeVersion(branch, build string) string {
	version := ""
	branch = strings.ToLower(branch)
	if strings.HasPrefix(branch, "release/") {
		version = strings.TrimPrefix(branch, "release/")
		version = "beta" + ":" + version + "." + build
	} else if strings.HasPrefix(branch, "master/") {
		version = strings.TrimPrefix(branch, "master/")
		version = "rel" + ":" + version + "." + build
	} else {
		version = "dev" + ":" + build
	}
	return version
}
