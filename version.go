package version

import (
	"errors"
	"fmt"
	"runtime/debug"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative version.proto

func (resp *VersionReply) Print(name string) {
	var mod string

	if resp.Modified == "true" {
		mod = "(*)"
	}
	fmt.Printf("%s:\n", name)

	fmt.Printf("\tversion:\t%s\n", resp.Version)
	fmt.Printf("\t%s commit:\t%s\n", resp.Vcs, resp.Revision+mod)
	fmt.Printf("\tGo version:\t%s\n", resp.GoVersion)
}

func GetSetting(settings []debug.BuildSetting, key string) string {
	for _, s := range settings {
		if s.Key == key {
			return s.Value
		}
	}

	return ""
}

func GetVersion() (*VersionReply, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// Goモジュールが無効など
		return nil, errors.New("no BuildInfo")
	}
	ver := info.Main.Version
	vcs := GetSetting(info.Settings, "vcs")
	rev := GetSetting(info.Settings, "vcs.revision")
	modified := GetSetting(info.Settings, "modified")
	return &VersionReply{Version: ver, Vcs: vcs, Revision: rev, Modified: modified, GoVersion: info.GoVersion}, nil
}
