package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	Controllers "github.com/nileshnk/media_upload_poc/communication/controllers"
)

func MainRouter() chi.Router {

	r := chi.NewRouter()
	daprRoutes := DaprRoutes()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Communication Service"))
	})

	r.Mount("/dapr", daprRoutes)

	r.Post("/send_email", Controllers.DaprEmailHandler)

	return r

}

func DaprRoutes() chi.Router {
	// Subscriptions
	r := chi.NewRouter()

	r.Get("/subscribe", Controllers.DaprSubscriptions)

	return r
}

