package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"runtime/debug"
	"time"

	pb "github.com/gnue/version"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type versionCommand struct {
}

func (c *versionCommand) Execute(args []string) error {
	network := "unix"
	if filepath.Ext(_opts.Addr) != ".sock" {
		network = "tcp"
	}

	dialer := func(a string, t time.Duration) (net.Conn, error) {
		return net.Dial(network, a)
	}

	conn, err := grpc.Dial(_opts.Addr, grpc.WithInsecure(), grpc.WithDialer(dialer))

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	//接続は最後に必ず閉じる
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	//サーバーに対してリクエストを送信する
	resp, err := client.Version(context.Background(), &pb.VersionRequest{})
	if err == nil {
		resp.Print("Server")
	}

	fmt.Println()

	resp, err = GetVersion()
	resp.Print("Client")

	return err
}

func GetVersion() (*pb.VersionReply, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// Goモジュールが無効など
		return nil, errors.New("no BuildInfo")
	}
	ver := info.Main.Version
	vcs := pb.GetSetting(info.Settings, "vcs")
	rev := pb.GetSetting(info.Settings, "vcs.revision")
	modified := pb.GetSetting(info.Settings, "modified")
	return &pb.VersionReply{Version: ver, Vcs: vcs, Revision: rev, Modified: modified, GoVersion: info.GoVersion}, nil
}
