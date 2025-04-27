package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/yufu085312/User-Management-Service/proto" // 自分のパスに変更
	"google.golang.org/grpc"

	// gRPCのエラーハンドリングを使うために追加
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	// ダミーのユーザーID（実際のアプリではDBに保存など）
	userID := int64(1)
	return &proto.CreateUserResponse{
		Id: userID,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	// ダミーのユーザー情報
	if req.Id == 1 {
		return &proto.GetUserResponse{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}, nil
	}
	return nil, status.Errorf(codes.NotFound, "User with ID %d not found", req.Id)
}

func main() {
	// gRPCサーバーの作成
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &server{})

	fmt.Println("gRPC server listening on port :50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
