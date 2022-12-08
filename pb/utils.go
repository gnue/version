package pb

import "fmt"

func (resp *VersionReply) Print(name string) {
	fmt.Println("%s:", name)

	fmt.Printf("\tversion:\t%s\n", resp.Version)
	fmt.Printf("\tGit commit:\t%s\n", resp.Revision[:7])
	fmt.Printf("\tGo version:\t%s\n", resp.GoVersion)
}
