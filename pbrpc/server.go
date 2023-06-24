package main

import (
	"context"
	"log"
	"net"
	pbrpc "rpc_demo/pbrpc/demo"

	"google.golang.org/grpc"
	// 替换为你的 PB 生成包路径
)

type demoServiceServer struct {
	pbrpc.UnimplementedMyServiceServer
}

func (s *demoServiceServer) MyMethod(ctx context.Context, req *pbrpc.MyMessage) (*pbrpc.MyMessage, error) {
	// 处理请求并返回响应
	name := req.GetName()
	age := req.GetAge()
	// 进行你的逻辑处理...
	log.Printf("Received: %v, %v", name, age)
	// 构造响应消息
	resp := &pbrpc.MyMessage{
		Name: "John Doe",
		Age:  30,
	}
	return resp, nil
}

func main() {
	// 创建 TCP 监听器
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// 创建 gRPC 服务器
	s := grpc.NewServer()
	pbrpc.RegisterMyServiceServer(s, &demoServiceServer{})

	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
