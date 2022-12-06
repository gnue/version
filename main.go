package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	i, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Println("version = ", i.Main.Version)
	}
}
