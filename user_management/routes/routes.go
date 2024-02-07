package routes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	Controllers "github.com/nileshnk/media_upload_poc/communication/controllers"
	"github.com/nileshnk/media_upload_poc/communication/dapr"
)

func MainRouter(ctx context.Context) chi.Router {

	r := chi.NewRouter()
	daprRoutes := DaprRoutes()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth Service"))
		dapr.DaprClient.InvokeMethod(ctx, "communication", "test", "test")
	})

	r.Mount("/dapr", daprRoutes)

	r.Post("/send_email", Controllers.DaprEmailHandler)

	r.Mount("/user", UserManagementRouter(ctx))

	return r
}

func DaprRoutes() chi.Router {
	// Subscriptions
	r := chi.NewRouter()

	r.Get("/subscribe", Controllers.DaprSubscriptions)

	return r
}


func UserManagementRouter(ctx context.Context) chi.Router {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Management Service"))
	})

	return r
}