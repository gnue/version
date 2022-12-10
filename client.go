package version

import (
	context "context"
	"log"
	"net"
	"time"

	grpc "google.golang.org/grpc"
)

func ServerVersion(addr string) (*VersionReply, error) {
	network := NetworkFromAddr(addr)

	dialer := func(a string, t time.Duration) (net.Conn, error) {
		return net.Dial(network, a)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithDialer(dialer))

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	//接続は最後に必ず閉じる
	defer conn.Close()

	client := NewGreeterClient(conn)

	//サーバーに対してリクエストを送信する
	return client.Version(context.Background(), &VersionRequest{})
}
