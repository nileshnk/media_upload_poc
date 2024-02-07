package controllers

import (
	pb "github.com/nileshnk/media_upload_poc/auth/proto"
)

func GRPCLogin(email string, password string) *pb.LoginResponse {
	return &pb.LoginResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}
}

func GRPCRegister(email string, password string) *pb.RegisterResponse {
	return &pb.RegisterResponse{
		Success: "true",
		Message: "User Registered Successfully",
	}
}

func GRPCGenerateAccessToken(userId string) *pb.GenerateAccessTokenResponse {
	return &pb.GenerateAccessTokenResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}
}

func GRPCGenerateAccessTokenWithRefreshToken(refreshToken string) *pb.RefreshTokenResponse {
	return &pb.RefreshTokenResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}
}
