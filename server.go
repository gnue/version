package main

import (
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"syscall"

	pb "github.com/gnue/version/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type serverCommand struct {
}

func (c *serverCommand) Execute(args []string) error {
	//requestを受け付けるportを指定する
	network := "unix"
	if filepath.Ext(_opts.Addr) != ".sock" {
		network = "tcp"
	}

	lis, err := net.Listen(network, _opts.Addr)
	if err != nil {
		return err
	}

	//新しいgRPCサーバーのインスタンスを作成
	s := grpc.NewServer()
	//gRPCサーバーを保存する
	pb.RegisterGreeterServer(s, &server{})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ch
		s.GracefulStop()
	}()

	s.Serve(lis)

	return err
}

type server struct{}

func getSetting(settings []debug.BuildSetting, key string) string {
	for _, s := range settings {
		if s.Key == key {
			return s.Value
		}
	}

	return ""
}

// リクセスト(Name)を受け取り、レスポンス(Message)を返す
func (s *server) Version(ctx context.Context, in *pb.VersionRequest) (*pb.VersionReply, error) {
	return GetVersion()
}
