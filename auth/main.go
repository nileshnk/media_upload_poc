package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	Config "github.com/nileshnk/media_upload_poc/auth/config"
	dapr "github.com/nileshnk/media_upload_poc/auth/dapr"
	pb "github.com/nileshnk/media_upload_poc/auth/proto"
	Routes "github.com/nileshnk/media_upload_poc/auth/routes"
	"google.golang.org/grpc"
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

	serverGrpcHost := Config.GetConfig.Server.GRPCHost
	serverGrpcPort := Config.GetConfig.Server.GRPCPort
	grpcAddr := serverGrpcHost + ":" + strconv.Itoa(serverGrpcPort)

	grpcListener, grpcListenerErr := net.Listen("tcp", grpcAddr)
	CheckErrorWithPanic(grpcListenerErr)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &GrpcAuthServiceServer{})

	serverHttpHost := Config.GetConfig.Server.HTTPHost
	serverHttpPort := Config.GetConfig.Server.HTTPPort
	httpAddr := serverHttpHost + ":" + strconv.Itoa(serverHttpPort)
	httpListner, httpListnerErr := net.Listen("tcp", httpAddr)
	CheckErrorWithPanic(httpListnerErr)
	fmt.Println("HTTP Server is running on ", httpAddr)
	http.Serve(httpListner, r)

	fmt.Println("GRPC Server is running on ", grpcAddr)
	if err := grpcServer.Serve(grpcListener); err != nil {
		fmt.Println("Failed to serve gRPC server over addr ", grpcAddr)
	}

}

func CheckErrorWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}
