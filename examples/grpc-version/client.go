package main

import (
	"fmt"

	"github.com/gnue/version"
)

type versionCommand struct {
}

func (c *versionCommand) Execute(args []string) error {
	//サーバーに対してリクエストを送信する
	resp, err := version.ServerVersion(_opts.Addr)
	if err == nil {
		resp.Print("Server")
	}

	fmt.Println()

	resp, err = version.GetVersion()
	resp.Print("Client")

	return err
}
