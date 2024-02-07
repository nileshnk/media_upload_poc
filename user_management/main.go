package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	Config "github.com/nileshnk/media_upload_poc/user_management/config"
	dapr "github.com/nileshnk/media_upload_poc/user_management/dapr"
	Routes "github.com/nileshnk/media_upload_poc/user_management/routes"
)



func main() {

	// Load environment variables
	Config.Load()

	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize Dapr Client
	DaprClientErr := dapr.DaprInit(ctx)
	CheckErrorWithPanic(DaprClientErr)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	mainRouter := Routes.MainRouter(ctx)
	r.Mount("/", mainRouter)

	/* DAPR GRPC implementation. Do it later*/
	// serverGrpcHost := Config.GetConfig.Server.GRPCHost
	// serverGrpcPort := Config.GetConfig.Server.GRPCPort
	// grpcAuthAddr := serverGrpcHost + ":" + strconv.Itoa(serverGrpcPort)
	
	// conn, err := grpc.Dial(grpcAuthAddr, grpc.WithInsecure(), grpc.WithBlock())
	// // credentials.TransportCredentials

	// CheckErrorWithPanic(err)
	// defer conn.Close()
	// client := pb.NewAuthServiceClient(conn)


	// ctx = metadata.AppendToOutgoingContext(ctx, "auth", "auth")

	// // r, err2 := client.Login(ctx, &pb.RefreshTokenRequest{RefreshToken: "refresh_token"})
	
	serverHttpHost := Config.GetConfig.Server.HTTPHost
	serverHttpPort := Config.GetConfig.Server.HTTPPort
	httpAddr := serverHttpHost + ":" + strconv.Itoa(serverHttpPort)
	fmt.Println("Server is running on ", httpAddr)
	http.ListenAndServe(httpAddr, r)
}

func CheckErrorWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}
