package main

import (
	"context"
	"fmt"
	pb "github.com/xiaobaiskill/grpc/hello/grpc"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9904", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc conn err:", err)
		return
	}

	defer conn.Close()
	hc := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	name := &pb.Name{
		Name: "jmz",
	}
	result, err := hc.SayHello(ctx, name)
	if err != nil {
		fmt.Println("处理失败", err)
		return
	}
	fmt.Println(result.Res)
}
