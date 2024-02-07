package routes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	Controllers "github.com/nileshnk/media_upload_poc/auth/controllers"
)

func MainRouter(ctx context.Context) chi.Router {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth Service"))
		// dapr.DaprClient.InvokeMethod(ctx, "communication", "/", "test")
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		w.Write(body)
		w.Write([]byte("Auth Service"))
	})

	r.Post("/validate_token", Controllers.ValidateTokenFromRequestBody)

	return r

}
