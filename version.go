package version

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative version.proto

func (resp *VersionReply) Print(name string) {
	rev := resp.Revision

	if 7 < len(rev) {
		rev = rev[:7]
	}
	if resp.Modified {
		rev += "(*)"
	}
	fmt.Printf("%s:\n", name)

	fmt.Printf("\tversion:\t%s\n", resp.Version)
	fmt.Printf("\t%s commit:\t%s\n", strcase.ToCamel(resp.Vcs), rev)
	fmt.Printf("\tGo version:\t%s\n", resp.GoVersion)
}
