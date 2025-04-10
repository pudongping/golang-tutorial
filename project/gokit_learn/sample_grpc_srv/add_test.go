// add_test.go
package main

import (
	"context"
	"log"
	"net"
	"testing"

	"go-tutorial/project/gokit_learn/sample_grpc_srv/pb"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// 使用bufconn构建测试链接，避免使用实际端口号启动服务

const bufSize = 1024 * 1024

var bufListener *bufconn.Listener

func init() {
	bufListener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	gs := NewGRPCServer(addService{})
	pb.RegisterAddServer(s, gs)
	go func() {
		if err := s.Serve(bufListener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return bufListener.Dial()
}

func TestSum(t *testing.T) {
	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddClient(conn)

	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 10,
		B: 2,
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(12), resp.V)
}

func TestConcat(t *testing.T) {
	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddClient(conn)

	resp, err := c.Concat(context.Background(), &pb.ConcatRequest{
		A: "10",
		B: "2",
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "102", resp.V)
}
