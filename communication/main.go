package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	Config "github.com/nileshnk/media_upload_poc/communication/config"
	dapr "github.com/nileshnk/media_upload_poc/communication/dapr"
	router "github.com/nileshnk/media_upload_poc/communication/routes"
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
	router := router.MainRouter()
	r.Mount("/", router)

	serverHost := Config.GetConfig.Server.Host
	serverPort := Config.GetConfig.Server.Port
	addr := serverHost + ":" + strconv.Itoa(serverPort)
	fmt.Println("Server is running on ", addr)
	http.ListenAndServe(addr, r)
}

func CheckErrorWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}
