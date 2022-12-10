package main

import (
	"github.com/gnue/version"
)

type serverCommand struct {
}

func (c *serverCommand) Execute(args []string) error {
	return version.Run(_opts.Addr)
}
