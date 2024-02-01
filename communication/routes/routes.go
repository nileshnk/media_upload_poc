package routes

import (
	"github.com/go-chi/chi/v5"
)

func MainRouter() chi.Router {

	r := chi.NewRouter()
	// daprRoutes := DaprRoutes()

	// r.Mount("/dapr", daprRoutes)

	return r

}

// func DaprRoutes() chi.Router {
// 	// Subscriptions
// 	r := chi.NewRouter()

// 	emailSubscribeResponse := &common.Subscription{
// 		PubsubName: "communication",
// 		Topic: "email_delivery",
// 		Route: "/send_email",
// 	}


// 	r.Get("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		
// 	})


// 	return  r
// }