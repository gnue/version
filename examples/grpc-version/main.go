package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

var _opts struct {
	Addr string `short:"a" long:"addr" default:"/tmp/version.sock" description:"server address host:post"`

	Server serverCommand  `command:"server" subcommands-optional:"true" description:"server commands"`
	Client versionCommand `command:"version" subcommands-optional:"true" description:"server and client version"`
}

func main() {
	_, err := flags.Parse(&_opts)
	if err != nil {
		os.Exit(1)
	}
}
