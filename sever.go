package version

import (
	context "context"
	"errors"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"syscall"

	grpc "google.golang.org/grpc"
)

type Server struct {
}

func NetworkFromAddr(addr string) string {
	network := "unix"
	if filepath.Ext(addr) != ".sock" {
		network = "tcp"
	}
	return network
}

func Run(addr string) error {
	network := NetworkFromAddr(addr)

	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	//新しいgRPCサーバーのインスタンスを作成
	s := grpc.NewServer()
	//gRPCサーバーを保存する
	RegisterGreeterServer(s, &Server{})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ch
		s.GracefulStop()
	}()

	s.Serve(lis)

	return err
}

// リクセスト(Name)を受け取り、レスポンス(Message)を返す
func (s *Server) Version(ctx context.Context, in *VersionRequest) (*VersionReply, error) {
	return GetVersion()
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
	mod := false
	if modified == "true" {
		mod = true
	}
	return &VersionReply{Version: ver, Vcs: vcs, Revision: rev, Modified: mod, GoVersion: info.GoVersion}, nil
}
