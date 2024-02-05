package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	Config "github.com/nileshnk/media_upload_poc/media/config"
	routes "github.com/nileshnk/media_upload_poc/media/routes"
)

// MediaService is the interface for the media service
type MediaService interface {
	UploadMedia(ctx context.Context, media []byte) (string, error)
}

// MediaRepository is the interface for the media repository
type MediaRepository interface {
	UploadMedia(media []byte) (string, error)
}

// Media is the media structure
type Media struct {
	ID   string
	Data []byte
}

func main() {

	// load environment variables
	// godotenv.Load(".env")
	Config.Load()


	// create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel();

	// create router
	r := chi.NewRouter()

	r.Mount("/", routes.MainRouter(ctx))



	serverHost := Config.GetConfig.Server.HTTPHost
	serverPort := Config.GetConfig.Server.HTTPPort
	httpAddr := serverHost + ":" + strconv.Itoa(serverPort)

	lis, err := net.Listen("tcp", httpAddr)
	CheckErrorWithPanic(err)
	fmt.Println("Server is running on ", httpAddr)
	http.Serve(lis, r)
}

func CheckErrorWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}