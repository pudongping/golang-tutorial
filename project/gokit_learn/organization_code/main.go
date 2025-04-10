// main.go

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"go-tutorial/project/gokit_learn/sample_grpc_srv/pb"
)

var (
	httpAddr = flag.String("http-addr", ":8080", "HTTP listen address")
	grpcAddr = flag.String("grpc-addr", ":8972", "gRPC listen address")
)

func main() {
	bs := NewService()

	var g errgroup.Group

	// HTTP服务
	g.Go(func() error {
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			fmt.Printf("http: net.Listen(tcp, %s) failed, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		httpHandler := NewHTTPServer(bs)
		return http.Serve(httpListener, httpHandler)
	})

	g.Go(func() error {
		// gRPC服务
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			fmt.Printf("grpc: net.Listen(tcp, %s) faield, err:%v\n", *grpcAddr, err)
			return err
		}
		defer grpcListener.Close()
		s := grpc.NewServer()
		pb.RegisterAddServer(s, NewGRPCServer(bs))
		return s.Serve(grpcListener)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("server exit with err:%v\n", err)
	}
}
