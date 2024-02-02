package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	Services "github.com/nileshnk/media_upload_poc/communication/services"
	Types "github.com/nileshnk/media_upload_poc/communication/types"
	Objx "github.com/stretchr/objx"
)

func DaprEmailHandler(w http.ResponseWriter, r *http.Request) {
	var payload Types.EmailPayload
	var pubsubPayload common.TopicEvent
	json.NewDecoder(r.Body).Decode(&pubsubPayload)

	m, ObjxErr := Objx.FromJSON(pubsubPayload.Data.(string))
	if ObjxErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ObjxErr.Error()))
		return
	}

	if !m.Get("recipient").IsStr() || !m.Get("recipient").IsStrSlice() {
		log.Println("Recipient is not a string or string slice")
	}

	payload = ParseDaprEmailPayload(pubsubPayload)

	err := Services.SendEmail(payload.Recipient, payload.Subject, payload.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Email Sent Successfully!"))
}

func DaprSubscriptions(w http.ResponseWriter, r *http.Request) {
	emailSubscribeResponse := &common.Subscription{
		PubsubName: "communication",
		Topic:      "email_delivery",
		Route:      "/send_email",
	}

	var subscriptionResponse []common.Subscription
	subscriptionResponse = append(subscriptionResponse, *emailSubscribeResponse)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptionResponse)
}

func ParseDaprEmailPayload(payload common.TopicEvent) Types.EmailPayload {
	var emailPayload Types.EmailPayload
	m, ObjxErr := Objx.FromJSON(payload.Data.(string))
	if ObjxErr != nil {
		log.Println(ObjxErr)
	}

	strIn := m.Get("recipient").InterSlice()
	var strArr []string
	for _, v := range strIn {
		if len(v.(string)) == 0 {
			continue
		}
		strArr = append(strArr, v.(string))
	}

	emailPayload.Recipient = strArr
	emailPayload.Subject = m.Get("subject").String()
	emailPayload.Body = m.Get("body").String()
	return emailPayload
}
