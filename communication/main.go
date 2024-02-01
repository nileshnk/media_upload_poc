package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	Config "github.com/nileshnk/media_upload_poc/communication/config"
	dapr "github.com/nileshnk/media_upload_poc/communication/dapr"
)

func main() {

	// Load environment variables
	envErr := godotenv.Load(".env")
	CheckErrorWithPanic(envErr)
	Config.Load()
	
	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize Dapr Client
	DaprClientErr := dapr.DaprInit(ctx)
	CheckErrorWithPanic(DaprClientErr)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	serverHost := Config.GetConfig.Server.Host
	serverPort := Config.GetConfig.Server.Port
	addr := serverHost + ":" + strconv.Itoa(serverPort)

	http.ListenAndServe(addr, r)
}

func CheckErrorWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}
