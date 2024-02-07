package routes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	Controllers "github.com/nileshnk/media_upload_poc/media/controllers"
	Middleware "github.com/nileshnk/media_upload_poc/media/middleware"
)

func MainRouter(ctx context.Context) chi.Router {

	r := chi.NewRouter()
	r.Use(Middleware.AuthMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Media Service"))
	})

	r.Post("/upload", Controllers.UploadHandlerfunc)

	return r
}
