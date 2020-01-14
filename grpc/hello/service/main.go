package main

import (
	"context"
	"fmt"
	pb "github.com/xiaobaiskill/grpc/hello/grpc"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, name *pb.Name) (result *pb.Result, err error) {
	result = &pb.Result{
		Res: fmt.Sprintf("hello %s", name.Name),
	}
	return
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9904")
	if err != nil {
		fmt.Println("network err:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		fmt.Println("failed to serve:", err)
	}

}
