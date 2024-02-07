package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	Controllers "github.com/nileshnk/media_upload_poc/communication/controllers"
	Types "github.com/nileshnk/media_upload_poc/communication/types"
	Utils "github.com/nileshnk/media_upload_poc/communication/utils"
)

func MainRouter() chi.Router {

	r := chi.NewRouter()
	daprRoutes := DaprRoutes()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Communication Service"))
	})

	r.Mount("/dapr", daprRoutes)

	r.Post("/send_email", Controllers.DaprEmailHandler)

	r.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		parsedData, errParse := decodePayload(body)
		Utils.CheckErrorWithCommonResponse(errParse)
		fmt.Println(parsedData)

		w.Write([]byte("Response from communication service! Yay!!"))
	})

	return r

}

func decodePayload(payloadBytes []byte) (Types.CommonResponse, error) {
	var parsedData Types.CommonResponse

	errJson2 := json.Unmarshal(payloadBytes, &parsedData)
	if errJson2 != nil {
		fmt.Println("errJson2---->", errJson2)
		return parsedData, errJson2
	}
	return parsedData, nil
}

func DaprRoutes() chi.Router {
	// Subscriptions
	r := chi.NewRouter()

	r.Get("/subscribe", Controllers.DaprSubscriptions)

	return r
}
