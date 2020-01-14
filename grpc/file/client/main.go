package main

import (
	"context"
	"fmt"
	pb "github.com/xiaobaiskill/grpc/file/grpc"
	"google.golang.org/grpc"
	"io"
	"os"
	"sync"
	"time"
)

var addr = "127.0.0.1:9042"
var clientConn *grpc.ClientConn

func init() {
	var err error
	clientConn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial err:", err)
		os.Exit(1)
	}
}

// 单列模式
var once sync.Once
var singleton *grpc.ClientConn

func Getinstance() *grpc.ClientConn {
	once.Do(func() {
		singleton = clientConn
	})
	return singleton
}

func IsFile(client pb.FileClient, name *pb.FileName) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	ok, err := client.IsFile(ctx, name)
	if err != nil {
		fmt.Println("isFile err:", err)
		return
	}
	fmt.Println("ok:", ok.Ok)
}

func ReadFile(client pb.FileClient, name *pb.FileName) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stream, err := client.ReadFile(ctx, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		filebyte, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println(string(filebyte.Byte))
	}
}

func WriteFile(client pb.FileClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stream, err := client.WriteField(ctx)
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	stream.Send(&pb.FileByte{
		Byte: []byte("jjjjjjjj\n"),
	})
	stream.Send(&pb.FileByte{
		Byte: []byte("mmmmmmmm\n"),
	})

	stream.Send(&pb.FileByte{
		Byte: []byte("zzzzzzz\n"),
	})

	fileok, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("write fileok err:", err)
		return
	}
	fmt.Println("fileok.Ok", fileok.Ok)
}

func main() {
	conn := Getinstance()

	client := pb.NewFileClient(conn)
	IsFile(client, &pb.FileName{
		Name: "./jmz.txt",
	})

	ReadFile(client, &pb.FileName{
		Name: "./jmz.txt",
	})

	WriteFile(client)

}
