package pb

import (
	"fmt"
	"runtime/debug"
)

func (resp *VersionReply) Print(name string) {
	fmt.Printf("%s:\n", name)

	fmt.Printf("\tversion:\t%s\n", resp.Version)
	fmt.Printf("\tGit commit:\t%s\n", resp.Revision)
	fmt.Printf("\tGo version:\t%s\n", resp.GoVersion)
}

func getSetting(settings []debug.BuildSetting, key string) string {
	for _, s := range settings {
		if s.Key == key {
			return s.Value
		}
	}

	return ""
}
