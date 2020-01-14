package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/xiaobaiskill/grpc/file/grpc"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

type File struct{}

func (f *File) File(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// 简单grpc
func (f *File) IsFile(ctx context.Context, file *pb.FileName) (*pb.FileOk, error) {
	return &pb.FileOk{
		Ok: f.File(file.Name),
	}, nil
}

// grpc服务端流
func (f *File) ReadFile(file *pb.FileName, stream pb.File_ReadFileServer) error {
	if !f.File(file.Name) {
		return errors.New("file is not path")
	}
	fileconetxt, err := os.Open(file.Name)
	if err != nil {
		return err
	}
	defer fileconetxt.Close()

	buf := make([]byte, 1024)
	for {
		n, err := fileconetxt.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		stream.Send(&pb.FileByte{
			Byte: buf[:n],
		})
	}

	return nil
}

// grpc客户端流
func (f *File) WriteField(stream pb.File_WriteFieldServer) error {
	filePath := "./" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		return err
	}
	defer file.Close()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.FileOk{
				Ok: true,
			})
			return nil
		}
		if err != nil {

			os.Remove(filePath)
			return err
		}
		_, err = file.Write(in.Byte)
		if err != nil {
			os.Remove(filePath)
			return err
		}
	}
	return nil
}

func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:9042")
	if err != nil {
		fmt.Println("listener err", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterFileServer(s, &File{})
	if err := s.Serve(conn); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
