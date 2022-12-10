package main

import (
	"fmt"
	"log"
	"net"
	"path/filepath"
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

	resp, err = pb.GetVersion()
	resp.Print("Client")

	return err
}
