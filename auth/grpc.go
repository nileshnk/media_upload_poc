package main

import (
	"context"

	controllers "github.com/nileshnk/media_upload_poc/auth/controllers"
	pb "github.com/nileshnk/media_upload_poc/auth/proto"
)

type GrpcAuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *GrpcAuthServiceServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return controllers.GRPCLogin(in.Email, in.Password), nil
}

func (s *GrpcAuthServiceServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return controllers.GRPCRegister(in.Email, in.Password), nil
}

func (s *GrpcAuthServiceServer) GenerateAccessToken(ctx context.Context, in *pb.GenerateAccessTokenRequest) (*pb.GenerateAccessTokenResponse, error) {
	return controllers.GRPCGenerateAccessToken(in.UserId), nil
}

func (s *GrpcAuthServiceServer) GenerateAccessTokenWithRefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return controllers.GRPCGenerateAccessTokenWithRefreshToken(in.RefreshToken), nil
}
